package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type TwilioParams struct {
	MessageSid string `json:"messagesid" xml:"messagesid" form:"messagesid"`
	AccountSid string `json:"accountsid" xml:"accountsid" form:"accountsid"`
	From       string `json:"from" xml:"from" form:"from"`
	Body       string `json:"body" xml:"body" form:"body"`
}

func Webhook(c *fiber.Ctx) error {
	//authToken := os.Getenv("TWILLIO_TOKEN")

	// Initialize the request validator
	//requestValidator := client.NewRequestValidator(authToken)

	// Store Twilio's request URL (the url of your webhook) as a variable

	params := new(TwilioParams)

	if err := c.BodyParser(params); err != nil {
		return err
	}

	//Log the paramaters
	fmt.Println("Body: ", params.Body)
	fmt.Println("messageSid: ", params.MessageSid)
	fmt.Println("AccountSid: ", params.AccountSid)
	fmt.Println("From: ", params.From)

	// Store the X-Twilio-Signature header attached to the request as a variable
	signature := c.Get("X-Twilio-Signature")
	fmt.Println("Signature: ", signature)

	// Check if the incoming signature is valid for your application URL and the incoming parameters
	//fmt.Println(requestValidator.Validate(url, params, signature))
	return c.SendString("done webhook")
}
