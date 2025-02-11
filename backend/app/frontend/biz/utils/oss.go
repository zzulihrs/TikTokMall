package utils

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/retry"
)

// 配置
const (
	BucketName = "titktokmall"                                     // 替换为您的Bucket名称
	Region     = "cn-beijing"                                      // 替换为您的Bucket地域（如cn-hangzhou）
	Endpoint   = "https://titktokmall.oss-cn-beijing.aliyuncs.com" // 替换为您的Endpoint
)

var once sync.Once
var client *oss.Client
var clientErr error

// 上传结果结构
type UploadResult struct {
	HashedName string // 哈希后文件名
	URL        string // 完整访问URL
}

// OssUtils 工具类
type OssUtils struct {
	client *oss.Client // OSS 客户端
}

// 新建 OSS 工具类实例
func NewOssUtils() (*OssUtils, error) {
	// 确保客户端只初始化一次
	once.Do(func() {
		fmt.Println("newOSSClient1")
		client, clientErr = newOSSClient()
	})

	if clientErr != nil {
		return nil, clientErr
	}

	return &OssUtils{client: client}, nil
}

// 生成随机哈希文件名（不保留扩展名）
func generateHashedName() (string, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("生成随机数失败: %v", err)
	}

	timestamp := time.Now().Unix()
	return fmt.Sprintf("%d_%s", timestamp, hex.EncodeToString(randomBytes)), nil
}

// 上传图片到 OSS
func (o *OssUtils) UploadImage(ctx context.Context, buffer bytes.Buffer, fileName string) (*UploadResult, error) {
	ext := filepath.Ext(fileName)
	contentType := "image/jpeg"
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	case ".webp":
		contentType = "image/webp"
	}
	return o.uploadFile(ctx, buffer, contentType, "images/")
}

// 上传视频到 OSS
func (o *OssUtils) UploadVideo(ctx context.Context, buffer bytes.Buffer, fileName string) (*UploadResult, error) {
	ext := filepath.Ext(fileName)
	contentType := "video/mp4"
	switch ext {
	case ".mov":
		contentType = "video/quicktime"
	case ".avi":
		contentType = "video/x-msvideo"
	case ".webm":
		contentType = "video/webm"
	}
	return o.uploadFile(ctx, buffer, contentType, "videos/")
}

// 通用上传方法
func (o *OssUtils) uploadFile(ctx context.Context, file bytes.Buffer, contentType, prefix string) (*UploadResult, error) {
	// 生成哈希文件名
	hashedName, err := generateHashedName()
	if err != nil {
		return nil, err
	}

	// 设置对象键
	objectKey := prefix + hashedName

	// 构造上传请求
	request := &oss.PutObjectRequest{
		Bucket:      oss.Ptr(BucketName),
		Key:         oss.Ptr(objectKey),
		Body:        &file,
		ContentType: oss.Ptr(contentType),
	}

	// 执行上传
	_, err = o.client.PutObject(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("上传失败: %v", err)
	}

	// 生成访问 URL
	url := fmt.Sprintf("%s/%s", Endpoint, objectKey)
	log.Printf("文件上传成功！URL: %s\n", url)

	return &UploadResult{
		HashedName: hashedName,
		URL:        url,
	}, nil
}

// 初始化 OSS 客户端
func newOSSClient() (*oss.Client, error) {
	// 自定义重试器
	customRetriever := retry.NewStandard(func(so *retry.RetryOptions) {
		so.MaxAttempts = 3
		so.MaxBackoff = 5 * time.Second
	})

	// 配置凭证提供者
	credProvider := credentials.NewEnvironmentVariableCredentialsProvider()

	// 加载默认配置并设置区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credProvider).
		WithRegion(Region).
		WithRetryer(customRetriever)

	// 创建客户端实例
	client := oss.NewClient(cfg)
	return client, nil
}

// 获取 OSS 客户端
func (o *OssUtils) GetClient() *oss.Client {
	return o.client
}
