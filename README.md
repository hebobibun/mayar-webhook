# Webhook Server

This is a basic webhook server written in Go, designed to handle various events of [Mayar](https://mayar.id/).

## Prerequisites

Make sure you have the following installed:

- Go (Golang)
- Dependencies (use `go mod tidy`)

## Getting Started

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/webhook-server.git
   cd webhook-server
2. **Set up your environment variables:**

   Create a .env file and add your Mayar token:
   ```bash
   MAYAR_TOKEN=your_mayar_token
   PORT=8080
3. Build and run the server:
   ```bash
   go run main.go

The server will run on http://localhost:8080 by default.

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
| shipper.status               | Handle shipper status events               |failed    |
| membership.memberUnsubscribed| Handle member unsubscribed events          |done    |
| membership.memberExpired      | Handle member expired events               |done |

## TODO 
Verify the Mayar token before processing the payload.



