package direct7

import (
	"log"
	"github.com/google/uuid"
	"fmt"
)

// VERIFY struct represents the VERIFY service in Go.
type Verify struct {
	client *Client
}

// NewVerify creates a new instance of the VERIFY service.
func NewVerify(client *Client) *Verify {
	return &Verify{client: client}
}

// SendOTP sends an OTP to a single recipient.
func (v *Verify) SendOTP(originator, recipient, content, dataCoding string, expiry int, templateID int) (string, error) {
	params := map[string]interface{}{
		"originator":  originator,
		"recipient":   recipient,
		"content":     content,
		"data_coding": dataCoding,
		"expiry":      expiry,
		"template_id": templateID,
	}
	if templateID != 0 {
		params = map[string]interface{}{
			"originator":  originator,
			"recipient":   recipient,
			"template_id": templateID,
		}
	}
	log.Println(params)
	response, err := v.client.Post("/verify/v1/otp/send-otp", true, params)
	if err != nil {
		return "", err
	}
	log.Println("OTP message sent successfully.")
	return string(response), nil
}

// ResendOTP re-sends an OTP to a single recipient.
func (v *Verify) ResendOTP(otpID string) (string, error) {
	// Parse the OTP ID string into a UUID
	otpUUID, err := uuid.Parse(otpID)
	if err != nil {
		return "", fmt.Errorf("error parsing OTP ID: %v", err)
	}

	params := map[string]interface{}{
		"otp_id": otpUUID,
	}
	response, err := v.client.Post("/verify/v1/otp/resend-otp", true, params)
	if err != nil {
		return "", err
	}
	log.Println("OTP message re-sent successfully.")
	return string(response), nil
}

// VerifyOTP verifies an OTP.
func (v *Verify) VerifyOTP(otpID, otpCode string) (string, error) {
	otpUUID, err := uuid.Parse(otpID)
	if err != nil {
		return "", fmt.Errorf("error parsing OTP ID: %v", err)
	}
	params := map[string]interface{}{
		"otp_id":   otpUUID,
		"otp_code": otpCode,
	}
	response, err := v.client.Post("/verify/v1/otp/verify-otp", true, params)
	if err != nil {
		return "", err
	}
	log.Println("OTP message verified successfully.")
	return string(response), nil
}

// GetStatus gets the status for an OTP request.
func (v *Verify) GetStatus(otpID string) (string, error) {
	response, err := v.client.Get("/verify/v1/report/"+otpID, nil)
	if err != nil {
		return "", err
	}
	log.Println("OTP message status retrieved successfully.")
	return string(response), nil
}
