package aws

import "github.com/aws/aws-sdk-go/aws/credentials"

// Config represents AWS IAM configuration.
type Config struct {
	Service string
	Credentials *credentials.Credentials
	Region string
}
