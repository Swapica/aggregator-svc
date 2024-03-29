package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Swapica/aggregator-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func AppendTxToBody(r *http.Request, tx *resources.EvmTransaction) ([]byte, error) {
	clonedBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}
	r.Body = io.NopCloser(bytes.NewReader(clonedBody))

	if tx == nil {
		return clonedBody, nil
	}

	var fields map[string]map[string]interface{}

	err = json.Unmarshal(clonedBody, &fields)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request body")
	}

	fields["data"]["raw_tx_data"] = &tx.Attributes.TxBody.Data

	newBody, err := json.Marshal(fields)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal new body")
	}

	return newBody, nil
}

func SendRequest(body io.Reader, endpoint string) (*resources.EvmTransaction, int, error) {
	res, err := http.Post(endpoint, "application/json", body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to send request, endpoint: "+endpoint)
	}
	if res.StatusCode == 400 {
		return nil, res.StatusCode, nil
	}

	tx, err := ParseEvmTransactionResponse(res)
	return tx, res.StatusCode, err
}

type TxResponse struct {
	Data     resources.EvmTransaction `json:"data"`
	Included resources.Included       `json:"included"`
}

func ParseEvmTransactionResponse(r *http.Response) (*resources.EvmTransaction, error) {
	var response TxResponse

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return &response.Data, errors.Wrap(err, "failed to unmarshal EvmTransaction")
	}

	return &response.Data, nil
}
