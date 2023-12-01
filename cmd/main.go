package main

import (
	"log"
	"os"
	"regexp"

	"github.com/dwessendorf/gocdk/pkg/cdkwrapper"
	"github.com/dwessendorf/gocdk/pkg/messaging"
)

func main() {
	// Get Slack OAuth token and Channel ID from environment variables
	slackToken := os.Getenv("SLACK_TOKEN")
	slackChannelID := os.Getenv("SLACK_CHANNEL_ID")
	if slackToken == "" || slackChannelID == "" {
		log.Fatalf("Environment variables SLACK_TOKEN or SLACK_CHANNEL_ID are not set")
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

	// replacedString := strings.ReplaceAll(output, "==", "=")
	// replacedString = strings.ReplaceAll(replacedString, "$", "Dollar")
	// replacedString = strings.ReplaceAll(replacedString, "function", "function_")

	// Prepare the output message
	re := regexp.MustCompile(`[\x1b\x0f]\[[0-9;]*m`)
	output = re.ReplaceAllString(output, "")

	if len(output) > 3000 {
		start := len(output) - 3000
		output = output[start:]
	}

	output = "```\n" + output + "\n```"
	// Append "SUCCESS" or "FAILURE"
	statusMessage := "FAILURE"
	if exitStatus == 0 {
		statusMessage = "SUCCESS"
	}
	output = statusMessage + "\n" + output
	if err := messaging.SendMessage(slackToken, slackChannelID, output); err != nil {
		log.Fatalf("Error sending Slack message: %v", err)
	}
	// Send the message in chunks
}
