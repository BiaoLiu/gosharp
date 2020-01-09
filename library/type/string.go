package utils

import (
	"gosharp/library/conf/viper"
	"gosharp/library/snowflake"
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
	if endTimeValue != "" {
		endTime, err := time.Parse(DATE_FORMAT, endTimeValue)
		if err != nil {
			panic("time value error:" + err.Error())
		}
		endTime = endTime.AddDate(0, 0, 1)
		return endTime.Format(DATE_FORMAT)
	}
	return endTimeValue
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

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func FormatTime(val interface{}) string {
	switch t := val.(type) {
	case *time.Time:
		if t == nil {
			return ""
		}
		return t.Format(TIME_FORMAT)
	case time.Time:
		return t.Format(TIME_FORMAT)
	default:
		return ""
	}
}

func GenerateId() int64 {
	node, _ := snowflake.NewNode(0)
	snowId := node.Generate()
	return snowId.Int64()
}

func ConvertInt64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}

func ConvertFloatToString(value float64) string {
	return strconv.FormatFloat(value, 'E', -1, 64)
}
