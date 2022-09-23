package config

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/joho/godotenv"
	"os"
)

var (
	Client       *oss.Client
	AccessKey    string
	AccessSecret string
	OssEndPoint  string
	BucketName   string
)

func GetClientBucket() (*oss.Bucket, error) {
	bucket, err := Client.Bucket(BucketName)
	if err != nil {
		fmt.Printf("获取bucket异常：%s\n", err.Error())
		return nil, err
	}
	return bucket, nil
}

func init() {
	//初始化配置文件
	err := godotenv.Load("./etc/test.env")
	if err != nil {
		fmt.Printf("获取配置文件错误：%s\n", err)
		os.Exit(1)
	}
	AccessKey = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndPoint = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName = os.Getenv("ALI_BUCKET_NAME")
	//初始化oss链接
	c, err := oss.New(OssEndPoint, AccessKey, AccessSecret)
	if err != nil {
		fmt.Printf("初始化阿里云OSS错误：%s\n", err)
		os.Exit(1)
	}
	Client = c
}
