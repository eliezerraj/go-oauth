package service

import(
	"time"
	"github.com/rs/zerolog/log"

	"github.com/go-oauth/internal/core/model"	
	go_core_aws_dynamo "github.com/eliezerraj/go-core/aws/dynamo"
	go_core_aws_secret_manager "github.com/eliezerraj/go-core/aws/secret_manager" 
)

var childLogger = log.With().Str("core", "service").Logger()

type WorkerService struct {
	coreSecretManager 	*go_core_aws_secret_manager.AwsSecretManager
	coreDynamoDB 		*go_core_aws_dynamo.DatabaseDynamo
	awsService			*model.AwsService
	Keys				*model.RsaKey
	TokenSignedValidation 	func(string, interface{}) (*model.JwtData, error)
	CreatedToken 			func(interface{}, time.Time, model.JwtData) (*model.Authentication, error)
}

// About create a ner worker service
func NewWorkerService(	coreSecretManager 	*go_core_aws_secret_manager.AwsSecretManager,
						coreDynamoDB 		*go_core_aws_dynamo.DatabaseDynamo,
						awsService			*model.AwsService,
						keys				*model.RsaKey,
						tokenSignedValidation 	func(string, interface{}) (*model.JwtData, error),
						createdToken 			func(interface{}, time.Time, model.JwtData) (*model.Authentication, error) ) (*WorkerService, error) {
	childLogger.Debug().Msg("NewWorkerService")

	return &WorkerService{
		coreSecretManager: coreSecretManager,
		coreDynamoDB: coreDynamoDB,
		awsService: awsService,
		Keys: keys,
		TokenSignedValidation: tokenSignedValidation,
		CreatedToken: createdToken,
	}, nil
}