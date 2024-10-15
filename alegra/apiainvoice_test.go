package alegra

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthAlegraAPI(t *testing.T) {
	username := "test@example.com"
	password := "test_token"

	auth, err := basicAuthAlegraAPI(username, password)
	if err != nil {
		t.Fatalf("Error generating basic auth: %v", err)
	}

	expectedAuth := "Basic dGVzdEBleGFtcGxlLmNvbTp0ZXN0X3Rva2Vu"
	if auth != expectedAuth {
		t.Errorf("Expected %s, got %s", expectedAuth, auth)
	}
}

func TestRequestAlegraPayments(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verificar que la URL de la solicitud sea la esperada
		if r.URL.Path != "/api/v1/payments" {
			t.Errorf("Expected request to '/api/v1/payments', got '%s'", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockPayments)
	}))
	defer ts.Close()

	mockConfig.ALEGRAApi.UrlApiAlegra = ts.URL

	payments, err := RequestAlegraPayments("2024-09-10", mockConfig)
	if err != nil {
		t.Fatalf("Error requesting Alegra payments: %v", err)
	}

	if len(payments) == 0 {
		t.Errorf("Expected to receive payments, got none")
	}
}

func TestQueryApiByteAlegra(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockPayments)
	}))
	defer ts.Close()

	mockConfig.ALEGRAApi.UrlApiAlegra = ts.URL

	jsonData, totalInvoices, totalAmount, err := QueryApiByteAlegra("2024-09-10", mockConfig)
	if err != nil {
		t.Fatalf("Error querying API byte Alegra: %v", err)
	}

	if jsonData == nil {
		t.Error("Expected non-nil JSON data")
	}

	if totalInvoices != 2 || totalAmount != 301.05 {
		t.Errorf("Expected totalInvoices 2 and totalAmount 301.05, got totalInvoices %d and totalAmount %f", totalInvoices, totalAmount)
	}
}
