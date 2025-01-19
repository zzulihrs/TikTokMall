package utils

import (
	"github.com/cloudwego/kitex/pkg/klog"
)

func MustHandleError(err error) {
	if err != nil {
		klog.Fatal(err)
	}
}
