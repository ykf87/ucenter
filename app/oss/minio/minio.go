package minio

import (
	"context"
	"errors"
	"os"
	"strings"
	"ucenter/app/config"
	"ucenter/app/funcs"
	"ucenter/app/logs"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	Conf  *config.OssConf
	Minio *minio.Client
}

var MailObj *Client

func GetClient() (*Client, error) {
	if MailObj != nil {
		return MailObj, nil
	}
	conf, ok := config.Config.Oss[config.Config.Useoss]
	if !ok {
		return nil, errors.New("配置信息不正确,找不到oss的配置: " + config.Config.Useoss)
	}

	minioClient, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.Accesskeyid, conf.Secret, ""),
		Secure: conf.Ssl,
	})
	if err != nil {
		return nil, err
	}
	// err = minioClient.MakeBucket(context.Background(), conf.Bucket, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: false})
	// if err != nil {
	// 	logs.Logger.Error(err)
	// 	return nil, err
	// }
	cc := new(Client)
	cc.Conf = conf
	cc.Minio = minioClient
	MailObj = cc
	return cc, nil
}

//获取桶名
func (this *Client) BluckName() string {
	return this.Conf.Bucket
}

//objectName 云端保存的路径和文件名
func (this *Client) Upload(filePath, objectName string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	contentType, err := funcs.GetFileContentType(f)
	if err != nil {
		return "", err
	}

	// Upload the zip file with FPutObject
	ctx := context.Background()
	// info, err := this.Minio.PutObject(ctx, this.Conf.Bucket, objectName, f, fileInfo.Size(), minio.PutObjectOptions{ContentType: contentType})
	info, err := this.Minio.FPutObject(ctx, this.Conf.Bucket, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logs.Logger.Error("上传云端失败: ", err.Error(), this.Conf.Bucket)
		return "", err
	}

	return info.Bucket + "/" + info.Key, nil
}

func (this *Client) Remove(src string) error {
	opts := minio.RemoveObjectOptions{
		ForceDelete:      true,
		GovernanceBypass: true,
	}
	src = strings.Replace(src, this.Conf.Bucket+"/", "", -1)
	return this.Minio.RemoveObject(context.Background(), this.Conf.Bucket, src, opts)
}

func (this *Client) Url(filename string) string {
	url := this.Conf.Url
	if url == "" {
		sheme := "http://"
		if this.Conf.Ssl == true {
			sheme = "https://"
		}
		url = sheme + this.Conf.Endpoint
	}
	return strings.TrimRight(url, "/") + "/" + filename
}
