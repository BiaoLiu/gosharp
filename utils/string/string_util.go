package string_util

import (
	"gosharp/utils/config"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
	DATE_FORMAT = "2006-01-02"
)

func SafeInt(value string, defaultValue int) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func CastEndTime(endTimeValue string) string {
	endTime, err := time.Parse(DATE_FORMAT, endTimeValue)
	if err != nil {
		panic("时间类型格式错误:" + err.Error())
	}
	endTime = endTime.AddDate(0, 0, 1)
	return endTime.Format(DATE_FORMAT)
}

func FormatUrl(value string) string {
	if value != "" {
		return config.Viper.GetString("QINIU_HOST") + "/" + value
	}
	return value
}

func ParseUrl(rawUrl string) string {
	if u, err := url.Parse(rawUrl); err == nil {
		return strings.TrimLeft(u.Path, "/")
	}
	return ""
}

func CastMap(m map[interface{}]interface{}) map[string]interface{} {
	m2 := make(map[string]interface{})
	for key, value := range m {
		switch key := key.(type) {
		case string:
			m2[key] = value
		}
	}
	return m2
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
