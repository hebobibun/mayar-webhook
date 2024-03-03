package main

import "time"

// RequestPayload represents the "payment.received" event
type RequestPayload struct {
	Event string `json:"event"`
	Data  Data   `json:"data"`
}

// Data represents the data associated with the "payment.received" event
type Data struct {
	ID                          string        `json:"id"`
	Status                      string        `json:"status"`
	TransactionStatus           string        `json:"transactionStatus"`
	CreatedAt                   time.Time     `json:"createdAt"`
	UpdatedAt                   time.Time     `json:"updatedAt"`
	MerchantID                  string        `json:"merchantId"`
	MerchantName                string        `json:"merchantName"`
	MerchantEmail               string        `json:"merchantEmail"`
	CustomerName                string        `json:"customerName"`
	CustomerEmail               string        `json:"customerEmail"`
	CustomerMobile              string        `json:"customerMobile"`
	Amount                      int           `json:"amount"`
	IsAdminFeeBorneByCustomer   bool          `json:"isAdminFeeBorneByCustomer"`
	IsChannelFeeBorneByCustomer bool          `json:"isChannelFeeBorneByCustomer"`
	PaymentUrl                  string        `json:"paymentUrl"`
	ProductID                   string        `json:"productId"`
	ProductName                 string        `json:"productName"`
	ProductType                 string        `json:"productType"`
	PixelFbp                    string        `json:"pixelFbp"`
	PixelFbc                    string        `json:"pixelFbc"`
	TrackingUrl                 string        `json:"trackingUrl"`
	TrackingCode                string        `json:"trackingCode"`
	TrackingDescription         string        `json:"trackingDescription"`
	CustomeField                []interface{} `json:"custom_field"`
}

// ShipperStatusEvent represents the "shipper.status" event
type ShipperStatusEvent struct {
	Event string            `json:"event"`
	Data  ShipperStatusData `json:"data"`
}

// ShipperStatusData represents the data associated with the "shipper.status" event
type ShipperStatusData struct {
	ID                  string    `json:"id"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"createdAt"`
	MerchantID          string    `json:"merchantId"`
	MerchantName        string    `json:"merchantName"`
	MerchantEmail       string    `json:"merchantEmail"`
	CustomerName        string    `json:"customerName"`
	CustomerEmail       string    `json:"customerEmail"`
	CustomerMobile      string    `json:"customerMobile"`
	TrackingUrl         string    `json:"trackingUrl"`
	TrackingCode        string    `json:"trackingCode"`
	TrackingDescription string    `json:"trackingDescription"`
	ProductID           string    `json:"productId"`
	ProductName         string    `json:"productName"`
	ProductType         string    `json:"productType"`
}
