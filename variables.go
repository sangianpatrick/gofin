package finserv

import "encoding/json"

// TransactionID is a type of transaction id
type TransactionID string

// Constant is collection of constant
const (
	QRPayment  = "qr_payment"
	DateFormat = "2006-01-02 15:04:05"
)

// QRPayload is model payload to be sent
type QRPayload struct {
	TransactionType string  `json:"transaction_type"`
	TID             string  `json:"tid"`
	MID             string  `json:"mid"`
	MerchantCode    string  `json:"spbu"`
	DateTime        string  `json:"date_time"`
	ProductCode     string  `json:"product_id"`
	Volume          string  `json:"volume"`
	Amount          string  `json:"amount"`
	CallbackURL     string  `json:"callback_url"`
	Items           []*Item `json:"items"`
	BillingNumber   *string `json:"bill_no,omitempty"`
	Token           *string `json:"token,omitempty"`
}

// PaymentProperty is finserv payment property for linkaja
type PaymentProperty struct {
	TransactionID             string  `json:"transaction_id"`
	MSISDN                    string  `json:"msisdn"`
	MID                       string  `json:"mid"`
	TID                       string  `json:"tid"`
	Amount                    string  `json:"amount"`
	Items                     []*Item `json:"items"`
	VehicleRegistrationNumber string  `json:"nopol"`
}

// Item is finserv item to be include in finserv payment property
type Item struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity string `json:"qty"`
}

const (
	confirmPaymentPath = "/v1/transaction/qr_payment_confirm"
	qrCapturePath      = "/v1/qr-payment"
)

const (
	rcSuccess   = "00"
	rcDBError   = "02"
	rcFormError = "03"
	rcNotFound  = "04"
)

type qrEnvlope struct {
	RC            string `json:"rc"`
	TransactionID string `json:"transaction_id"`
	Message       string `json:"message"`
}

type paymentEnvlope struct {
	Succsess bool            `json:"success"`
	Data     json.RawMessage `json:"data"`
	Message  string          `json:"message"`
	Code     int             `json:"code"`
}
