package response

import (
	"github.com/gin-gonic/gin"
	"supermarket/resources/lang"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// Response conf gin.JSON
func (g *Gin) Res(httpCode int, status, msg string, data interface{}) {
	// response json
	g.C.JSON(httpCode, Response{
		Status: status,
		Msg:    lang.GetMsg(msg),
		Data:   data,
	})
	return
}

func (g *Gin) Success(msg string, data interface{}) {
	g.Res(200, "success", msg, data)
}

func (g *Gin) Error(msg string, data interface{}) {
	g.Res(200, "error", msg, data)
}
