package tools

import "time"

func TimeNowString() string {
	return time.Now().UTC().Format("2015-01-02 15:04:05")
}