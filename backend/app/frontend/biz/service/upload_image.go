package service

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	oss "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/oss"
)

type UploadImageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUploadImageService(Context context.Context, RequestContext *app.RequestContext) *UploadImageService {
	return &UploadImageService{RequestContext: RequestContext, Context: Context}
}

func (h *UploadImageService) Run(req *oss.UploadFileRequest) (resp *oss.UploadResponse, err error) {

	// 解析 multipart/form-data
	file, err := h.RequestContext.FormFile("file")
	if err != nil {
		return nil, err
	}

	// 打开文件
	fileReader, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileReader.Close()

	// 读取文件内容到 bytes.Buffer
	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, fileReader); err != nil {
		return nil, err
	}

	// 获取文件名
	fileName := file.Filename

	// 初始化 OSS 工具类
	ossUtils, err := utils.NewOssUtils()
	if err != nil {
		log.Printf("初始化 OSS 工具失败: %v", err)
	}
	temp, err := ossUtils.UploadImage(h.Context, buffer, fileName)
	if err != nil {
		id := os.Getenv("OSS_ACCESS_KEY_ID")
		secret := os.Getenv("OSS_ACCESS_KEY_SECRET")
		log.Printf("上传图片失败: %v", err)
		log.Println("id: ", id)
		log.Println("secret: ", secret)
		return nil, err
	}
	resp = &oss.UploadResponse{
		HashedName: temp.HashedName,
		Url:        temp.URL,
	}

	return resp, nil
}
