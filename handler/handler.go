package handler

import (
	"encoding/json"
	"github.com/istudko/go-project-template/contract"
	"github.com/istudko/go-project-template/logger"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, resp *contract.Response) {
	w.Header().Set("Content-Type", "application/json")

	responseJSON, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		logger.Errorf("unable to marshal request response")
		return
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = w.Write(responseJSON)
}
