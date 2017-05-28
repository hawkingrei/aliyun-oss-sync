package main

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"strings"
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

func (cli *Client) GenerateTask(n *NSQD, config *Config) ([]string, error) {
	marker := oss.Marker("")
	for {
		select {
		case prestr := <-n.PreChan:
			pre := oss.Prefix(prestr)
			for {
				lor, err := cli.Bucket.ListObjects(oss.MaxKeys(10), marker, pre)
				if err != nil {
					log.Println(err.Error())
					return []string{}, err
				}
				pre = oss.Prefix(lor.Prefix)
				marker = oss.Marker(lor.NextMarker)
				for _, object := range lor.Objects {
					n.TaskChan <- object.Key
				}
				if !lor.IsTruncated {
					break
				}
			}
		case <-n.ExitChan:
			break
		}
	}
}

func (cli *Client) Worker(n *NSQD) {
	for {
		select {
		case key := <-n.TaskChan:
			err := cli.ChangeContentType(string(key))
			if err != nil {
				log.Println("error " + string(key) + " " + err.Error())
			} else {
				log.Println("ok " + key)
			}

		case <-n.ExitChan:
			break
		}
	}
}
