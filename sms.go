// Package sms provides a small wrapper around AWS SNS SMS support.
package sms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

// SMS configures an SNS SMS client.
type SMS struct {
	svc      snsiface.SNSAPI // Service implementation
	SenderID string          // SenderID (optional)
}

// Send `message` to `number`.
func (s *SMS) Send(message, number string) error {
	attrs := map[string]*sns.MessageAttributeValue{}

	if s.SenderID != "" {
		attrs["AWS.SNS.SMS.SenderID"] = &sns.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: &s.SenderID,
		}
	}

	params := &sns.PublishInput{
		Message:           &message,
		PhoneNumber:       &number,
		MessageAttributes: attrs,
	}

	_, err := s.svc.Publish(params)
	return err
}

func NewSMS(AccessKeyId, SecretAccessKey, Region string) *SMS {
	awsSession := session.New(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(AccessKeyId, SecretAccessKey, ""),
	})
	if awsSession == nil {
		return nil
	}
	svc := sns.New(awsSession)
	if svc == nil {
		return nil
	}
	return &SMS{svc: svc}
}
