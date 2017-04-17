package main

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"strings"
	//"time"
)

var checkmap = map[string]oss.Option{
	"fax":  oss.ContentType("image/fax"),
	"gif":  oss.ContentType("image/gif"),
	"ico":  oss.ContentType("image/x-icon"),
	"jfif": oss.ContentType("image/jpeg"),
	"jpe":  oss.ContentType("image/jpeg"),
	"jpeg": oss.ContentType("image/jpeg"),
	"jpg":  oss.ContentType("image/jpeg"),
	"net":  oss.ContentType("image/pnetvue"),
	"png":  oss.ContentType("image/png"),
	"rp":   oss.ContentType("image/vnd.rn-realpix"),
	"tif":  oss.ContentType("image/tiff"),
	"tiff": oss.ContentType("image/tiff"),
	"wbmp": oss.ContentType("image/vnd.wap.wbmp"),
}

type Client struct {
	Bucket *oss.Bucket
}

func NewClient(config *Config) (Client, error) {
	ACCESS_ID := config.ACCESS_ID
	ACCESS_SEC_KEY := config.ACCESS_SEC_KEY
	endpoint := config.Endpoint
	bucketname := config.Bucket

	client, err := oss.New(endpoint, ACCESS_ID, ACCESS_SEC_KEY)
	if err != nil {
		return Client{}, errors.New("init aliyun client error")
	}
	bucket, err := client.Bucket(bucketname)
	if err != nil {
		return Client{}, errors.New("get aliyun oss bucket error")
	}
	return Client{
		Bucket: bucket,
	}, nil
}

func ReturnFilenameExtension(name string) string {
	tmp := strings.Split(name, ".")
	return tmp[len(tmp)-1]
}

func (cli *Client) ChangeContentType(path string) error {
	option := checkmap[ReturnFilenameExtension(path)]
	return cli.Bucket.SetObjectMeta(path, option)
}

func (cli *Client) GenerateTask(path string, config *Config) (Task, error) {
	lsRes, err := cli.Bucket.ListObjects(oss.Prefix(path))
	if err != nil {
		return Task{}, err
	}

	var taskslice []string
	for _, object := range lsRes.Objects {
		taskslice = append(taskslice, object.Key)
	}
	return Task{Keys: taskslice}, nil
}
