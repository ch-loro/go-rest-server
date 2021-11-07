package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/tetsk/go-server/restapi/operations/posttest"
)

type PostTestHandler struct{}

func (h *PostTestHandler) Handle(params posttest.PosttestParams) middleware.Responder {
	//add here

	return posttest.NewPosttestCreated()
}
