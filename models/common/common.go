package common

import "encoding/json"

func MapJson(status string, statusval interface{}, info string, infoval interface{}) string {
	m := make(map[string]interface{})
	m[status] = statusval
	m[info] = infoval
	mData, err := json.Marshal(m)
	if err != nil {
		Log.Warning(err.Error())
		return ""
	}
	return string(mData)
}
