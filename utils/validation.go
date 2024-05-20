package utils

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
)

func ValidateAndFormatPhoneNumber(phone, region string) (string, error) {
	num, err := phonenumbers.Parse(phone, region)
	if err != nil {
		return "", fmt.Errorf("failed to parse phone number: %v", err)
	}
	if !phonenumbers.IsValidNumber(num) {
		return "", fmt.Errorf("invalid phone number: %s", phone)
	}
	formattedNumber := phonenumbers.Format(num, phonenumbers.E164)
	return formattedNumber, nil
}
