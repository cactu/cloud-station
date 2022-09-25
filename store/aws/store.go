package aws

import (
	"fmt"
)

type AwsOssStore struct {
}

func (s *AwsOssStore) Upload(bucketName, objectKet, fileName string) error {
	return fmt.Errorf("当前不支持aws上传")
}

func NewAwsOssStore() (*AwsOssStore, error) {
	return &AwsOssStore{}, nil
}
