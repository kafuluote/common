package errcode

import (
	"fmt"
)

var message map[int]string

type ErrData struct {
	Code    int
	Message string
}

const (
	ERRCODE_SYS_UNKONWN = 0
)

var code = []ErrData{
	{ERRCODE_SYS_UNKONWN, "unkonwn"},
}

func insertErr(code int, mes string) {
	_, ok := message[code]
	if ok {
		panic(fmt.Sprintf("repeated code %d", code))
	}
	message[code] = mes
}

func LoadCode(d []ErrData) {
	for _, v := range d {
		insertErr(v.Code, v.Message)
	}
}
func GetMessage(code int) string {
	d, ok := message[code]
	if ok {
		return d
	}
	return ""
}

func init() {
	message = make(map[int]string)
	LoadCode(code)
}
