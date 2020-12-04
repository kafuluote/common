package errcode

import "net/http"

const MAX_HTTP_CODE = 512

type Ecoder interface {
	SetStatus(c int)
	SetData(h H)
	ServerJson(encoding ...bool)
}

func Process(c Ecoder, e PublicErrors) {
	code := e.GetCode()

	if code > MAX_HTTP_CODE {
		c.SetStatus(http.StatusOK)
	} else {
		c.SetStatus(code)
	}
	c.SetData(e.GetResult())
	c.ServerJson()
}
