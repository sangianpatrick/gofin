// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import finserv "gitlab.playcourt.id/new-mypertamina/myptm-gopayment-service/pkg/finserv"
import mock "github.com/stretchr/testify/mock"

// QRClient is an autogenerated mock type for the QRClient type
type QRClient struct {
	mock.Mock
}

// Capture provides a mock function with given fields: ctx, qrp
func (_m *QRClient) Capture(ctx context.Context, qrp *finserv.QRPayload) *finserv.QRCaptureResult {
	ret := _m.Called(ctx, qrp)

	var r0 *finserv.QRCaptureResult
	if rf, ok := ret.Get(0).(func(context.Context, *finserv.QRPayload) *finserv.QRCaptureResult); ok {
		r0 = rf(ctx, qrp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*finserv.QRCaptureResult)
		}
	}

	return r0
}
