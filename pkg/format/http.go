package format

import (
	// import built-in packages
	"fmt"
	"net/http"
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"github.com/impossible98/butterfly/pkg/ecode"
)

func HTTP(c *gin.Context, code int, err error, data interface{}) {
	type resp struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data,omitempty"`
	}

	var msg = fmt.Sprintf("%s", ecode.GetMsg(code))
	if err != nil {
		msg += fmt.Sprintf(": %s", err.Error())
	}
	c.JSON(http.StatusOK, &resp{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
