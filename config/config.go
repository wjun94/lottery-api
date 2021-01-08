package config

/**
 * 文档作者: wjun94
 * 创建时间：2019年09月22日
 * 修改时间：2019年10月10日
 * 描述信息：配置文件
 */
import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config 配置文件
var Config = struct {
	DB struct {
		User     string `yaml:"user"`
		Base     string `yaml:"base"`
		Password string `yaml:"password"`
	}
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		Base     int    `yaml:"base"`
	}
	Route struct {
		PathPrefix string `yaml:"pathPrefix"`
	}
	Qiniu struct {
		AccessKey string `yaml:"accessKey"`
		SecretKey string `yaml:"secretKey"`
		Bucket    string `yaml:"bucket"`
		Domain    string `yaml:"domain"`
	}
}{}

func init() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("err:", err)
	}
	yaml.Unmarshal(yamlFile, &Config)
}
