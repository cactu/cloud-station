package cli

import (
	"fmt"
	"github.com/cactu/cloud-station/env"
	"github.com/cactu/cloud-station/store"
	"github.com/cactu/cloud-station/store/aliyun"
	"github.com/cactu/cloud-station/store/aws"
	"github.com/cactu/cloud-station/store/tx"
	"github.com/spf13/cobra"
)

var (
	uploader     store.Uploader
	provider     string
	endpoint     string
	accessKey    string
	accessSecret string
	bucketName   string
	fileName     string
	err          error
)

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Long:  "upload 文件上传",
	Short: "upload 文件上传",
	RunE: func(cmd *cobra.Command, args []string) error {
		//默认值
		defaultValue()
		switch provider {
		case "aliyun":
			uploadParams := &aliyun.AliOssParams{
				OssEndPoint:  endpoint,
				AccessKey:    accessKey,
				AccessSecret: accessSecret,
			}
			uploader, err = aliyun.NewAliOssStore(uploadParams)
			if err != nil {
				return err
			}
		case "aws":
			uploader, err = aws.NewAwsOssStore()
			if err != nil {
				return err
			}
		case "tx":
			uploader, err = tx.NewTxOssStore()
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("not support oss storage provider")
		}
		if err != nil {
			return err
		}
		return uploader.Upload(bucketName, fileName, fileName)
	},
}

func defaultValue() {
	if accessKey == "" {
		accessKey = env.Config.AccessKey
	}
	if accessSecret == "" {
		accessSecret = env.Config.AccessSecret
	}
}

func init() {
	upload := UploadCmd.PersistentFlags()
	upload.StringVarP(&provider, "oss_provider", "p", env.Config.Provider, "the oss provider")
	upload.StringVarP(&endpoint, "oss_endpoint", "e", env.Config.Endpoint, "the ali oss endpoint")
	upload.StringVarP(&bucketName, "oss_bucket_name", "b", env.Config.BucketName, "the ali oss bucket name")
	upload.StringVarP(&accessKey, "oss_access_key", "k", "", "the ali oss access key")
	upload.StringVarP(&accessSecret, "oss_access_secret", "s", "", "the ali oss access secret")
	upload.StringVarP(&fileName, "oss_upload_file", "f", "", "the ali oss upload file")
	RootCmd.AddCommand(UploadCmd)
}
