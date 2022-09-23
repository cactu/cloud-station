package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cactu/cloud-station/config"
	"os"
)

var filePath = ""
var help = false

//文件上传
func uploadFile(filePath string) {
	bucket, err := config.GetClientBucket()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = bucket.PutObjectFromFile(filePath, filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	downloadUrl, err := bucket.SignURL(filePath, oss.HTTPGet, 60*60*24)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("上传文件 %s 成功\n", filePath)
	fmt.Printf("下载链接是:【%s】", downloadUrl)
}

//参数校验
func validate() error {
	if config.AccessKey == "" || config.AccessSecret == "" {
		return fmt.Errorf("AccessKey or AccessSecret missed")
	}
	if config.OssEndPoint == "" || config.BucketName == "" {
		return fmt.Errorf("OssEndPoint or BucketName missed")
	}
	return nil
}

//加载用户输入参数
func loadParam() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.StringVar(&filePath, "f", "", "upload file")
	flag.Parse()
	if help {
		usage()
		os.Exit(1)
	}
	if filePath == "" {
		usage()
		os.Exit(1)
	}
}

//打印帮助信息
func usage() {
	_, err := fmt.Fprintf(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <uplaod_file_path>
Options:`)
	if err != nil {
		fmt.Printf("使用说明异常:%s", err)
	}
	flag.PrintDefaults()
}

func main() {
	//校验参数
	err := validate()
	if err != nil {
		fmt.Printf("validate params error : %s\n", err)
		os.Exit(1)
	}
	//接受用户的命令行参数
	loadParam()
	if filePath == "" {
		fmt.Println("未指定需要上传的文件")
		os.Exit(1)
	}

	uploadFile(filePath)
}
