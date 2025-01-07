ceph/ceph.go

```go
// Package ceph Go对ceph s3文件增删改
package ceph

import (
	"fmt"
	"os"

	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

// MgrCephS3 ceph s3
type MgrCephS3 struct {
	EndPoint  string // ceph s3地址
	AccessKey string // ceph s3 AK
	SecretKey string // ceph s3 SK

	CephConn *s3.S3 // cehp 连接
}

// NewMgrCephS3 构造函数
func NewMgrCephS3(endPoint, accessKey, secretKey string) *MgrCephS3 {
	return &MgrCephS3{
		EndPoint:  endPoint,
		AccessKey: accessKey,
		SecretKey: secretKey,
	}
}

// InitCephS3 初始化ceph s3连接
func (ms *MgrCephS3) InitCephS3() {
	auth := aws.Auth{
		AccessKey: ms.AccessKey,
		SecretKey: ms.SecretKey,
	}

	region := aws.Region{
		Name:                 "default",
		EC2Endpoint:          ms.EndPoint,
		S3Endpoint:           ms.EndPoint,
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}

	ms.CephConn = s3.New(auth, region) // 获取ceph s3连接
}

// InitBucket 获取一个桶连接
func (ms *MgrCephS3) InitBucket(bucket string) (*MgrBucket, error) {
	if ms.CephConn == nil {
		return nil, fmt.Errorf("MgrCephS3.CephConn is nil")
	}

	return &MgrBucket{
		BucketName: bucket,
		BucketConn: ms.CephConn.Bucket(bucket),
	}, nil
}

// MgrBucket 桶
type MgrBucket struct {
	BucketName string     // ceph s3 bucket桶名
	BucketConn *s3.Bucket // 桶连接
}

// Put2Bucket 上传文件到bucket
func (mb *MgrBucket) Put2Bucket(localPath, cephPath string) error {
	// 权限
	err := mb.BucketConn.PutBucket(s3.AuthenticatedRead)
	if err != nil {
		return err
	}

	// 读取文件
	bytes, err := os.ReadFile(localPath)
	if err != nil {
		return err
	}

	// 上传文件
	return mb.BucketConn.Put(cephPath, bytes, "octet-stream", s3.AuthenticatedRead)
}

// DownloadFromBucket 从Bucket下载文件
func (mb *MgrBucket) DownloadFromBucket(localPath, cephPath string) error {
	data, err := mb.BucketConn.Get(cephPath)
	if err != nil {
		return err
	}

	return os.WriteFile(localPath, data, 0644)
}

// DelBucketData 删除ceph上指定文件
func (mb *MgrBucket) DelBucketData(cephPath string) error {
	return mb.BucketConn.Del(cephPath)
}

// GetBatchFromBucket 批量获取文件信息
func (mb *MgrBucket) GetBatchFromBucket(prefixCephPath string) ([]string, error) {
	result, err := mb.BucketConn.List(prefixCephPath, "", "", 0)
	if err != nil {
		return nil, err
	}

	keyList := make([]string, 0)
	for _, key := range result.Contents {
		keyList = append(keyList, key.Key)
	}

	return keyList, nil
}

```



main.go

```
/*
需要提供ceph s3的信息为:
ceph url
ceph AccessKey
ceph SecretKey
ceph bucket 路径（例如: des-backet/

https://osstest.alnk.com
BSETGVM88JBHNNHPALNK
HQo8eER1Drc9LCEatEZqUfZEPJKTtPCl8jjQALNk
des-bucket
*/

package main

import (
	"fmt"

	"github.com/alnk-ceph-s3/ceph"
)

func main() {
	// 1. 获取ceph s3对象
	ms3 := ceph.NewMgrCephS3("https://osstest.alnk.com", "BSETGVM88JBHNNHPALNK", "HQo8eER1Drc9LCEatEZqUfZEPJKTtPCl8jjQALNk")

	// 2. 初始化ceph s3连接
	ms3.InitCephS3()

	// 3. 初始化桶连接
	bk, err := ms3.InitBucket("des-bucket")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 上传文件
	if err := bk.Put2Bucket("./main.go", "alnk/main.go.1"); err != nil {
		fmt.Println(err)
		return
	}

	// 上传文件
	if err := bk.Put2Bucket("./main.go", "alnk/main.go.2"); err != nil {
		fmt.Println(err)
		return
	}

	// 上传文件
	if err := bk.Put2Bucket("./main.go", "alnk/main.go.3"); err != nil {
		fmt.Println(err)
		return
	}

	// 删除文件
	if err := bk.DelBucketData("alnk/main.go.2"); err != nil {
		fmt.Println(err)
		return
	}

	// 下载文件
	if err := bk.DownloadFromBucket("./main.go-ceph", "alnk/main.go.3"); err != nil {
		fmt.Println(err)
		return
	}

	// 批量获取文件
	fileSlice, err := bk.GetBatchFromBucket("alnk")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("fileSlice: ", fileSlice)  // fileSlice:  [alnk/main.go.1 alnk/main.go.3]
}

```
