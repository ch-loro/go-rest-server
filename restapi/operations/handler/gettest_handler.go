package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/tetsk/go-server/models"
	"github.com/tetsk/go-server/restapi/operations/gettest"
)

type GetTestHandler struct{}

func (h *GetTestHandler) Handle(params gettest.GettestParams) middleware.Responder {
	//add here
	payload := []*models.Obj{}

	return gettest.NewGettestOK().WithPayload(payload)
}
