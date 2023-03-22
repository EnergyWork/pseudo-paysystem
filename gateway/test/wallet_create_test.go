package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/energywork/pseudo-paysystem/gateway/api/usecase"
)

func TestWalletCreate(t *testing.T) {
	req := usecase.ReqGateWalletCreate{
		Phone: "89991112233",
	}

	js, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:9002/v2/wallet/create", bytes.NewReader(js))
	if err != nil {
		t.Fatal(err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: time.Second * 30,
	}

	httpReq.Close = true

	resp, err := client.Do(httpReq)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Response Body: %s", string(respBody))
}

func TestWalletCreate2(t *testing.T) {
	req := usecase.ReqGateWalletCreate{
		Phone: "89991112233",
	}

	js, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post("http://localhost:9002/v1/wallet/create", "application/json", bytes.NewReader(js))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Response Body: %s", string(respBody))
}
