package finserv

import "errors"

// Error list
var (
	ErrInvalidRequest      = errors.New("Finserv Client: Request is not valid")
	ErrTimeout             = errors.New("Finserv Client: Request is timeout")
	ErrInternalServer      = errors.New("Finserv Client: Internal Server Error")
	ErrBadRequest          = errors.New("Finserv Client: Bad Request")
	ErrUnauthorized        = errors.New("Finserv Client: Invalid Secret Key")
	ErrClient              = errors.New("Finserv Client: Client Error")
	ErrDatabase            = errors.New("Finserv Client: Gagal input data ke database")
	ErrInvalidParameter    = errors.New("Finserv Client: Form harus diisi")
	ErrTransactionNotFound = errors.New("Finserv Client: Data transaksi tidak di temukan")
)
