package errcode

import "net/http"

const ERR_CODE = "ret"
const ERR_MSG = "msg"
const ERR_DATA = "data"

type H map[string]interface{}
type PublicErrors struct {
	ret  H
	data map[string]interface{}
}

func NewPublicErrors() *PublicErrors {
	s := &PublicErrors{}
	s.init()
	return s
}

func (s *PublicErrors) init() {
	ret := H{}
	ret[ERR_CODE] = 0
	ret[ERR_MSG] = ""
	s.ret = ret
	s.data = make(map[string]interface{}, 0)
}

func (s *PublicErrors) SetErrCode(code int, err_msg ...string) {
	s.ret[ERR_CODE] = code
	if code > MAX_HTTP_CODE {
		if len(err_msg) > 0 {
			s.ret[ERR_MSG] = err_msg[0]
		} else {
			s.ret[ERR_MSG] = GetMessage(code)
		}
	} else {
		if len(err_msg) > 0 {
			s.ret[ERR_MSG] = err_msg[0]
		} else {
			s.ret[ERR_MSG] = http.StatusText(code)
		}
	}
}

func (s *PublicErrors) SetDataSection(key string, value interface{}) {
	s.data[key] = value
}

func (s *PublicErrors) GetResult() H {
	s.ret[ERR_DATA] = s.data
	return s.ret
}

func (s *PublicErrors) GetCode() int {
	return s.ret[ERR_CODE].(int)
}
