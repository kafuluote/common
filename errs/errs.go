package errs

import "net/http"

var ERR_CODE_RET = "code"
var ERR_CODE_MESSAGE = "msg"
var RET_DATA = "data"

var message map[int]string

const (
	ERRCODE_SUCCESS = 0
	ERRCODE_UNKNOWN = 1
	ERRCODE_PARAM   = 2
)
type H map[string]interface{}

func init() {
	message = make(map[int]string, 0)
	message[ERRCODE_SUCCESS] = "成功"
	message[ERRCODE_UNKNOWN] = "未知错误"
	message[ERRCODE_PARAM] = "参数错误"

}


func getErrorMessage(code int) string {
	return message[code]
}

type PublicErrorType struct {
	ret  H
	data map[string]interface{}
}

//创建统一错误返回格式
func NewPublciError() *PublicErrorType {
	s := new(PublicErrorType)
	s.init()
	return s
}

//初始化操作
func (s *PublicErrorType) init() {
	var ret = H{}
	ret[ERR_CODE_RET] = 0
	ret[ERR_CODE_MESSAGE] = 0
	s.ret = ret
	s.data = make(map[string]interface{}, 0)
}

//设置错误代码，如果有自定义错误信息填写err_msg参数
func (s *PublicErrorType) SetErrCode(code int, err_msg ...string) {
	s.ret[ERR_CODE_RET] = code

	if len(err_msg)>0 {
		s.ret[ERR_CODE_MESSAGE] = err_msg[0]
	}else{
		if code>512 {
			s.ret[ERR_CODE_MESSAGE]=getErrorMessage(code)
		}else{
			s.ret[ERR_CODE_MESSAGE]=http.StatusText(code)
		}

	}
}

//设置数据部分内容
func (s *PublicErrorType) SetDataSection(key string, value interface{}) {
	s.data[key] = value
}

//返回最终的数据
func (s *PublicErrorType) GetResult() H {
	s.ret[RET_DATA] = s.data
	return s.ret
}

// 补充-设置数据部分内容
func (s *PublicErrorType) SetDataValue(value interface{}) {
	s.ret[RET_DATA] = value
}

/*
// 补充-返回最终的数据
func (s *PublicErrorType) GetData() gin.H {
	return s.ret
}
*/
