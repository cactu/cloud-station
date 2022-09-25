package tx

import "fmt"

type TxOssStore struct {
}

func (s *TxOssStore) Upload(bucketName, objectKet, fileName string) error {
	return fmt.Errorf("当前不支持tx上传")
}

func NewTxOssStore() (*TxOssStore, error) {
	return &TxOssStore{}, nil
}
