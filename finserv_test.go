package finserv_test

// import (
// 	"context"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"

// 	"gitlab.playcourt.id/new-mypertamina/myptm-gopayment-service/pkg/finserv"
// )

// func TestFinserv(t *testing.T) {
// 	pp := &finserv.PaymentProperty{
// 		TransactionID: "trxid123",
// 		MID:           "mid123",
// 		TID:           "tid123",
// 		MSISDN:        "628124541588",
// 		Items: []*finserv.Item{
// 			new(finserv.Item),
// 			new(finserv.Item),
// 		},
// 		CallbackURL: "https://www.blabla.com",
// 	}

// 	t.Run("Should return wco", func(t *testing.T) {
// 		var expectedWCO finserv.WCOData = "wco"
// 		c := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			mockResponse := `{
// 				"success": true,
// 				"data": "wco",
// 				"message": "success",
// 				"code": 200
// 			}`
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusOK)
// 			io.WriteString(w, mockResponse)
// 		}))

// 		fc := finserv.NewClient(c.Client(), c.URL, "123", finserv.BasicSecret, "")
// 		payment := fc.CapturePayment(pp)

// 		wcoData, err := payment.DoWebCheckout(context.Background())

// 		if err != nil {
// 			t.Error(err)
// 		}

// 		assert.Equal(t, expectedWCO, wcoData)
// 	})

// 	t.Run("Should return err internal server", func(t *testing.T) {
// 		// var expectedWCO finserv.WCOData = "wco"
// 		c := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			mockResponse := `{
// 				"success": true,
// 				"data": "wco",
// 				"message": "success",
// 				"code": 200
// 			}`
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusOK)
// 			io.WriteString(w, mockResponse)
// 		}))

// 		fc := finserv.NewClient(c.Client(), " aa aa ", "123", finserv.BasicSecret, "")
// 		payment := fc.CapturePayment(pp)

// 		_, err := payment.DoWebCheckout(context.Background())

// 		if err != nil {
// 			assert.Equal(t, finserv.ErrInternalServer, err)
// 		}
// 	})

// 	t.Run("Should return err bad request", func(t *testing.T) {
// 		// var expectedWCO finserv.WCOData = "wco"
// 		c := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			mockResponse := `{
// 				"success": false,
// 				"data": "",
// 				"message": "Bad Request",
// 				"code": 400
// 			}`
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusBadRequest)
// 			io.WriteString(w, mockResponse)
// 		}))

// 		fc := finserv.NewClient(c.Client(), c.URL, "123", finserv.BasicSecret, "")
// 		payment := fc.CapturePayment(pp)

// 		_, err := payment.DoWebCheckout(context.Background())

// 		if err != nil {
// 			assert.Equal(t, finserv.ErrBadRequest, err)
// 		}
// 	})

// 	t.Run("Should return err not authorize in context", func(t *testing.T) {
// 		// var expectedWCO finserv.WCOData = "wco"
// 		c := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			mockResponse := `{
// 				"success": false,
// 				"data": "",
// 				"message": "Not Authorize",
// 				"code": 401
// 			}`
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			io.WriteString(w, mockResponse)
// 		}))

// 		fc := finserv.NewClient(c.Client(), c.URL, "123", finserv.BasicSecret, "")
// 		payment := fc.CapturePayment(pp)

// 		_, err := payment.DoWebCheckout(context.Background())

// 		if err != nil {
// 			assert.Equal(t, finserv.ErrUnauthorized, err)
// 		}
// 	})

// 	t.Run("Should return err timeout in context", func(t *testing.T) {
// 		// var expectedWCO finserv.WCOData = "wco"
// 		c := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			mockResponse := `{
// 				"success": true,
// 				"data": "wco",
// 				"message": "success",
// 				"code": 200
// 			}`
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusOK)
// 			io.WriteString(w, mockResponse)
// 		}))

// 		fc := finserv.NewClient(c.Client(), c.URL, "123", finserv.BasicSecret, "")
// 		payment := fc.CapturePayment(pp)

// 		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond*1)
// 		defer cancel()
// 		_, err := payment.DoWebCheckout(ctx)

// 		if err != nil {
// 			assert.Equal(t, finserv.ErrTimeout, err)
// 		}
// 	})

// 	t.Run("Should return err internal server when undifined code appeared", func(t *testing.T) {
// 		// var expectedWCO finserv.WCOData = "wco"
// 		c := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			mockResponse := `{
// 				"success": false,
// 				"data": "",
// 				"message": "Not Authorize",
// 				"code": 500
// 			}`
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusInternalServerError)
// 			io.WriteString(w, mockResponse)
// 		}))

// 		fc := finserv.NewClient(c.Client(), c.URL, "123", finserv.BasicSecret, "")
// 		payment := fc.CapturePayment(pp)

// 		_, err := payment.DoWebCheckout(context.Background())

// 		if err != nil {
// 			assert.Equal(t, finserv.ErrInternalServer, err)
// 		}
// 	})
// }
