package conciliacion

import (
	"cmsalegra/cms"
	"cmsalegra/configuration"
	"fmt"
)

type MockInvoiceAPISuccess struct{}

func (m *MockInvoiceAPISuccess) QueryApiByteAlegra(fecha string, config configuration.Configuration) ([]byte, int, float64, error) {
	return []byte(`[
        {"id":6421,"amount":30,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz soles"}},
        {"id":6420,"amount":90,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Paypal"}},
        {"id":6419,"amount":28,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz dólares"}},
        {"id":6418,"amount":9,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz dólares"}},
        {"id":6417,"amount":9,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz dólares"}},
        {"id":6416,"amount":60,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Paypal"}},
        {"id":6415,"amount":76.92,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz soles"}},
        {"id":6414,"amount":28,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz dólares"}},
        {"id":6403,"amount":8.97,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz soles"}},
        {"id":6402,"amount":30,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz soles"}},
        {"id":6401,"amount":180,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz dólares"}},
        {"id":6400,"amount":30,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz dólares"}},
        {"id":6399,"amount":24,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz dólares"}}
    ]`), 13, 603.89, nil
}

func (m *MockInvoiceAPISuccess) QueryApiByteCMSReports(fecha string, config configuration.Configuration, client cms.CMSClient) ([]byte, int, float64, error) {
	return []byte(`[
        {"id":271298,"user_id":10006739,"email":"usuariotest@gmail.com","first_name":"Diego","last_name":"Quesada","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":24}]},"alegra_payment_id":"6399","alegra_data":{"bank_account":""}},"in_usd":24,"exchange_rate":1,"currency":"","payment_method":"","original_price":24},
        {"id":271299,"user_id":14710608,"email":"usuariotest@gmail.com","first_name":"Jhon","last_name":"juares Blas","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":30}]},"alegra_payment_id":"6400","alegra_data":{"bank_account":""}},"in_usd":30,"exchange_rate":1,"currency":"","payment_method":"","original_price":30},
        {"id":271331,"user_id":6566897,"email":"usuariotest@gmail.com","first_name":"Geraldine","last_name":"Rivera","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":180}]},"alegra_payment_id":"6401","alegra_data":{"bank_account":""}},"in_usd":180,"exchange_rate":1,"currency":"","payment_method":"","original_price":180},
        {"id":271430,"user_id":15370616,"email":"usuariotest@gmail.com","first_name":"Julio jesús ","last_name":"Juares castro ","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":28}]},"alegra_payment_id":"6414","alegra_data":{"bank_account":""}},"in_usd":28,"exchange_rate":1,"currency":"","payment_method":"","original_price":28},
        {"id":271463,"user_id":15357485,"email":"usuariotest@gmail.com","first_name":"jj","last_name":"cr","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":300}]},"alegra_payment_id":"6415","alegra_data":{"bank_account":""}},"in_usd":76.92,"exchange_rate":3.9,"currency":"","payment_method":"","original_price":300},
        {"id":271364,"user_id":94016469,"email":"usuariotest@gmail.com","first_name":"ELVIS","last_name":"juanes TEJADA","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":117}]},"alegra_payment_id":"6402","alegra_data":{"bank_account":""}},"in_usd":30,"exchange_rate":3.9,"currency":"","payment_method":"","original_price":117},
        {"id":271397,"user_id":15376155,"email":"usuariotest@gmail.com","first_name":"VICTOR","last_name":"mirlo BORCHANI","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":35}]},"alegra_payment_id":"6403","alegra_data":{"bank_account":""}},"in_usd":8.97,"exchange_rate":3.9,"currency":"","payment_method":"","original_price":35},
        {"id":271496,"user_id":15064362,"email":"usuariotest@gmail.com","first_name":"Carlos Saul","last_name":"peres Sanchez","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":60}]},"alegra_payment_id":"6416","alegra_data":{"bank_account":""}},"in_usd":60,"exchange_rate":1,"currency":"","payment_method":"","original_price":60},
        {"id":271529,"user_id":15378780,"email":"usuariotest@gmail.com","first_name":"Jovan","last_name":"edura","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":9}]},"alegra_payment_id":"6417","alegra_data":{"bank_account":""}},"in_usd":9,"exchange_rate":1,"currency":"","payment_method":"","original_price":9},
        {"id":271530,"user_id":15387880,"email":"usuariotest@gmail.com","first_name":"Jovan","last_name":"educa","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":9}]},"alegra_payment_id":"6418","alegra_data":{"bank_account":""}},"in_usd":9,"exchange_rate":1,"currency":"","payment_method":"","original_price":9},
        {"id":271562,"user_id":13793506,"email":"usuariotest@gmail.com","first_name":"Aixa","last_name":"suarez","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":28}]},"alegra_payment_id":"6419","alegra_data":{"bank_account":""}},"in_usd":28,"exchange_rate":1,"currency":"","payment_method":"","original_price":28},
        {"id":271595,"user_id":13750752,"email":"usuariotest@gmail.com","first_name":"Felipe","last_name":"juanjo","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":90}]},"alegra_payment_id":"6420","alegra_data":{"bank_account":""}},"in_usd":90,"exchange_rate":1,"currency":"","payment_method":"","original_price":90},
        {"id":271628,"user_id":15938142,"email":"usuariotest@gmail.com","first_name":"Jhuliño","last_name":"mirlo Ramirez","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":117}]},"alegra_payment_id":"6421","alegra_data":{"bank_account":""}},"in_usd":30,"exchange_rate":3.9,"currency":"","payment_method":"","original_price":117}
    ]`), 13, 603.89, nil
}

