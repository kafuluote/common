package perror

import (
	"github.com/gin-gonic/gin"
	"log"
)

var errMsgDict = make(map[int32]string)

func RegisterErrMsgDict(dict map[int32]string) {
	for code, errMsg := range dict {
		if errMsgDict[code] != "" {
			log.Fatalf("错误码初始化错误,重复定义的code:%d", code)
		}

		errMsgDict[code] = errMsg
	}
}

func GetErrMsg(code int32) string {
	return errMsgDict[code]
}

type CommonErr struct {
	ErrCode int32  `json:"err_code"` // 业务错误码
	Msg     string `json:"msg"`      // 错误消息
}

func NewCommonErr(ret int32) error {
	return &CommonErr{
		ErrCode: ret,
		Msg:     GetErrMsg(ret),
	}

}

func (s CommonErr) Error() string {
	return s.Msg
}

var ERR_CODE_RET = "code"
var ERR_CODE_MESSAGE = "msg"
var RET_DATA = "data"

type PublicErrorType struct {
	ret  gin.H
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
	var ret = gin.H{}
	ret[ERR_CODE_RET] = 0
	ret[ERR_CODE_MESSAGE] = 0
	s.ret = ret
	s.data = make(map[string]interface{}, 0)
}

func (s *PublicErrorType) GetErrCode() int32 {
	g, ok := s.ret[ERR_CODE_RET]
	if ok {
		e, ok := g.(int32)
		if ok {
			return e
		}

		r, ok := g.(int)
		if ok {
			return int32(r)
		}
		/*
			k:=reflect.TypeOf(g)
			switch k.(type) {
			case int:
				return g.(int)
			case int32:
				return g.(int32)
			default:
				panic(k)
			}
		*/

		panic("err type")
	}
	return 0
}

//设置错误代码，如果有自定义错误信息填写err_msg参数
func (s *PublicErrorType) SetErrCode(code int32, err_msg ...string) {
	s.ret[ERR_CODE_RET] = code

	if len(err_msg) > 0 {
		if err_msg[0] == "" {
			s.ret[ERR_CODE_MESSAGE] = GetErrMsg(code)
		} else {
			s.ret[ERR_CODE_MESSAGE] = err_msg[0]
		}
	} else {
		s.ret[ERR_CODE_MESSAGE] = GetErrMsg(code)
	}
}

//设置数据部分内容
func (s *PublicErrorType) SetDataSection(key string, value interface{}) {
	s.data[key] = value
}

//返回最终的数据
func (s *PublicErrorType) GetResult() gin.H {
	s.ret[RET_DATA] = s.data
	return s.ret
}

// 补充-设置数据部分内容
func (s *PublicErrorType) SetDataValue(value interface{}) {
	s.ret[RET_DATA] = value
}

// 补充-返回最终的数据
func (s *PublicErrorType) GetData() gin.H {
	return s.ret
}
