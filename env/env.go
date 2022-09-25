package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var Config *AliOssEnv

type AliOssEnv struct {
	Provider     string
	AccessKey    string
	AccessSecret string
	Endpoint     string
	BucketName   string
}

func init() {
	//初始化配置文件
	err := godotenv.Load("/Users/cactus/Desktop/go_server/src/cloud-station/etc/test.env")
	if err != nil {
		fmt.Printf("获取配置文件错误：%s\n", err)
		os.Exit(1)
	}
	Config = &AliOssEnv{
		Provider:     "aliyun",
		AccessKey:    os.Getenv("ALI_AK"),
		AccessSecret: os.Getenv("ALI_SK"),
		Endpoint:     os.Getenv("ALI_OSS_ENDPOINT"),
		BucketName:   os.Getenv("ALI_BUCKET_NAME"),
	}
}
