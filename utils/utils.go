package utils

import (
	"strings"
	"encoding/json"
)

// string of tag to map
// from string (a=b&c=d) to map[string]string
func TagMap(s string) (map[string]string) {
	var m map[string]string
	var ss []string

	m = make(map[string]string)

	if s == "" {
		return m
	}

	ss = strings.Split(s, "&")
	for _, pair := range ss {
		z := strings.Split(pair, "=")

		if len(z) == 2 {
			m[z[0]] = z[1]
		}
	}
	return m
}

// Marshall Put Request Message into string
func MarshalMsg(msg interface{}) string {
	marshalMsg, err := json.Marshal(msg)
	if err != nil {
		Log.Panicf(ErrTemplate(), GetFunctionName(MarshalMsg), err)
		return ""
	}
	return string(marshalMsg)
}
