package example

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cactu/cloud-station/env"
	"testing"
)

var client *oss.Client

func TestListBucket(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket(env.Config.BucketName)
	if err != nil {
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("oss_test.go", "oss_test.go")
	if err != nil {
		t.Log(err)
	}
}

func init() {
	//初始化oss链接
	c, err := oss.New(env.Config.Endpoint, env.Config.AccessKey, env.Config.AccessSecret)
	if err != nil {
		panic(err)
	}
	client = c
}
