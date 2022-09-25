package config

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cactu/cloud-station/env"
	"os"
)

var (
	Client *oss.Client
)

func GetClientBucket() (*oss.Bucket, error) {
	bucket, err := Client.Bucket(env.Config.BucketName)
	if err != nil {
		fmt.Printf("获取bucket异常：%s\n", err.Error())
		return nil, err
	}
	return bucket, nil
}

func init() {
	//初始化oss链接
	c, err := oss.New(env.Config.Endpoint, env.Config.AccessKey, env.Config.AccessSecret)
	if err != nil {
		fmt.Printf("初始化阿里云OSS错误：%s\n", err)
		os.Exit(1)
	}
	Client = c
}
