package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Swapica/aggregator-svc/internal/service/helpers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	nodes, err := helpers.Noder(r).Select()
	if err != nil {
		helpers.Log(r).Error("failed to get node list")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	endpoint := fmt.Sprintf("%v%v", nodes[0], r.RequestURI)

	res, err := http.Get(endpoint)
	if err != nil {
		helpers.Log(r).Error("failed to send request, endpoint: " + endpoint)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		helpers.Log(r).Error("failed to read response body")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		helpers.Log(r).Error("failed to write to response")
		ape.RenderErr(w, problems.InternalError())
		return
	}
}
