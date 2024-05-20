package clients

import (
	"fmt"
	"github.com/twilio/twilio-go"
	"math/rand"
	"preeti-kansal-24/MidasLab.git/constants"
	"time"

	api "github.com/twilio/twilio-go/rest/api/v2010"
)

// Function to generate a random OTP of given length
func generateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < length; i++ {
		otp += fmt.Sprintf("%d", rand.Intn(10))
	}
	return otp
}

func GenerateOTP(phoneNumber string) (string, string, error) {
	// Your Twilio Account SID and Auth Token from twilio.com/console
	accountSid := constants.TwilioSID
	authToken := constants.TwilioAuth

	// Create a Twilio client
	client := twilio.NewRestClientWithParams(
		twilio.ClientParams{Username: accountSid, Password: authToken})

	from := constants.TwilioFromPhone
	// Generate OTP
	otp := generateOTP(6)

	fmt.Printf("otp is %v, phone number is %v\n", phoneNumber, otp)

	// Compose the message body
	message := fmt.Sprintf("Your OTP is: %s", otp)

	// Send message using Twilio API
	messageParams := &api.CreateMessageParams{
		To:   &phoneNumber,
		From: &from,
		Body: &message,
	}

	//Ignoring the msg as of now as I dont have a twilio number to send otp from

	client.Api.CreateMessage(messageParams)
	//if err != nil {
	//	fmt.Println("Error sending OTP:", err.Error())
	//	return "", otp, err
	//}
	return fmt.Sprintf("OTP sent successfully"), otp, nil
}
