# Webhook Server

This is a basic webhook server written in Go, designed to handle various events of [Mayar](https://mayar.id/).

## Prerequisites

Make sure you have the following installed:

- Go (Golang)
- Dependencies (use `go mod tidy`)
- WA sender service : [Whatsapp web js](https://wwebjs.dev/) / [WHMCS](https://github.com/Intprism-Technology/Whatsapp-WHMCS) or anything

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/webhook-server.git
   cd webhook-server
2. **Set up your environment variables:**

   Create a local.env file and add your Mayar token:
   ```bash
   MAYAR_TOKEN=your_mayar_token
   PORT=8090
   WASERVICEHOST=127.0.0.1
   WASERVICEPORT=8080
   ADMIN=+621111111111
3. Build and run the server:
   ```bash
   go run main.go
   
The server will run on http://localhost:8090 by default.

## Usage

Webhook Endpoints
| METHOD           | Endpoint          | Description |
| ---------------- | ----------------- |-------------|
| POST             | /                 |Receive incoming webhooks. |


## Supported Events

| Event                        | Description                               | Status |
| ---------------------------- | ----------------------------------------- |--------|
| testing                      | Handle testing events                     |done    |
| payment.received             | Handle payment received events            |done    |
| payment.reminder             | Handle payment reminder events            |done    |
| shipper.status               | Handle shipper status events               |done    |
| membership.memberUnsubscribed| Handle member unsubscribed events          |done    |
| membership.memberExpired      | Handle member expired events               |done |

## TODO 
[done] Verify the Mayar token before processing the payload.



