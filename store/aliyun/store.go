package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cactu/cloud-station/env"
	"github.com/schollz/progressbar/v3"
)

type AliOssParams struct {
	OssEndPoint  string
	AccessKey    string
	AccessSecret string
}

type AliOssStore struct {
	client   *oss.Client
	listener *Listener
}

type Listener struct {
	bar *progressbar.ProgressBar
}

func (l *Listener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("文件开始上传\n")
		l.bar = progressbar.DefaultBytes(event.TotalBytes)
	case oss.TransferDataEvent:
		err := l.bar.Add64(event.RwBytes)
		if err != nil {
			fmt.Printf("文件上传失败,%s\n", err)
		}
	case oss.TransferCompletedEvent:
		fmt.Printf("上传完成\n")
	case oss.TransferFailedEvent:
		fmt.Printf("上传失败\n")
	}
}

func NewProgressListener() *Listener {
	return &Listener{}
}

func (params *AliOssParams) Validate() error {
	if params.OssEndPoint == "" || params.AccessKey == "" || params.AccessSecret == "" {
		return fmt.Errorf("endpoint access_key access_secret has one is empty")
	}
	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	params := &AliOssParams{
		env.Config.Endpoint,
		env.Config.AccessKey,
		env.Config.AccessSecret,
	}
	return NewAliOssStore(params)
}

func NewAliOssStore(params *AliOssParams) (*AliOssStore, error) {
	err := params.Validate()
	if err != nil {
		return nil, err
	}
	//初始化oss链接
	client, err := oss.New(params.OssEndPoint, params.AccessKey, params.AccessSecret)
	if err != nil {
		return nil, err
	}
	store := &AliOssStore{
		client:   client,
		listener: NewProgressListener(),
	}
	return store, nil
}

func (s *AliOssStore) Upload(bucketName, objectKet, fileName string) error {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(objectKet, fileName, oss.Progress(s.listener))
	if err != nil {
		return err
	}
	downloadUrl, err := bucket.SignURL(objectKet, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("当前文件%s的下载链接为%s\n", fileName, downloadUrl)
	fmt.Println("请在一天内下载，过期失效")
	return nil
}
