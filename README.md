
# Background
Integrating Slack with CDK calls offers significant benefits, especially for long-running processes. By receiving notifications directly in Slack, users can promptly respond to the outcomes of CDK deployments, making timely corrections and restarts as necessary. This integration is particularly useful for:

- **Monitoring Long-Running Processes**: CDK deployments or updates can take time. Slack integration provides real-time notifications, allowing users to focus on other tasks without the need to constantly check the terminal.
- **Immediate Feedback**: Receive instant alerts on the status of CDK commands, enabling quicker responses to success or failure messages.
- **Collaborative Troubleshooting**: Share deployment statuses with team members in Slack channels, fostering collaborative debugging and problem-solving.
- **Automated Workflows**: Streamline CI/CD pipelines by integrating CDK command outputs directly into Slack, providing a centralized platform for deployment updates.

# Usage of GoCDK

## Overview
This README provides instructions on how to use the compiled binary of the GoCDK project and integrate it with Slack for streamlined CDK command execution and Slack notifications.

## Installation and Configuration

### Step 1: Clone and Compile the GoCDK Project
```bash
git clone https://github.com/yourusername/gocdk.git
cd gocdk
go build -o gocdk
```
This creates an executable named `gocdk`.

### Step 2: Move the Binary to a Desired Location
```bash
mv gocdk /path/to/your/other/project
# Or globally
sudo mv gocdk /usr/local/bin
```

### Step 3: Detailed Slack Integration Instructions
#### Setting Up Slack App for OAuth Token
1. Navigate to the [Slack API](https://api.slack.com/apps) and click "Create New App".
2. Choose "From scratch", name your app, and select your Slack workspace.
3. In "OAuth & Permissions", under "Scopes", add `chat:write` to Bot Token Scopes.
4. Install the app to your workspace. After installation, copy the "Bot User OAuth Access Token".

#### Configuring the GoCDK Project
Set the following environment variables:
- `SLACK_TOKEN` with your Slack Bot User OAuth Access Token.
- `SLACK_CHANNEL_ID` with the ID of the Slack channel where messages will be sent.
These can be set in your shell or through a `.env` file.

### Step 4: Using the Binary
Execute the binary in your project's directory to pass arguments directly to CDK:
```bash
cd /path/to/your/other/project
./gocdk [cdk arguments]
```
Or, if in your PATH:
```bash
gocdk [cdk arguments]
```
The `gocdk` binary acts as a wrapper, passing all arguments to the AWS CDK command line interface.

## Troubleshooting
- Ensure the `gocdk` binary has the correct execution permissions (`chmod +x gocdk`).
- Verify the correct path to the binary and correct setting of environment variables.
- For Slack integration, ensure a valid OAuth token and accurate channel ID.

## Contributing
Contributions to this project are welcome. Please follow the standard fork-and-pull request workflow.

## License
MIT License
