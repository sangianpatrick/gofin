package finserv

import "context"

// QRClient is a finserv qr client to serve qr payment request
type QRClient interface {
	Capture(ctx context.Context, qrp *QRPayload) *QRCaptureResult
}

// PaymentClient is a finserv payment client
type PaymentClient interface {
	ConfirmForWebCheckout(ctx context.Context, pp *PaymentProperty) *ConfirmPaymentResult
}
