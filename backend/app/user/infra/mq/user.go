package mq

import (
	"encoding/json"
	"time"
)

var userSubject = "user"

type UserMsg struct {
	Method   string         `json:method`
	UpdateAt time.Time      `json:"update_at"` // 数据更新时间
	data     map[string]any `json:data`
}

func SendUserMsg(msg UserMsg) error {
	bs, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return Nc.Publish(userSubject, bs)
}
