package api

import (
	// import third-party packages
	"github.com/gin-gonic/gin"
	// import local packages
	"github.com/impossible98/butterfly/global"
	"github.com/impossible98/butterfly/pkg/ecode"
	"github.com/impossible98/butterfly/pkg/format"
)

func Version(ctx *gin.Context) {
	format.HTTP(ctx, ecode.Success, nil, gin.H{
		"version": global.VERSION,
	})
}
