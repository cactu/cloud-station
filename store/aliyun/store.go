package aliyun

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cactu/cloud-station/env"
)

type AliOssParams struct {
	OssEndPoint  string
	AccessKey    string
	AccessSecret string
}

type AliOssStore struct {
	client *oss.Client
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
		client: client,
	}
	return store, nil
}

func (s *AliOssStore) Upload(bucketName, objectKet, fileName string) error {
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObjectFromFile(objectKet, fileName)
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
