package cms

import (
	"cmsalegra/configuration"
	"net/http"
)

// RequestCMS simula una llamada a la API
func (m *MockCMSClient) RequestCMS(fecha string, config configuration.Configuration) (*http.Response, error) {
	args := m.Called(fecha, config)
	// Asegúrate de manejar el caso donde args.Get(0) pueda ser nil
	if res, ok := args.Get(0).(*http.Response); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}
