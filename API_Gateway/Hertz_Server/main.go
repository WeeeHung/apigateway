package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func createGetHandler(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, utils.H{"message": "Got it"})
}

func main() {
	h := server.Default(server.WithHostPorts(":8080"))

	h.GET("/pingz", createGetHandler)

	h.GET("/externalRedirect", func(ctx context.Context, c *app.RequestContext) {
		c.Redirect(consts.StatusMovedPermanently, []byte("http://www.google.com/"))
	})

	h.GET("/internalRedirect", func(ctx context.Context, c *app.RequestContext) {
		c.Redirect(consts.StatusFound, []byte("/foo"))
	})

	h.GET("/foo", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hello, world")
	})

	h.Spin()
}
