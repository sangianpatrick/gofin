package finserv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/context/ctxhttp"
)

// QRCaptureResult is a result after capture a transaction
type QRCaptureResult struct {
	transactionID string
	err           error
}

// NewQRCaptureResult is a constructor for doing test
func NewQRCaptureResult(transactionID string, err error) *QRCaptureResult {
	return &QRCaptureResult{
		transactionID: transactionID,
		err:           err,
	}
}

// Transaction will return transaction id and error
func (qrcr *QRCaptureResult) Transaction() (string, error) {
	return qrcr.transactionID, qrcr.err
}

type qrClient struct {
	client *http.Client
	host   string
	secret string
}

// NewQRClient is a constructor
func NewQRClient(httpClient *http.Client, host, secret string) QRClient {
	return &qrClient{
		client: httpClient,
		host:   host,
		secret: secret,
	}
}

func (q *qrClient) do(ctx context.Context, resultChan chan<- *QRCaptureResult, client *http.Client, req *http.Request) {
	defer close(resultChan)
	res, err := ctxhttp.Do(ctx, client, req)

	result := new(QRCaptureResult)

	if err != nil {
		if err == context.DeadlineExceeded {
			result.err = ErrTimeout
			resultChan <- result
			return
		}

		result.err = ErrClient
		resultChan <- result
		return
	}

	if res.StatusCode != http.StatusOK {
		result.err = ErrInternalServer
		resultChan <- result
		return
	}

	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)

	var e qrEnvlope

	if err := json.Unmarshal(resBody, &e); err != nil {
		result.err = ErrInternalServer
		resultChan <- result
		return
	}

	switch e.RC {
	case rcDBError:
		result.err = ErrDatabase
		break
	case rcFormError:
		result.err = ErrInvalidParameter
		break
	case rcSuccess:
		result.err = nil
		result.transactionID = e.TransactionID
		break
	default:
		result.err = ErrInternalServer
		break
	}

	resultChan <- result
	return
}

// Capture is a function to send QRPayload fo finserv server.
// qrHost is a parameter that contains url that point to finserv qr service,
// if the value is not empty or "", it will replace the current host.
func (q *qrClient) Capture(ctx context.Context, qrp *QRPayload) *QRCaptureResult {
	resultChan := make(chan *QRCaptureResult)

	url := fmt.Sprintf("%s%s", q.host, qrCapturePath)

	qrp.TransactionType = QRPayment
	qrp.DateTime = time.Now().Format(DateFormat)

	qrpByte, _ := json.Marshal(qrp)

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(qrpByte))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", q.secret)

	go q.do(ctx, resultChan, q.client, req)

	result := <-resultChan
	return result
}
