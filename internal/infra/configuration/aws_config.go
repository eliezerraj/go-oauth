package configuration

import(
	"os"

	"github.com/joho/godotenv"
	"github.com/go-oauth/internal/core/model"
)

// About get AWS service env ver
func GetAwsServiceEnv() model.AwsService {
	childLogger.Info().Str("func","GetAwsServiceEnv").Send()

	err := godotenv.Load(".env")
	if err != nil {
		childLogger.Info().Err(err).Send()
	}
	
	var awsService	model.AwsService

	if os.Getenv("AWS_REGION") !=  "" {
		awsService.AwsRegion = os.Getenv("AWS_REGION")
	}

	if os.Getenv("SECRET_NAME") !=  "" {
		awsService.SecretName = os.Getenv("SECRET_NAME")
	}

	if os.Getenv("DYNAMO_TABLE_NAME") !=  "" {
		awsService.DynamoTableName = os.Getenv("DYNAMO_TABLE_NAME")
	}

	return awsService
}