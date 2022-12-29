package handlers

import (
	"bytes"
	"fmt"
	"github.com/Swapica/aggregator-svc/resources"
	"github.com/google/jsonapi"
	"net/http"

	"github.com/Swapica/aggregator-svc/internal/service/helpers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	nodes, err := helpers.Noder(r).Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get node list")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	var tx *resources.EvmTransaction
	var apiErr *jsonapi.ErrorObject

	for i := 0; i < len(nodes); i++ {
		body, err := helpers.AppendTxToBody(r, tx)
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to append tx to body")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		newTx, status, err := helpers.SendRequest(bytes.NewBuffer(body), fmt.Sprintf("%v%v", nodes[i], r.RequestURI))
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to send request")
			continue
		}
		if status == 400 {
			helpers.Log(r).Error("validation failed")
			apiErr = &jsonapi.ErrorObject{
				Title:  http.StatusText(http.StatusBadRequest),
				Status: fmt.Sprintf("%d", http.StatusBadRequest),
				Detail: "Validation failed",
				Code:   "validation_failed",
			}
			continue
		}

		if newTx == nil {
			helpers.Log(r).Error("failed to build transaction")
			continue
		}

		tx = newTx

		if *tx.Attributes.Confirmed {
			break
		}
	}

	if !*tx.Attributes.Confirmed {
		if apiErr != nil {
			ape.RenderErr(w, apiErr)
			return
		}

		helpers.Log(r).Error("not enough nodes presented")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusInternalServerError),
			Status: fmt.Sprintf("%d", http.StatusInternalServerError),
			Code:   "not_enough_active_validators",
		})
		return
	}

	ape.Render(w, tx)
}
