package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *oss.Client

var (
	AccessKey       = os.Getenv("ALI_AK")
	AccessKeySecret = os.Getenv("ALI_AS")
	OssEndpoint     = os.Getenv("ALI_ENDPOINT")
	BucketName      = os.Getenv("ALI_BUCKET_NAME")
)

func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log("ERROR:", err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("-----Buckets:", bucket.Name)
	}
}

func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("mydir/test.go", "oss_test.go")
	if err != nil {
		t.Log(err)
	}
}

// 初始化一个Oss Client
func init() {
	fmt.Println("++++++OssEndpoint:", OssEndpoint, AccessKey, AccessKeySecret)
	c, err := oss.New(OssEndpoint, AccessKey, AccessKeySecret)
	if err != nil {
		panic(err)
	}
	client = c
}
