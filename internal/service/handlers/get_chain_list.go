package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Swapica/aggregator-svc/internal/service/helpers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetChainList(w http.ResponseWriter, r *http.Request) {
	nodes, err := helpers.Noder(r).Select()
	if err != nil {
		helpers.Log(r).Error("failed to get node list")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	actionTarget := r.Context().Value("actionTarget").(string)
	endpoint := fmt.Sprintf("%v/%v", nodes[0], actionTarget)

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

	_, err = w.Write(body)
	if err != nil {
		helpers.Log(r).Error("failed to write to response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
