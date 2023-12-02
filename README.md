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

# Using the GoCDK Binary in Another Project

## Overview

This section provides instructions on how to use the compiled binary of the GoCDK project in other code repositories. This can be useful if you want to integrate the GoCDK functionalities into different projects or workflows.

## Steps to Use the Binary

### Step 1: Compile the GoCDK Binary

First, ensure that you have compiled the GoCDK project to create an executable binary.

1. Clone the GoCDK repository if you haven't already:
   <pre><div class="bg-black rounded-md"><div class="flex items-center relative text-gray-200 bg-gray-800 dark:bg-token-surface-primary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>bash</span></div></div></pre>

* <pre><div class="bg-black rounded-md"><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-bash">git clone https://github.com/yourusername/gocdk.git
  cd gocdk
  </code></div></div></pre>
* Compile the project:
  <pre><div class="bg-black rounded-md"><div class="flex items-center relative text-gray-200 bg-gray-800 dark:bg-token-surface-primary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>bash</span></div></div></pre>

1. <pre><div class="bg-black rounded-md"><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-bash">go build -o gocdk
   </code></div></div></pre>

   This will create an executable named `gocdk` in your current directory.

### Step 2: Move the Binary to a Desired Location

Decide where you want to use the binary. You can place it in a specific project's directory, or in a directory that's in your system's PATH for easier access.

* To move the binary to a specific project:

  <pre><div class="bg-black rounded-md"><div class="flex items-center relative text-gray-200 bg-gray-800 dark:bg-token-surface-primary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>bash</span></div></div></pre>
* <pre><div class="bg-black rounded-md"><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-bash">mv gocdk /path/to/your/other/project
  </code></div></div></pre>
* To make the binary globally accessible (assuming `/usr/local/bin` is in your PATH):

  <pre><div class="bg-black rounded-md"><div class="flex items-center relative text-gray-200 bg-gray-800 dark:bg-token-surface-primary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>bash</span></div></div></pre>
* <pre><div class="bg-black rounded-md"><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-bash">sudo mv gocdk /usr/local/bin
  </code></div></div></pre>

### Step 3: Using the Binary in Another Project

To use the binary in another project, navigate to the project's directory in the terminal and execute the binary:

<pre><div class="bg-black rounded-md"><div class="flex items-center relative text-gray-200 bg-gray-800 dark:bg-token-surface-primary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>bash</span></div></div></pre>

<pre><div class="bg-black rounded-md"><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-bash">cd /path/to/your/other/project
./gocdk [arguments]
</code></div></div></pre>

Or, if you've placed it in a location in your PATH:

<pre><div class="bg-black rounded-md"><div class="flex items-center relative text-gray-200 bg-gray-800 dark:bg-token-surface-primary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>bash</span></div></div></pre>

<pre><div class="bg-black rounded-md"><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-bash">gocdk [arguments]
</code></div></div></pre>

Replace `[arguments]` with any arguments required by the GoCDK application.

### Step 4: Integrating with Project Workflows

You can integrate the GoCDK binary into your project's workflows or scripts. For example, you can call it from a shell script or include it in your project's build process.

## Troubleshooting

If you encounter issues while using the binary, ensure that:

* The binary has the correct execution permissions (`chmod +x gocdk`).
* You are using the correct path to the binary.
* All necessary environment variables are set correctly.

## Contributing

Contributions to this project are welcome. Please follow the standard fork-and-pull request workflow.

## License

MIT License
