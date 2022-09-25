package store

type Uploader interface {
	Upload(bucketName, objectKet, fileName string) error
}
