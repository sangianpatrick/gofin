package finserv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"golang.org/x/net/context/ctxhttp"
)

// ConfirmPaymentResult is a result after capture a transaction
type ConfirmPaymentResult struct {
	err             error
	pgpToken        string
	referenceNumber string
	fastTime        string
}

// NewConfirmPaymentResult is a constructor for doing test
func NewConfirmPaymentResult(pgpToken string, referenceNumber string, fastTime string, err error) *ConfirmPaymentResult {
	return &ConfirmPaymentResult{
		pgpToken:        pgpToken,
		referenceNumber: referenceNumber,
		fastTime:        fastTime,
		err:             err,
	}
}

// Err return error
func (cpr *ConfirmPaymentResult) Err() error {
	return cpr.err
}

// PGPToken will return transaction id and error
func (cpr *ConfirmPaymentResult) PGPToken() string {
	return cpr.pgpToken
}

// ReferenceNumber will return transaction id and error
func (cpr *ConfirmPaymentResult) ReferenceNumber() string {
	return cpr.referenceNumber
}

// FastTime will return transaction id and error
func (cpr *ConfirmPaymentResult) FastTime() string {
	return cpr.fastTime
}

type paymentClient struct {
	client *http.Client
	host   string
	secret string
	logger *logrus.Logger
}

// NewPaymentClient is a constructor
func NewPaymentClient(httpClient *http.Client, host, secret string, logger *logrus.Logger) PaymentClient {
	if logger == nil {
		logger = logrus.New()
	}

	return &paymentClient{
		client: httpClient,
		host:   host,
		secret: secret,
		logger: logger,
	}
}

func (pc *paymentClient) logResponse(res *http.Response) {
	defer res.Body.Close()
	resBody, _ := ioutil.ReadAll(res.Body)
	resMap := map[string]interface{}{
		"message": string(resBody),
	}
	pc.logger.Error(resMap)
	return
}

func (pc *paymentClient) do(ctx context.Context, resultChan chan<- *ConfirmPaymentResult, client *http.Client, req *http.Request) {
	defer close(resultChan)
	res, err := ctxhttp.Do(ctx, client, req)

	result := new(ConfirmPaymentResult)

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
		if res.StatusCode == http.StatusUnauthorized {
			pc.logResponse(res)
			result.err = ErrUnauthorized
			resultChan <- result
			return
		}

		if res.StatusCode == http.StatusUnprocessableEntity {
			pc.logResponse(res)
			result.err = ErrInvalidParameter
			resultChan <- result
			return
		}

		if res.StatusCode == http.StatusNotFound {
			pc.logResponse(res)
			result.err = ErrTransactionNotFound
			resultChan <- result
			return
		}
		pc.logResponse(res)
		result.err = ErrInternalServer
		resultChan <- result
		return
	}

	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)

	var e paymentEnvlope

	if err := json.Unmarshal(resBody, &e); err != nil {
		result.err = ErrInternalServer
		resultChan <- result
		return
	}

	successData := make(map[string]string, 0)

	json.Unmarshal(e.Data, &successData)

	result.pgpToken = successData["pgpToken"]
	result.referenceNumber = successData["refNum"]
	result.fastTime = successData["fastTime"]

	resultChan <- result
	return
}

// ConfirmForWebCheckout is a function to confirm payment for web checkout.
func (pc *paymentClient) ConfirmForWebCheckout(ctx context.Context, pp *PaymentProperty) *ConfirmPaymentResult {
	resultChan := make(chan *ConfirmPaymentResult)

	url := fmt.Sprintf("%s%s", pc.host, confirmPaymentPath)

	ppByte, _ := json.Marshal(pp)
	fmt.Println(string(ppByte))
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(ppByte))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", pc.secret)

	go pc.do(ctx, resultChan, pc.client, req)

	result := <-resultChan
	return result
}
