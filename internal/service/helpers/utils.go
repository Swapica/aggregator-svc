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
	bodyCloned, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}
	r.Body = io.NopCloser(bytes.NewReader(bodyCloned))

	if tx == nil {
		return bodyCloned, nil
	}

	var fields map[string]map[string]interface{}

	err = json.Unmarshal(bodyCloned, &fields)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal request body")
	}

	fields["data"]["raw_tx_data"] = &tx.Attributes.TxBody.Data

	rawBody, err := json.Marshal(fields)
	if err != nil {
		return nil, err
	}

	return rawBody, nil
}

func SendRequest(body io.Reader, endpoint string) (*resources.EvmTransaction, error) {
	res, err := http.Post(endpoint, "application/json", body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request, endpoint: "+endpoint)
	}

	return ParseEvmTransactionResponse(res)
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
