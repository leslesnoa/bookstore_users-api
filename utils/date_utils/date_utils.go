package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2019-01-02T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()

}

func GetNowString() string {
	// 現在時間を取得し、DBのDatecreatadeパラメータに挿入
	return GetNow().Format(apiDateLayout)
}
