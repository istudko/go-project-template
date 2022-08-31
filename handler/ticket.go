package handler

import (
	"github.com/istudko/go-project-template/contract"
	"github.com/istudko/go-project-template/logger"
	"github.com/istudko/go-project-template/service"
	"github.com/istudko/go-project-template/validator"
	"net/http"
)

func CreateTicketHandler(w http.ResponseWriter, r *http.Request) {
	req, err := contract.NewCreateTicketRequest(r)
	if err != nil {
		logger.Warnf("invalid payload, %+v", err)
		resp := contract.NewErrorResponse(err)
		WriteResponse(w, &resp)
		return
	}

	if err := validator.Validate(req); err != nil {
		logger.Warnf("validation error, %+v", err)
		resp := contract.NewErrorResponse(err)
		WriteResponse(w, &resp)
		return
	}
	ticketResponse, err := service.CreateTicket(r.Context(), req)
	if err != nil {
		logger.Errorf("failure to create ticket, %+v", err)
		resp := contract.NewErrorResponse(err)
		WriteResponse(w, &resp)
		return
	}
	resp := contract.NewResponse(http.StatusCreated, ticketResponse)
	WriteResponse(w, &resp)
}
