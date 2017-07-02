package common

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"fmt"
)

//响应json
func ResponseJson(status interface{}, info interface{}) string {
	m := make(map[string]interface{})
	m["status"] = status
	m["info"] = info
	mData, err := json.Marshal(m)
	if err != nil {
		Log.Warning(err.Error())
		return ""
	}
	return string(mData)
}
func Md5(s string) string {
	hash := md5.New()
	buf := []byte(s)
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func Sha1(s string) string {
	hash := sha1.New()
	buf := []byte(s)
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//Sha1 Plus Md5
func Sha1PlusMd5(s string) string {
	return Sha1(Md5(s))
}
