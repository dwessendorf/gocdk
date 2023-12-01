# GoCDK Slack Integration

## Overview

This application integrates the GoCDK project with Slack, allowing you to send output from CDK commands directly to a Slack channel. This README provides instructions on setting up the integration and obtaining the necessary Slack OAuth token.

## Prerequisites

* A Slack account and a workspace where you can create apps.
* Golang environment set up for running and compiling the GoCDK project.

## Setting Up Slack App for OAuth Token

To use this integration, you need to create a Slack app and obtain an OAuth token:

1. **Create a New Slack App** :

* Navigate to the [Slack API](https://api.slack.com/apps) page and click on "Create New App".
* Choose "From scratch" and give your app a name.
* Select the Slack workspace where you want to install the app.

2. **Add Permissions to Your Slack App** :

* In the app settings, go to "OAuth & Permissions".
* Under "Scopes", add the necessary permissions. For sending messages, add `chat:write`.

3. **Install App to Workspace** :

* At the top of the "OAuth & Permissions" page, click "Install to Workspace" and authorize the permissions.

1. **Copy the OAuth Access Token** :

* After installation, you will see an "OAuth Access Token" on the same page. Copy this token.

## Configuring the GoCDK Project

Set up your environment with the Slack OAuth token and the channel ID where messages will be sent:

1. **Set Environment Variables** :

* Set `SLACK_TOKEN` with your OAuth Access Token.
* Set `SLACK_CHANNEL_ID` with the ID of the channel where messages will be sent.
* These can be set in your shell or through a `.env` file.

2. **Compile and Run the GoCDK Project** :

* Navigate to your GoCDK project directory.
* Compile and run your project. It should now send messages to the specified Slack channel.

## Usage

Run the application with CDK commands as arguments. The output, along with a status message indicating success or failure, will be sent to the specified Slack channel.

## Troubleshooting

* Ensure the Slack app has the correct permissions and the OAuth token is valid.
* Check the channel ID for accuracy.
* For detailed errors, check the application logs.

## Contributing

Contributions to this project are welcome. Please follow the standard fork-and-pull request workflow.

## License

MIT License
