package util

import "time"

// TimeLocation 时区
func TimeLocation() *time.Location {
	location, err := time.LoadLocation("Asia/Shanghai")
	// return strconv.FormatInt(time.Now().Unix(), 13)
	if err != nil {
		location = time.FixedZone("CST", 8*3600) //替换上海时区
	}
	return location
}
