// Package util 时间工具类
package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type HTime struct {
	time.Time
}

var (
	formatTime = "2006-01-02 15:04:05"
)

// MarshalJSON 统一改为指针接收器
func (t *HTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	return []byte(formatted), nil
}

func (t *HTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"`+formatTime+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*t = HTime{Time: now}
	return nil
}

// Value 修改为值接收器，避免空指针问题
func (t HTime) Value() (driver.Value, error) {
	// 检查是否为零值
	if t.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *HTime) Scan(v interface{}) error {
	if v == nil {
		*t = HTime{Time: time.Time{}}
		return nil
	}
	value, ok := v.(time.Time)
	if ok {
		*t = HTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
