package cms

import (
	"cmsalegra/configuration"

	"github.com/stretchr/testify/mock"
)

type MockCMSClient struct {
	mock.Mock
}

func (m *MockCMSClient) ApiCMSLogin(config configuration.Configuration) (string, error) {
	args := m.Called(config)
	return args.String(0), args.Error(1)
}
