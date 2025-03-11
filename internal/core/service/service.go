package service

import(
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
	keys				*model.RsaKey
}

// About create a ner worker service
func NewWorkerService(	coreSecretManager 	*go_core_aws_secret_manager.AwsSecretManager,
						coreDynamoDB 		*go_core_aws_dynamo.DatabaseDynamo,
						awsService			*model.AwsService,
						keys				*model.RsaKey ) (*WorkerService, error) {
	childLogger.Debug().Msg("NewWorkerService")

	return &WorkerService{
		coreSecretManager: coreSecretManager,
		coreDynamoDB: coreDynamoDB,
		awsService: awsService,
		keys: keys,
	}, nil
}