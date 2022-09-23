package example

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var client *oss.Client

var (
	AccessKey    string
	AccessSecret string
	OssEndPoint  string
	BucketName   string
)

func TestListBucket(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
		fmt.Println()
	}
}

func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		t.Log(err)
	}
}

func init() {
	//初始化配置文件
	err := godotenv.Load("../etc/test.env")
	if err != nil {
		panic(err)
	}
	AccessKey = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndPoint = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName = os.Getenv("ALI_BUCKET_NAME")
	//初始化oss链接
	c, err := oss.New(OssEndPoint, AccessKey, AccessSecret)
	if err != nil {
		panic(err)
	}
	client = c
}
