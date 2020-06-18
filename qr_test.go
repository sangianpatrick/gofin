package finserv_test

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"testing"

// 	"gitlab.playcourt.id/new-mypertamina/myptm-gopayment-service/pkg/finserv"
// )

// func TestCapture(t *testing.T) {

// 	qrc := finserv.NewQRClient(http.DefaultClient, "http://apidev.my-pertamina.id/finserv-qr", "Basic Zmluc2VydjpteVAzclQ0bUluNC1maW41M1JW")
// 	qrcr := qrc.Capture(context.Background(), new(finserv.QRPayload))

// 	jsonByte, _ := json.Marshal(qrcr)
// 	fmt.Println(string(jsonByte))
// }
