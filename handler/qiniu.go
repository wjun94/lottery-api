package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"sync"

	"encoding/json"
	"io"
	"net/http"

	// "os"
	"strings"
	"time"

	"fmt"
	"mango-api/config"

	// "github.com/google/uuid"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

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

// UploadFile 微信小程序文件上传
func UploadFile(w http.ResponseWriter, r *http.Request) {
	var res map[string]string
	file, head, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	name := head.Filename
	// multipart.File转[]byte
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		fmt.Println("上传图片报错")
	}
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
	dataLen := int64(len(buf.Bytes()))
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()

		err1 := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(buf.Bytes()), dataLen, &putExtra)
		if err1 != nil {
			fmt.Println(err1)
		}

	}()
	wg.Wait()
	// fmt.Fprintln(w, downFile(name))
	url := downFile(name)
	res = map[string]string{"img": url}
	response, _ := json.Marshal(res)
	w.Write(response)
}

func ByteQiniu(val []byte, name string) string {
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	formUploader := storage.NewFormUploader(&cfg)
	dataLen := int64(len(val))
	var upToken = putPolicy.UploadToken(mac)
	putPolicy.Expires = 7200 //示例2小时有效期
	ret := storage.PutRet{}
	err1 := formUploader.Put(context.Background(), &ret, upToken, name, bytes.NewReader(val), dataLen, &putExtra)
	if err1 != nil {
		fmt.Println(err1)
	}
	return downFile(name)
}

// Base64Qiniu base64上传到七牛云
func Base64Qiniu(encoded string, name string) string {
	ddd, _ := base64.StdEncoding.DecodeString(strings.Split(encoded, ",")[1])

	// key := "99.png"
	key := name
	var upToken = putPolicy.UploadToken(mac)
	putPolicy.Expires = 7200 //示例2小时有效期

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// fmt.Println(ret.Key, ret.Hash)
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	dataLen := int64(len(ddd))
	err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(ddd), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return downFile(name)
}

// downFile 下载七牛云文件
func downFile(name string) string {
	mac := qbox.NewMac(accessKey, secretKey)
	key := name
	deadline := time.Now().Add(time.Second * 7600000000).Unix() // 3600为1小时有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	return privateAccessURL
}

// DeleteQiniu 删除七牛云文件
func DeleteQiniu(url string) {
	key := strings.Split(url, fmt.Sprintf("%s/", domain))[1]
	// days := 2 // 设置文件生存时间
	// err := bucketManager.DeleteAfterDays(bucket, key, days)
	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/* // UploadFile 上传到七牛云
func UploadFile(w http.ResponseWriter, r *http.Request) {
	//本地保存的文件夹名称
	uploadPath := "/files/"
	//获取文件内容 要这样获取
	file, head, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//创建文件夹
	pwd, _ := os.Getwd()
	//文件夹存在的话会返回一个错误，可以用`_`抛出去
	err = os.Mkdir(pwd+uploadPath, os.ModePerm)
	if err != nil {
		fmt.Println("dir is create Error")
	}
	fW, err := os.Create(pwd + uploadPath + head.Filename)
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fW.Close()
	//复制文件，保存到本地
	_, err = io.Copy(fW, file)
	if err != nil {
		fmt.Println("文件保存失败")
		return
	}
	// strings.Split("", ";")
	// head.Header["Content-Type"][0]
	name := fmt.Sprintf("%s.%s", uuid.New(), strings.Split(head.Header["Content-Type"][0], "/")[1])

	//调用七牛上传函数
	uploadQiniu(pwd+uploadPath+head.Filename, name)
	result := map[string]string{
		"image": downFile(name),
	}

	response, _ := json.Marshal(result)
	w.Write(response)
}

// uploadQiniu 上传图片到七牛云
func uploadQiniu(localFile string, name string) {
	// key := "99.png"
	key := name
	var upToken = putPolicy.UploadToken(mac)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// fmt.Println(ret.Key, ret.Hash)
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
} */