type MockInvoiceAPIDiscrepancies struct{}

func (m *MockInvoiceAPIDiscrepancies) QueryApiByteAlegra(fecha string, config configuration.Configuration) ([]byte, int, float64, error) {
	return []byte(`[
        {"id":6421,"amount":30,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz soles"}},
        {"id":6417,"amount":90,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz dólares"}},
        {"id":6416,"amount":60,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Paypal"}},
        {"id":6415,"amount":76.92,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz soles"}},
        {"id":6403,"amount":80.97,"currency":[{"code":"USD","exchangeRate":3.9}],"bankAccount":{"name":"Pasarela Niubiz soles"}}
    ]`), 5, 603.89, nil
}

func (m *MockInvoiceAPIDiscrepancies) QueryApiByteCMSReports(fecha string, config configuration.Configuration, client cms.CMSClient) ([]byte, int, float64, error) {
	return []byte(`[
        {"id":271430,"user_id":1537061,"email":"usuariotest@gmail.com","first_name":"Julio jesús ","last_name":"Pérez castro ","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":28}]},"alegra_payment_id":"6414","alegra_data":{"bank_account":""}},"in_usd":28,"exchange_rate":1,"currency":"","payment_method":"","original_price":28},
        {"id":271463,"user_id":1537485,"email":"usuariotest@gmail.com","first_name":"jj","last_name":"cr","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":300}]},"alegra_payment_id":"6415","alegra_data":{"bank_account":""}},"in_usd":76.92,"exchange_rate":3.9,"currency":"","payment_method":"","original_price":300},
        {"id":271397,"user_id":1537155,"email":"usuariotest@gmail.com","first_name":"VICTOR","last_name":"VILCA BORCHANI","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":35}]},"alegra_payment_id":"6403","alegra_data":{"bank_account":""}},"in_usd":8.97,"exchange_rate":3.9,"currency":"","payment_method":"","original_price":35},
        {"id":271496,"user_id":1506462,"email":"usuariotest@gmail.com","first_name":"Carlos Saul","last_name":"Cante Sanchez","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":60}]},"alegra_payment_id":"6416","alegra_data":{"bank_account":""}},"in_usd":60,"exchange_rate":1,"currency":"","payment_method":"","original_price":60},
        {"id":271529,"user_id":1537880,"email":"usuariotest@gmail.com","first_name":"Jovan","last_name":"Pacheco","alegra_transaction":{"invoice_relation":{"invoice_items":[{"original_price":9}]},"alegra_payment_id":"6417","alegra_data":{"bank_account":""}},"in_usd":9,"exchange_rate":1,"currency":"","payment_method":"","original_price":9}
    ]`), 5, 603.89, nil
}

type MockInvoiceAPIWithErrorAlegra struct{}

// Implementación del método que devuelve un error
func (m *MockInvoiceAPIWithErrorAlegra) QueryApiByteAlegra(fecha string, config configuration.Configuration) ([]byte, int, float64, error) {
	return nil, 0, 0, fmt.Errorf("error al consultar API de Alegra")
}

func (m *MockInvoiceAPIWithErrorAlegra) QueryApiByteCMSReports(fecha string, config configuration.Configuration, client cms.CMSClient) ([]byte, int, float64, error) {
	// Devuelve datos de prueba o un resultado esperado en caso de que no necesites que falle
	return []byte(`[{"ID":1, "AlegraTransactionListResponse": {"AlegraPaymentID": "1", "InvoiceRelationListResponse": {"InvoiceItems": [{"OriginalPrice": 100.00}]}}}]`), 1, 100.00, nil
}
