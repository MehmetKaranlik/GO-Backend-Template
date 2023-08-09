package Environment

type EnvKeys string

const (
	// Mongo

	mongo_connection_string EnvKeys = "MONGO_CONNECTION_STRING"
	// AWS
	aws_base_url    EnvKeys = "AWS_BASE_URL"
	aws_bucket_name EnvKeys = "aws_bucket_name"
	aws_region      EnvKeys = "aws_region"
	aws_id          EnvKeys = "aws_id"
	aws_secret      EnvKeys = "aws_secret"
	// JWT
	access_token_salt  EnvKeys = "access_token_salt"
	refresh_token_salt EnvKeys = "refresh_token_salt"
	verify_token_salt
)

func (e EnvKeys) String() string {
	return string(e)
}
