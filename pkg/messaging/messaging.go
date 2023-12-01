package messaging

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SendMessage sends a WhatsApp message using the specified parameters.
func SendMessage(phoneNumber, message, apiKey string) error {
	// URL encode the message
	encodedMessage := url.QueryEscape(message)

	// Construct the GET request URL
	apiURL := fmt.Sprintf("https://api.callmebot.com/whatsapp.php?phone=%s&text=%s&apikey=%s",
		url.QueryEscape(phoneNumber), encodedMessage, url.QueryEscape(apiKey))

	fmt.Println(apiURL)

	// Perform the GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read and log the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	fmt.Printf("HTTP Response Status: %s\n", resp.Status)
	fmt.Printf("HTTP Response Body: %s\n", string(responseBody))

	return nil
}
