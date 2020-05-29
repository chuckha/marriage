# marriage

This is designed to be run with [AWS Lambda](https://aws.amazon.com/lambda/).

To get started, build the bundle with `make`. The default target builds a zip bundle that will work with lambda.

You will need a slack oauth token and a slack channel to post to.

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

## Lambda set up

1. Log in to the aws console
0. Go to the [lambda service page](https://console.aws.amazon.com/lambda/home?region=us-east-1#/create/function)

    1. Fill in some name like `project-cupid` or whatever you want, this doesn't matter
    0. Select `Go 1.x` as the runtime
    0. No permissions are necessary
    0. Create function

0. In the configuration page

    1. Under `Designer`
    
        0. Click `Add trigger`
        0. Select `CloudWatch Events/EventBridge`
        0. For the rule, you want to pick `Create a new rule`
    
            1. Rule name isn't relevant, something like "project-cupid-schedule' will work
            0. Select `Schedule expression`
            0. Fill in an expression like `rate(2 hours)` which means once every two hours
            0. Make sure enable trigger is selected
            0. click `Add`

    0. Under `Function code` change the `Handler` to `main`
        
        1. Upload the zip file built at the very beginning

    0. Under `Environmnet Variables`
    
        1. set `SLACK_CHANNEL_ID` to the channel ID of the slack channel you want this function to post to
        0. set `SLACK_OAUTH_TOKEN` to the oauth token provided by slack
        
    0. Save your configuration
    0. At the top, next to `Test`, click the dropdown labled `Select a test event`
    0. Select `Configure test event`
    0. Name it anything you like, I called mine `whatever`
    0. Push `Create`
    0. Push `Test`
    0. If you don't see anything in your slack channel then there is an error with the configuraiton, check the cloudwatch logs
    
