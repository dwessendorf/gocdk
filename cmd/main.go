package main

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/dwessendorf/gocdk/pkg/cdkwrapper"
	"github.com/dwessendorf/gocdk/pkg/messaging"
)

func main() {
	// Get phone number and API key from environment variables
	phoneNumber := os.Getenv("GOCDK_PHONE")
	apiKey := os.Getenv("GOCDK_APIKEY")
	if phoneNumber == "" || apiKey == "" {
		log.Fatalf("Environment variables GOCDK_PHONE or GOCDK_APIKEY are not set")
	}

	// Check if CDK arguments are provided
	if len(os.Args) < 2 {
		log.Fatalf("No CDK command provided")
	}

	// os.Args[0] is the program name, os.Args[1:] are the actual CDK arguments
	cdkArgs := os.Args[1:]

	// Execute the CDK command
	output, exitStatus, err := cdkwrapper.ExecuteCDKCommand(cdkArgs)
	if err != nil {
		log.Printf("Error executing CDK command: %v", err)
	}

	replacedString := strings.ReplaceAll(output, "==", "=")
	replacedString = strings.ReplaceAll(replacedString, "$", "Dollar")
	replacedString = strings.ReplaceAll(replacedString, "function", "function_")

	// Prepare the output message
	re := regexp.MustCompile(`[\x1b\x0f]\[[0-9;]*m|┌─┐|└─┘|├─┤|│|─|└┴┘├┼┤┌┬┐└┴┘├┼┼┼┼┼┤`)
	outputString := re.ReplaceAllString(replacedString, "")

	// Keep only the last 2200 characters
	if len(outputString) > 1400 {
		outputString = outputString[:len(outputString)-1400]
	}

	// Append "SUCCESS" or "FAILURE"
	statusMessage := "FAILURE"
	if exitStatus == 0 {
		statusMessage = "SUCCESS"
	}
	outputString = statusMessage + "\n" + outputString

	// Send the message in chunks
	const maxMessageLength = 500
	for start := 0; start < len(outputString); start += maxMessageLength {
		end := start + maxMessageLength
		if end > len(outputString) {
			end = len(outputString)
		}

		messageChunk := outputString[start:end]
		if err := messaging.SendMessage(phoneNumber, messageChunk, apiKey); err != nil {
			log.Fatalf("Error sending WhatsApp message: %v", err)
		}
	}
}
