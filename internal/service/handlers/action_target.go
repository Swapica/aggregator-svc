package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/Swapica/aggregator-svc/internal/service/helpers"
	"github.com/Swapica/swapica-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func ActionTarget(w http.ResponseWriter, r *http.Request) {
	nodes, err := helpers.Noder(r).Select()
	if err != nil {
		helpers.Log(r).Error("failed to get node list")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var tx *resources.EvmTransaction
	actionTarget := r.Context().Value("actionTarget").(string)

	for i := 0; i < len(nodes); i++ {
		body, err := helpers.AppendTxToBody(r, tx, actionTarget)
		if err != nil {
			helpers.Log(r).Error("failed to append tx to body")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		tx, err = helpers.SendRequest(bytes.NewBuffer(body), fmt.Sprintf("%v/%v", nodes[i], actionTarget))
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

		if *tx.Attributes.Confirmed {
			break
		}
	}

	if !*tx.Attributes.Confirmed {
		helpers.Log(r).Error("not enough nodes presented")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, tx)
}
