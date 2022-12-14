package handlers

import (
	"fmt"
	"net/http"

	"github.com/Swapica/aggregator-svc/internal/service/helpers"
	"github.com/Swapica/swapica-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func Action(w http.ResponseWriter, r *http.Request) {
	nodes, err := helpers.Noder(r).Select()
	if err != nil {
		helpers.Log(r).Error("failed to get node list")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var tx *resources.EvmTransaction
	actionTarget := r.Context().Value("actionTarget").(string)

	tx, err = helpers.SendRequest(r.Body, fmt.Sprintf("%v/%v", nodes[0], actionTarget))
	if err != nil {
		helpers.Log(r).Error("failed to send request")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if tx == nil {
		helpers.Log(r).Error("failed to build transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, tx)
}
