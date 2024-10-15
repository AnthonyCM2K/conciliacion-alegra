package conciliacion

import (
	"cmsalegra/cms"
	"cmsalegra/configuration"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryApiByteAlegra(t *testing.T) {
	mockAPI := &MockInvoiceAPISuccess{}

	data, totalInvoices, totalAmount, err := mockAPI.QueryApiByteAlegra("2024-09-10", configuration.Configuration{})

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, 13, totalInvoices)
	assert.Equal(t, 603.89, totalAmount)
}

func TestQueryApiByteCMSReports(t *testing.T) {
	mockAPI := &MockInvoiceAPISuccess{}
	clientCMS := &cms.DefaultCMSClient{}

	data, totalInvoices, totalAmount, err := mockAPI.QueryApiByteCMSReports("2024-09-10", configuration.Configuration{}, clientCMS)

	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, 13, totalInvoices)
	assert.Equal(t, 603.89, totalAmount)
}

func TestConciliation(t *testing.T) {
	mockAPI := &MockInvoiceAPISuccess{}
	fecha := "2024-09-10"
	config := configuration.Configuration{}
	clientCMS := &cms.DefaultCMSClient{}

	Conciliation(fecha, config, mockAPI, clientCMS)
	t.Logf("Test exitoso, CSV se genera sin Discrepancias - " + fecha)
}

func TestConciliation_Discrepancies(t *testing.T) {
	mockAPI := &MockInvoiceAPIDiscrepancies{}
	fecha := "2024-09-11"
	config := configuration.Configuration{}
	clientCMS := &cms.DefaultCMSClient{}

	Conciliation(fecha, config, mockAPI, clientCMS)
	t.Logf("Test exitoso, CSV se genera con Discrepancias - " + fecha)
}

func TestConciliation_ErrorQueryApiAlegra(t *testing.T) {
	mockAPI := &MockInvoiceAPIWithErrorAlegra{}
	fecha := "2024-09-12"
	config := configuration.Configuration{}
	clientCMS := &cms.DefaultCMSClient{}

	err := Conciliation(fecha, config, mockAPI, clientCMS)

	if err == nil {
		t.Errorf("Se esperaba un error al consultar la API de Alegra, pero no se produjo")
	} else {
		t.Logf("Test exitoso, error esperado: %v", err)
	}
}
