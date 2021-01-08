package models

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mango-api/config"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var clientID = config.Config.Yzy.ClientID
var appkey = config.Config.Yzy.Appkey

// CreateYzyOrder 创建订单
// id:订单id name:用户姓名 address:地址 phone:手机号 code:产品编码 count:总张数 phs:照片对象
func CreateYzyOrder(id string, name string, address string, phone string, code string, count int, phs []Photos, options byte) string {
	comp := map[string]string{
		"ZHP101-FJL-3R":      "3R5寸 富士绒面 照片",
		"ZHP101-FJL-4R":      "4R6寸 富士绒面 照片",
		"ZHP106-FJG-3R-A101": "3R5寸 富士光面 塑封照片",
		"ZHP106-FJG-4R-A101": "4R6寸 富士光面 塑封照片",
	}
	url, fMd5 := createImg(phs)
	tm := strconv.Itoa(int(time.Now().Unix()))
	m5 := md5.New()
	io.WriteString(m5, clientID+appkey+tm+appkey)
	sign := fmt.Sprintf("%x", m5.Sum(nil)) // w.Sum(nil)将w的hash转成[]byte格式
	if options == 1 {
		// 封塑
		switch code {
		case "ZHP101-FJL-3R":
			code = "ZHP106-FJG-3R-A101"
		case "ZHP101-FJL-4R":
			code = "ZHP106-FJG-4R-A101"
		}
	}
	data := map[string]interface{}{
		"client_id":         clientID,
		"appkey":            appkey,
		"method":            "add_order",
		"timestamp":         tm,
		"sign":              strings.ToUpper(sign),
		"goods_id":          code,
		"goods_name":        comp[code],
		"goods_qty":         count,
		"receiver_name":     name,
		"receiver_mobile":   phone,
		"receiver_address":  address,
		"receiver_areacode": "020",
		"client_order_no":   id,
		"client_goods_id":   code,
		"client_goods_name": comp[code],
		"client_goods_qty":  count,
		"addresser_name":    "小项",
		"addresser_mobile":  "13588227124",
		"addresser_address": "杭州市金家渡南苑46号",
		"file_type":         "1",
		"file_url":          url,
		"file_md5":          fMd5,
	}
	jsonStr, _ := json.Marshal(data)
	resp, err := http.Post("http://erp.pyjyphoto.com/erp_open/",
		"application/x-www-form-urlencoded",
		strings.NewReader(string(jsonStr)))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	m := make(map[string]string)
	err1 := json.Unmarshal(body, &m)
	if err1 != nil {
		return ""
	}
	// fmt.Println(m)
	return m["order_no"]
}

// -----------
func createImg(arr []Photos) (string, string) {
	path1 := strconv.FormatInt(time.Now().Unix(), 10)
	//图片正则
	//创建多级目录
	err2 := os.MkdirAll("./"+path1, os.ModePerm)
	if err2 != nil {
		log.Fatal(err2)
	}
	for i, v := range arr {
		// reg, _ := regexp.Compile(`(\w|\d|_|-)*.(jpg|png|jpeg)`)
		name := "./" + path1 + "/" + strconv.Itoa(i+1) + "(" + strconv.Itoa(v.Count) + ")" + strings.Split(strings.Split(v.Img, "?")[0], "http://www.cnicu.cn/")[1]
		//通过http请求获取图片的流文件
		resp, _ := http.Get(v.Img)
		body, _ := ioutil.ReadAll(resp.Body)
		out, _ := os.Create(name)
		io.Copy(out, bytes.NewReader(body))
		defer out.Close()
	}

	dir, _ := os.Getwd()
	zipName := path1 + ".zip"
	Zip("./"+path1, zipName)
	UploadZip(dir+"/"+zipName, zipName)
	str := downFile(zipName)
	zMd5 := ZipMd5(zipName)
	os.RemoveAll(dir + "/" + path1)
	os.RemoveAll(dir + "/" + zipName)
	return str, zMd5
}

// Zip 压缩文件
func Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()
	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		if info.IsDir() {
			if srcFile == path {
				return nil
			}
			path += "/"
		} else {
			header.Method = zip.Deflate
		}
		base := filepath.Base(srcFile)
		header.Name = path[len(base)+1:]
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
	return err
}

var accessKey string = config.Config.Qiniu.AccessKey
var secretKey string = config.Config.Qiniu.SecretKey
var bucket string = config.Config.Qiniu.Bucket
var domain string = config.Config.Qiniu.Domain
var mac = qbox.NewMac(accessKey, secretKey)
var putPolicy = storage.PutPolicy{
	Scope: bucket,
}
var cfg = storage.Config{
	// 是否使用https域名进行资源管理
	UseHTTPS:      false,
	Zone:          &storage.ZoneHuadong,
	UseCdnDomains: false,
}
var bucketManager = storage.NewBucketManager(mac, &cfg)

// UploadZip zip上传
func UploadZip(localFile string, name string) {
	key := name
	var upToken = putPolicy.UploadToken(mac)
	putPolicy.Expires = 7200 //示例2小时有效期

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 构建表单上传的对象
		err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(ret.Key, ret.Hash)
	}()
	wg.Wait()
}

func downFile(name string) string {
	mac := qbox.NewMac(accessKey, secretKey)
	key := name
	deadline := time.Now().Add(time.Second * 7600000000).Unix() // 3600为1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	return privateAccessURL
}

func ZipMd5(filename string) string {
	h := md5.New()
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	io.Copy(h, f)

	return hex.EncodeToString(h.Sum(nil))
}

/** ------------------------------ */
func SelectYzy(orderID string, yzyID string) (string, string, int) {
	tm := strconv.Itoa(int(time.Now().Unix()))
	m5 := md5.New()
	io.WriteString(m5, clientID+appkey+tm+appkey)
	sign := fmt.Sprintf("%x", m5.Sum(nil)) // w.Sum(nil)将w的hash转成[]byte格式
	data := map[string]interface{}{
		"client_id":       clientID,
		"appkey":          appkey,
		"sign":            strings.ToUpper(sign),
		"timestamp":       tm,
		"method":          "get_order_info",
		"client_order_no": orderID,
		"order_no":        yzyID,
	}
	jsonStr, _ := json.Marshal(data)
	resp, err := http.Post("http://erp.pyjyphoto.com/erp_open/",
		"application/x-www-form-urlencoded",
		strings.NewReader(string(jsonStr)))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var m struct {
		Success string                   `json:"success"`
		Detail  []map[string]interface{} `json:"detail"`
	}
	err1 := json.Unmarshal(body, &m)
	if err1 != nil {
		return "", "", 0
	}
	if m.Success == "true" {
		res1 := m.Detail[0]
		sta, _ := strconv.Atoi(res1["order_status"].(string))
		return res1["express_no"].(string), res1["express_name"].(string), sta
	}
	return "", "", 0
}
