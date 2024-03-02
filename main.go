package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

var (
	mayarToken string
)

func init() {
	// Set up Viper to read from environment variables
	viper.AutomaticEnv()

	// Retrieve the value of the environment variable
	mayarToken = viper.GetString("MAYAR_TOKEN")

	log.Println("MAYAR TOKEN", mayarToken)

	// Check if the environment variable is set
	if mayarToken == "" {
		fmt.Println("Environment variable not set.")
		// os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", webhookHandler)

	port := 8080
	fmt.Printf("Webhook server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {

	// TODO : VERIFY MAYAR TOKEN HERE

	if r.Method != http.MethodPost {
		log.Println("Wrong method")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	log.Println("HOOK Running...")

	var payload RequestPayload
	var payloadShipping ShipperStatusEvent
	var isShip bool
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&payloadShipping)
		if err != nil {
			log.Println("Error decoding JSON payload")
			http.Error(w, "Error decoding JSON payload", http.StatusBadRequest)
			return
		}
		isShip = true
	}
	defer r.Body.Close()

	if isShip {
		handleShipperStatusEvent(payloadShipping.Data)

		w.WriteHeader(http.StatusOK)
		return
	}

	event := payload.Event

	// Handle the webhook data based on the event type and data
	switch event {
	case "testing":
		handleTestingEvent(payload.Data)
	case "payment.received":
		handleRequestPayload(payload.Data)
	case "payment.reminder":
		handlePaymentReminderEvent(payload.Data)
	case "membership.memberUnsubscribed":
		handleMemberUnsubscribedEvent(payload.Data)
	case "membership.memberExpired":
		handleMemberExpiredEvent(payload.Data)

	default:
		fmt.Println("Unhandled event:", event)
	}

	w.WriteHeader(http.StatusOK)
}

func handleTestingEvent(data Data) {
	fmt.Println("Handling testing event...")
	fmt.Println("ID:", data.ID)
	fmt.Println("Status:", data.Status)
}

func handleRequestPayload(data Data) {
	// Implement logic for handling payment.received event
	fmt.Println("Handling payment.received event...")
	fmt.Println("ID:", data.ID)
	fmt.Println("Status:", data.Status)
}

func handlePaymentReminderEvent(data Data) {
	// Implement logic for handling payment.reminder event
	fmt.Println("Handling payment.reminder event...")
	fmt.Println("ID:", data.ID)
	fmt.Println("Status:", data.Status)
}

func handleShipperStatusEvent(data ShipperStatusData) {
	// Implement logic for handling shipper.status event
	fmt.Println("Handling shipper.status event...")
	fmt.Println("ID:", data.ID)
	fmt.Println("Status:", data.Status)
}

func handleMemberUnsubscribedEvent(data Data) {
	// Implement logic for handling membership.memberUnsubscribed event
	fmt.Println("Handling membership.memberUnsubscribed event...")
	fmt.Println("ID:", data.ID)
	fmt.Println("Status:", data.Status)
}

func handleMemberExpiredEvent(data Data) {
	// Implement logic for handling membership.memberExpired event
	fmt.Println("Handling membership.memberExpired event...")
	fmt.Println("ID:", data.ID)
	fmt.Println("Status:", data.Status)
}
