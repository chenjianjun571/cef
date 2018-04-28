package cef

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSUtils struct {
	bucketName   string   // bucket名称
	baseDir      []string // 合法域名
	endpoint     string   // oss域名
	basePutPath  string   // 存放目录
	accessKeyID  string   // 密钥key
	accessSecret string   // 密钥
}

// SetBucketName 设置bucket名称
func (its *OSSUtils) SetBucketName(bucketName string) {
	its.bucketName = bucketName
}

// SetBaseDir 设置baseDir
func (its *OSSUtils) SetBaseDir(baseDirs ...string) {
	its.baseDir = baseDirs
}

// SetBasePutPath 设置basePutPath
func (its *OSSUtils) SetBasePutPath(basePutPath string) {
	its.basePutPath = basePutPath
}

// SetAccessInfo 设置
func (its *OSSUtils) SetAccessInfo(endpoint string, accessKeyID string, accessSecret string) {
	its.endpoint = endpoint
	its.accessKeyID = accessKeyID
	its.accessSecret = accessSecret
}

func (its *OSSUtils) GetOSSClient() (*oss.Client, error) {
	return oss.New(its.endpoint, its.accessKeyID, its.accessSecret)
}
