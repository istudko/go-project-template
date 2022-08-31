package service

import (
	"context"
	"github.com/istudko/go-project-template/contract"
	"github.com/istudko/go-project-template/db"
	"github.com/istudko/go-project-template/errors"
)

const (
	CreateTicketQuery = `INSERT INTO ticket(name) VALUES(:name) RETURNING id, name;`
)

func CreateTicket(ctx context.Context, req *contract.CreateTicketRequest) (*contract.CreateTicketResponse, error) {
	var resp contract.CreateTicketResponse
	stmt, err := db.Get().PrepareNamedContext(ctx, CreateTicketQuery)
	if err != nil {
		return nil, errors.NewDBErrorf("error create ticket, %+v", err)
	}
	if err := stmt.QueryRowxContext(ctx, req).StructScan(&resp); err != nil {
		return nil, errors.NewDBErrorf("error create ticket, %+v", err)
	}
	return &resp, nil
}
