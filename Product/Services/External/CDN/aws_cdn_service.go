package CDN

import (
	"Backend/Core/Globals"
	"Backend/Core/Utilities/Methods"
	"bytes"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rs/zerolog/log"
)

const (
	InvalidFilenameError = "Invalid filename"
)

// Implements ICDNService interface
type AWSCDNService struct {
}

// Implemenets AWS'S Logger interface to log aws operations
type awsLogger struct {
}

func (A awsLogger) Log(i ...interface{}) {
	log.Info().Msgf("%v", i)
}

func (self *AWSCDNService) connect() *s3.S3 {
	config := self.config()
	session := session.Must(session.NewSession(config))
	s3Object := s3.New(session)
	return s3Object
}

func (self *AWSCDNService) UploadFile(file []byte, fileName string) (string, error) {
	s3Object := self.connect()
	newFilename := Methods.ConstructNewNameForFile(fileName)
	if newFilename == "" {
		return "", errors.New(InvalidFilenameError)
	}
	input := s3.PutObjectInput{
		Bucket: aws.String(Globals.EnvValues.AWSSecrets.BucketName),
		Body:   bytes.NewReader(file),
		Key:    aws.String(Methods.ConstructNewNameForFile(newFilename)),
	}

	if _, err := s3Object.PutObject(&input); err != nil {
		return "", err
	}

	return Globals.EnvValues.AWSSecrets.BaseURL + newFilename, nil

}

func (self *AWSCDNService) disconnect() error {
	// There is no need to close s3 connection its shutting down automatically
	// Thus we are returning nil
	return nil
}

func (self *AWSCDNService) config() *aws.Config {
	return &aws.Config{
		Region:      aws.String(Globals.EnvValues.AWSSecrets.Region),
		Credentials: credentials.NewStaticCredentials(Globals.EnvValues.AWSSecrets.ID, Globals.EnvValues.AWSSecrets.Secret, ""),
		Logger:      awsLogger{},
	}
}
