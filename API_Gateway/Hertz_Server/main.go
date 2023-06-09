package main

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// the array to test interactions btwn cliets and server
type Numbers struct {
	Numbers []int `json:"numbers"`
}

var numbers Numbers

func createGetHandler(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, utils.H{"numbers": numbers})
}

func createPostHandler(c context.Context, ctx *app.RequestContext) {
	var reqBody struct {
		Number int `json:"number"`
	}
	// Decode the request body into the reqBody struct
	json.Unmarshal(ctx.Request.Body(), &reqBody)
	// Update the internal data structure with the new number
	numbers.Numbers = append(numbers.Numbers, reqBody.Number)
	// Send a JSON response indicating success
	ctx.JSON(consts.StatusOK, utils.H{"message": "Number added successfully"})
}

func main() {
	h := server.Default(server.WithHostPorts(":8080"))

	// init
	numbers = Numbers{Numbers: []int{1, 2, 3}}

	h.GET("/numbers", createGetHandler)

	h.POST("/numbers", createPostHandler)

	h.GET("/internalRedirect", func(ctx context.Context, c *app.RequestContext) {
		c.Redirect(consts.StatusFound, []byte("/foo"))
	})

	h.GET("/foo", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "hello, world")
	})

	h.Spin()
}
