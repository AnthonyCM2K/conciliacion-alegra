package cms

import (
	"cmsalegra/configuration"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiCMSLogin(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, "/login", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"data": {"token": "mockToken"}}`)
	}))
	defer server.Close()

	config := configuration.Configuration{
		CMSApi: configuration.CmsApi{
			UrlApiCmsLogin:         server.URL,
			UrlApiCmsLoginEmail:    "test@example.com",
			UrlApiCmsLoginPassword: "password",
		},
	}

	mockClient := new(MockCMSClient)

	mockClient.On("ApiCMSLogin", config).Return("mockToken", nil)

	token, err := mockClient.ApiCMSLogin(config)
	assert.NoError(t, err)
	assert.Equal(t, "mockToken", token)

	mockClient.AssertExpectations(t)
}
