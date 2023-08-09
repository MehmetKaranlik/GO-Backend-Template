package Environment

import (
	"Backend/Core/App/Configuration"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

type Env struct {
	MongoConnectionString string
	AWSSecrets            AWSSecrets
	JWTSecrets            JWTSecrets
}

type AWSSecrets struct {
	BaseURL    string
	BucketName string
	Region     string
	ID         string
	Secret     string
}

type JWTSecrets struct {
	AccessSecretKey  string
	RefreshSecretKey string
	VerifySecretKey  string
}

func (e *Env) LoadEnvironment(configuration Configuration.IAppConfiguration) {
	err := godotenv.Load(configuration.GetEnvDir())
	fmt.Println()
	if err != nil {
		log.Err(err)
		os.Exit(1)
	}

}

func (e *Env) MapEnvironmentValues() {
	e.MongoConnectionString = os.Getenv(mongo_connection_string.String())
	e.AWSSecrets = AWSSecrets{
		BaseURL:    os.Getenv(aws_base_url.String()),
		BucketName: os.Getenv(aws_bucket_name.String()),
		Region:     os.Getenv(aws_region.String()),
		ID:         os.Getenv(aws_id.String()),
		Secret:     os.Getenv(aws_secret.String()),
	}
	e.JWTSecrets = JWTSecrets{
		AccessSecretKey:  os.Getenv(access_token_salt.String()),
		RefreshSecretKey: os.Getenv(refresh_token_salt.String()),
		VerifySecretKey:  os.Getenv(verify_token_salt.String()),
	}
}
