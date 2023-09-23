package config

/*
Note : Read the Meta Docs to get Phone number Id and Access token
*/

const VERIFY_TOKEN string = "***********" // Use the same token in Meta configuration
const PREFIX byte = '.'                   // Prefix to invoke the service
const ACCESS_TOKEN string = "**********"
const PHONE_NUMBER_ID string = "*********"

const ENDPOINT string = "https://graph.facebook.com/v18.0/" + PHONE_NUMBER_ID + "/messages"
