package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPaymentReceivedHandler(t *testing.T) {
	// JSON payload for testing payment.reminder event
	testJSON := `{
		"event": "payment.received",
		"data": {
		  "id": "9356ec92-32ae-4d99-a1a7-51b11dff4d84",
		  "status": "SUCCESS",
		  "transactionStatus": "created",
		  "createdAt": 1693817623264,
		  "updatedAt": 1693817626638,
		  "merchantId": "348e083d-315a-4e5c-96b1-5a2a98c48413",
		  "merchantName": "Malo Gusto",
		  "merchantEmail": "aldodwier@gmail.com",
		  "customerName": "Student Test",
		  "customerEmail": "student@student.com",
		  "customerMobile": "0815",
		  "amount": 1029,
		  "isAdminFeeBorneByCustomer": true,
		  "isChannelFeeBorneByCustomer": true,
		  "productId": "e2b3f5d5-0c62-47ba-8a01-6c1c209e0f77",
		  "productName": "Kelas Pemrograman Web Dasar",
		  "productType": "course",
		  "pixelFbp": "fb.1.1693462870069.763035785",
		  "pixelFbc": null
		}
	  }`

	// Create a new HTTP request with the JSON payload
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(testJSON)))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Call the webhook handler function with the test request
	webhookHandler(recorder, req)

	// Check the HTTP status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestPaymentReminderHandler(t *testing.T) {
	// JSON payload for testing payment.reminder event
	testJSON := `{
	  "event": "payment.reminder",
	  "data": {
		"id": "b147e908-bba4-4337-ab03-bf9be504a539",
		"status": "SUCCESS",
		"transactionStatus": "created",
		"createdAt": 1693550488964,
		"updatedAt": 1693550703627,
		"merchantId": "4dba4996-7d74-483e-99fe-b52c60368cb5",
		"merchantEmail": "kugutsu.hiruko@gmail.com",
		"customerName": "andre",
		"customerEmail": "alikusnadie@gmail.com",
		"customerMobile": "085",
		"amount": 120000,
		"isAdminFeeBorneByCustomer": false,
		"isChannelFeeBorneByCustomer": false,
		"productId": "f03694c1-2f6d-45c2-8be2-138f9fd34bbf",
		"productName": "juak e book ranl",
		"productType": "digital_product",
		"pixelFbp": null,
		"pixelFbc": null,
		"paymentUrl": "https://korban-motivator.mayar.shop/plt/b147e908-bba4-4337-ab03-bf9be504a539?webhook=true"
	  }
	}`

	// Create a new HTTP request with the JSON payload
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(testJSON)))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Call the webhook handler function with the test request
	webhookHandler(recorder, req)

	// Check the HTTP status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestMembershipUnsubHandler(t *testing.T) {
	// JSON payload for testing payment.reminder event
	testJSON := `{
		"event": "membership.memberUnsubscribed",
		"data": {
		  "merchantId": "4dba4996-7d74-483e-99fe-b52c60368cb5",
		  "merchantName": "andrea",
		  "status": "INACTIVE",
		  "memberId": "BQ1UVCFB",
		  "customerName": "andret",
		  "customerEmail": "kugutsu.hiruko@gmail.com",
		  "customerMobile": "085797522261"
		}
	  }`

	// Create a new HTTP request with the JSON payload
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(testJSON)))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Call the webhook handler function with the test request
	webhookHandler(recorder, req)

	// Check the HTTP status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestMembershipMemberExpiredHandler(t *testing.T) {
	// JSON payload for testing payment.reminder event
	testJSON := `{
		"event": "membership.memberExpired",
		"data": {
		  "merchantId": "4dba4996-7d74-483e-99fe-b52c6873jrjb",
		  "merchantName": "andrea",
		  "status": "INACTIVE",
		  "memberId": "5U2FV9PD",
		  "customerName": "andika",
		  "customerEmail": "alikusnadie@gmail.com",
		  "customerMobile": "0815"
		}
	  }`

	// Create a new HTTP request with the JSON payload
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(testJSON)))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Call the webhook handler function with the test request
	webhookHandler(recorder, req)

	// Check the HTTP status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestShipperStatusHandler(t *testing.T) {
	// JSON payload for testing payment.reminder event
	testJSON := `{
		"event": "shipper.status",
		"data": {
		  "id": "ebee706c-7216-48b1-a770-8971ca0c32bb",
		  "status": "SUCCESS",
		  "createdAt": "2023-09-05T05:17:05.964Z",
		  "merchantId": "4dba4996-7d74-483e-99fe-b52c60368cb5",
		  "merchantName": "andrea",
		  "merchantEmail": "kugutsu.hiruko@gmail.com",
		  "customerName": "andret",
		  "customerEmail": "kugutsu.hiruko@gmail.com",
		  "customerMobile": "0909",
		  "trackingUrl": "https://shipper.id/track/239K3W5VJYWD3",
		  "trackingCode": "239K3W5VJYWD3",
		  "trackingDescription": "Data order sudah masuk ke sistem",
		  "productId": "4be2cb72-b0c6-474e-b057-fc7cd7f34dbe",
		  "productName": "produk miniatur ular",
		  "productType": "physical_product"
		}
	  }`

	// Create a new HTTP request with the JSON payload
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(testJSON)))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	recorder := httptest.NewRecorder()

	// Call the webhook handler function with the test request
	webhookHandler(recorder, req)

	// Check the HTTP status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
