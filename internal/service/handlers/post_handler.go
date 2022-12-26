package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/Swapica/aggregator-svc/resources"
	"net/http"

	"github.com/Swapica/aggregator-svc/internal/service/helpers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	nodes, err := helpers.Noder(r).Select()
	if err != nil {
		helpers.Log(r).Error("failed to get node list")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var tx *resources.EvmTransaction
	var status int

	for i := 0; i < len(nodes); i++ {
		body, err := helpers.AppendTxToBody(r, tx)
		if err != nil {
			helpers.Log(r).Error("failed to append tx to body")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		tx, status, err = helpers.SendRequest(bytes.NewBuffer(body), fmt.Sprintf("%v%v", nodes[i], r.RequestURI))
		if status == 400 {
			helpers.Log(r).Error("validation failed")
			ape.RenderErr(w, problems.BadRequest(errors.New("validation failed"))...)
			return
		}
		if err != nil {
			helpers.Log(r).Error("failed to send request")
			ape.RenderErr(w, problems.BadRequest(err)...)
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
