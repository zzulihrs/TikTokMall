package service

import (
	"bytes"
	"context"
	"github.com/tiktokmall/backend/app/frontend/biz/utils"
	"io"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	oss "github.com/tiktokmall/backend/app/frontend/hertz_gen/frontend/oss"
)

type UploadVideoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUploadVideoService(Context context.Context, RequestContext *app.RequestContext) *UploadVideoService {
	return &UploadVideoService{RequestContext: RequestContext, Context: Context}
}

func (h *UploadVideoService) Run(req *oss.UploadFileRequest) (resp *oss.UploadResponse, err error) {
	/// 解析 multipart/form-data
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
		log.Fatalf("初始化 OSS 工具失败: %v", err)
	}
	temp, err := ossUtils.UploadVideo(h.Context, buffer, fileName)
	if err != nil {
		log.Fatalf("上传视频失败: %v", err)
		return nil, err
	}
	resp = &oss.UploadResponse{
		HashedName: temp.HashedName,
		Url:        temp.URL,
	}

	return resp, nil
}
