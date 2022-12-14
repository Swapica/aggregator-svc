package requests

import (
	"encoding/json"
	"io"

	"github.com/Swapica/swapica-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type ExecuteMatchRequest struct {
	Data resources.ExecuteMatchRequest
}

func ParseExecuteMatchBody(r io.Reader, tx *resources.EvmTransaction) ([]byte, error) {
	request := ExecuteMatchRequest{}

	if err := json.NewDecoder(r).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to decode request")
	}

	request.Data.RawTxData = &tx.Attributes.TxBody.Data
	rawBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return rawBody, nil
}
