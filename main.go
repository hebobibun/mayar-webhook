package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"webhook/config"
)

var (
	mayarToken    string
	servicePort   int
	waServiceHost string
	waServicePort string
	admin         string
)

func init() {
	cfg := config.InitConfig()

	mayarToken = cfg.MayarToken
	servicePort = cfg.Port
	waServiceHost = cfg.WAServiceHost
	waServicePort = cfg.WAServicePort
	admin = cfg.Admin

	if servicePort <= 0 {
		servicePort = 8090
	}
}

func main() {
	http.HandleFunc("/", webhookHandler)

	fmt.Printf("Webhook server is running on port %d...\n", servicePort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {

	// Verify request source
	if r.Header["X-Callback-Token"][0] != mayarToken {
		log.Println("Unauthorized webhook", mayarToken)
		http.Error(w, "Unauthorized webhook", http.StatusUnauthorized)
		return
	}

	log.Println("Authorized.")

	// Verify method
	if r.Method != http.MethodPost {
		log.Println("Wrong method")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Replace the json.Decode with json.Unmarshal
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var payload RequestPayload
	err = json.Unmarshal(bodyBytes, &payload)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
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
	case "shipper.status":
		handleShipperStatusEvent(payload.Data)

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
	log.Println("Cust Name  :", data.CustomerName)
	log.Println("Cust Email :", data.CustomerEmail)
	log.Println("Cust Phone :", data.CustomerMobile)

	custPhone := formatPhoneNumber(data.CustomerMobile)

	// Create payload for the WhatsApp message
	payload := map[string]string{
		"phone": custPhone,
		"message": `Pesanan Telah Terkonfirmasi

Halo ` + data.CustomerName + `!
		
Terima kasih telah berbelanja di toko kami! Pesanan dan pembayaran Anda telah terkonfirmasi. 🎉
		
Rincian Pesanan:
		
*Nomor Pesanan:* 
` + fmt.Sprintf("%v", data.ID) + `
		
*Nama Produk:* 
` + data.ProductName + `
	
*Email Anda:* 
` + data.CustomerEmail + `
		
*Total Pembayaran:* 
` + fmt.Sprintf("%v", data.Amount) + `
		
*Status Pembayaran:* 
Sudah dibayar
	
Akses ke pesanan Anda akan segera dikirimkan melalui Email.
		
Jika ada pertanyaan, jangan ragu untuk menghubungi kami.
	
Salam hangat,
Tim Kami`,
	}

	// send meesage to customer's WA
	err := sendWhatsAppMessage(payload, custPhone)
	if err != nil {
		log.Println("Error sending WA to cust", err)
		return
	}
	log.Println("Sent message to :", custPhone)
	// send notif to admin WA
	err = sendWhatsAppMessage(payload, admin)
	if err != nil {
		log.Println("Error sending WA to admin", err)
		return
	}
	log.Println("Sent message to admin")
}

func handlePaymentReminderEvent(data Data) {
	// Implement logic for handling payment.reminder event
	fmt.Println("Handling payment.reminder event...")
	fmt.Println("ID:", data.ID)
	fmt.Println("Status:", data.Status)
}

func handleShipperStatusEvent(data Data) {
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

func sendWhatsAppMessage(payload map[string]string, phone string) error {
	apiURL := fmt.Sprintf("http://%v:%v/api/send", waServiceHost, waServicePort)

	// Convert the payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Println("error marshalling payload WA", err)
		return err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Println("error new request WA", err)
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Perform the HTTP request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error client Do WA", err)
		return err
	}
	defer resp.Body.Close()

	// Check the HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Println("error status", resp.Status)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
func formatPhoneNumber(number string) string {
	if strings.HasPrefix(number, "0") {
		// If the number starts with 0, replace it with +62
		return "+62" + number[1:]
	}
	// If the number doesn't start with 0, assume it's already in the desired format
	return "+" + number
}
