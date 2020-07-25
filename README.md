# marriage

# Lambda function set up

## For the trusting

This project is built and uploaded to s3 as a lambda zip.

1. Figure out your slack channel ID.
2. Figure out your slack oauth token.
3. Upload the cloudformation.yaml file to
   https://console.aws.amazon.com/cloudformation/home?region=us-east-1#/stacks?filteringText=&filteringStatus=active&viewNested=true&hideStacks=false
4. Supply the non-default fields.
5. Click through the forms.
6. Wait and watch slack!

## For the untrusting

1. Gain trust in this code.
2. Build it and upload to s3.
3. Follow the steps above for the trusting, except replace the default values in
   the cloud formation template with your custom values.

## Build

This uses a Makefile to help build. Run `make build` to build the resources and
`make clean` to clean everything up.

## Slack set up

1. [Go here](https://api.slack.com/apps?new_app=1)
0. Enter your bot name, something like `Project Cupid Bot`, you can change this later if you like
0. Select the workspace to use (you will have to be signed in)
0. Select your project which will bring you to the configuration page
0. Under `Add features and functionality`

    1. Make sure `Bots` gets marked
    0. Under `Permissions` set the `chat:write` scope under `Bot Token Scopes`

0. Grab the `Bot User Oauth Token` for later
0. Find the channel ID by going to your workspace and joining the channel and looking at the URL: `https://app.slack.com/client/<workspace_id>/<channel_id>`
