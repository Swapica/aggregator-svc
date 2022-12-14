package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Swapica/aggregator-svc/internal/service/requests"
	"github.com/Swapica/swapica-svc/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func AppendTxToBody(r *http.Request, tx *resources.EvmTransaction, actionTarget string) ([]byte, error) {
	bodyCloned, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read request body")
	}
	r.Body = io.NopCloser(bytes.NewReader(bodyCloned))

	if tx == nil {
		return bodyCloned, nil
	}

	var body []byte

	bodyReader := bytes.NewReader(bodyCloned)

	switch actionTarget {
	case "v1/create/match":
		body, err = requests.ParseCreateMatchBody(bodyReader, tx)
	case "v1/cancel/match":
		body, err = requests.ParseCancelMatchBody(bodyReader, tx)
	case "v1/execute/order":
		body, err = requests.ParseExecuteOrderBody(bodyReader, tx)
	case "v1/execute/match":
		body, err = requests.ParseExecuteMatchBody(bodyReader, tx)
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed parse request body")
	}

	return body, nil
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
