package cms

import (
	"cmsalegra/configuration"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestCMS(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Contains(t, r.URL.RawQuery, "begins=")
		assert.Contains(t, r.URL.RawQuery, "ends=")
		assert.Equal(t, "Bearer mockToken", r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"data": []}`)
	}))
	defer server.Close()

	config := configuration.Configuration{
		CMSApi: configuration.CmsApi{
			UrlApiCmsLogin:         server.URL,
			UrlApiCmsConsulta:      server.URL,
			UrlApiCmsLoginEmail:    "emailtest@mail.com",
			UrlApiCmsLoginPassword: "passwordtest",
		},
	}

	mockCMSClient := new(MockCMSClient)

	mockCMSClient.On("ApiCMSLogin", config).Return("mockToken", nil)

	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"data": []}`)),
	}
	mockCMSClient.On("RequestCMS", "2024-01-01", config).Return(mockResponse, nil).Once()

	resp, err := RequestCMS("2024-01-01", config, mockCMSClient)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestQueryApiByteCMSReports(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Contains(t, r.URL.RawQuery, "begins=")
		assert.Contains(t, r.URL.RawQuery, "ends=")
		assert.Equal(t, "Bearer mockToken", r.Header.Get("Authorization"))

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{
            "data": [
                {
                    "id": 1,
                    "user_id": 123,
                    "email": "test@example.com",
                    "first_name": "John",
                    "last_name": "Doe",
                    "alegra_transaction": {
                        "invoice_relation": {
                            "invoice_items": [
                                {
                                    "original_price": 100.0
                                }
                            ]
                        }
                    },
                    "in_usd": 100.0,
                    "exchange_rate": 1.0,
                    "currency": "USD",
                    "payment_method": "credit_card",
                    "original_price": 100.0
                },
				{
                    "id": 1,
                    "user_id": 123,
                    "email": "test@example.com",
                    "first_name": "John",
                    "last_name": "Doe",
                    "alegra_transaction": {
                        "invoice_relation": {
                            "invoice_items": [
                                {
                                    "original_price": 20.0
                                }
                            ]
                        }
                    },
                    "in_usd": 20.0,
                    "exchange_rate": 1.0,
                    "currency": "USD",
                    "payment_method": "credit_card",
                    "original_price": 20.0
                }
            ]
        }`)
	}))
	defer server.Close()

	config := configuration.Configuration{
		CMSApi: configuration.CmsApi{
			UrlApiCmsLogin:         server.URL,
			UrlApiCmsConsulta:      server.URL,
			UrlApiCmsLoginEmail:    "test@example.com",
			UrlApiCmsLoginPassword: "password",
		},
	}

	mockClient := new(MockCMSClient)

	mockClient.On("ApiCMSLogin", config).Return("mockToken", nil)

	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(strings.NewReader(`{
            "data": [
                {
                    "id": 1,
                    "user_id": 123,
                    "email": "test@example.com",
                    "first_name": "John",
                    "last_name": "Doe",
                    "alegra_transaction": {
                        "invoice_relation": {
                            "invoice_items": [
                                {
                                    "original_price": 101.0
                                }
                            ]
                        }
                    },
                    "in_usd": 101.0,
                    "exchange_rate": 1.0,
                    "currency": "USD",
                    "payment_method": "credit_card",
                    "original_price": 101.0
                }
            ]
        }`)),
	}
	mockClient.On("RequestCMS", "2024-01-01", config).Return(mockResponse, nil)

	jsonData, totalInvoices, totalAmount, err := QueryApiByteCMSReports("2024-01-01", config, mockClient)

	assert.NoError(t, err)
	assert.NotNil(t, jsonData)
	assert.Equal(t, 1, totalInvoices)
	assert.Equal(t, 120.0, totalAmount)

}
