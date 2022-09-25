package aliyun

import (
	"fmt"
	"github.com/cactu/cloud-station/env"
	"github.com/cactu/cloud-station/store"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var uploader store.Uploader

var (
	BucketName string
)

func TestAliYunUploadFile(t *testing.T) {
	should := assert.New(t)
	err := uploader.Upload(BucketName, "test.txt", "store_test.go")
	if should.NoError(err) {
		fmt.Println("upload ok")
	}
}

func init() {
	u, err := NewDefaultAliOssStore()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	uploader = u
	BucketName = env.Config.BucketName
}
