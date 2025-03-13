package main

import(
	"time"
	"os"
	"context"
	"encoding/json"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/go-oauth/internal/infra/configuration"
	"github.com/go-oauth/internal/adapter/api"
	"github.com/go-oauth/internal/core/model"
	"github.com/go-oauth/internal/core/service"
	"github.com/go-oauth/internal/infra/server"

	go_core_cert "github.com/eliezerraj/go-core/cert"
	go_core_aws_config "github.com/eliezerraj/go-core/aws/aws_config"
	go_core_aws_dynamo "github.com/eliezerraj/go-core/aws/dynamo"
	go_core_aws_secret_manager "github.com/eliezerraj/go-core/aws/secret_manager" 
)

var(
	logLevel = 	zerolog.DebugLevel
	appServer	model.AppServer
	awsConfig 	go_core_aws_config.AwsConfig
	databaseDynamo		go_core_aws_dynamo.DatabaseDynamo
	awsSecretManager	go_core_aws_secret_manager.AwsSecretManager
)

// About initialize the enviroment var
func init(){
	log.Debug().Msg("init")
	zerolog.SetGlobalLevel(logLevel)

	infoPod, server := configuration.GetInfoPod()
	configOTEL 		:= configuration.GetOtelEnv()
	awsService 	:= configuration.GetAwsServiceEnv() 

	appServer.InfoPod = &infoPod
	appServer.Server = &server
	appServer.ConfigOTEL = &configOTEL
	appServer.AwsService = &awsService
}

// About loads all key (HS256 and RSA)
func loadKey(ctx context.Context, secretName string, coreSecretManager *go_core_aws_secret_manager.AwsSecretManager) (*model.RsaKey, error){
	log.Debug().Msg("loadKey")

	// Load symetric key from secret manager
	var certCore go_core_cert.CertCore

	keys := model.RsaKey{}
	secret, err := coreSecretManager.GetSecret(ctx, secretName)
	if err != nil {
		return nil, err
	}
	var secretData map[string]string
	if err := json.Unmarshal([]byte(*secret), &secretData); err != nil {
		return nil, err
	}
	keys.JwtKey = secretData["JWT_KEY"]

	// Load the private key
	private_key, err := os.ReadFile("../assets/certs/server-private.key")
	if err != nil{
		return nil, err
	}
	keys.Key_rsa_priv_pem = string(private_key)

	// Convert private key
	key_rsa_priv, err := certCore.ParsePemToRSAPriv(&keys.Key_rsa_priv_pem)
	if err != nil{
		return nil, err
	}
	keys.Key_rsa_priv = key_rsa_priv

	// Load the public key
	public_key, err := os.ReadFile("../assets/certs/server-public.key")
	if err != nil {
		return nil, err
	}
	keys.Key_rsa_pub_pem = string(public_key)
	// Convert public key
	key_rsa_pub, err := certCore.ParsePemToRSAPub(&keys.Key_rsa_pub_pem)
	if err != nil{
		return nil, err
	}
	keys.Key_rsa_pub = key_rsa_pub

	// Load the crl
	crl_pem, err := os.ReadFile("../assets/certs/crl-ca.crl")
	if err != nil {
		return nil, err
	}
	keys.Crl_pem = string(crl_pem)

	return &keys, nil
}

// About main
func main (){
	log.Debug().Msg("main")
	log.Debug().Msg("----------------------------------------------------")
	log.Debug().Interface("appServer :",appServer).Msg("")
	log.Debug().Msg("----------------------------------------------------")

	ctx, cancel := context.WithTimeout(	context.Background(), 
										time.Duration( appServer.Server.ReadTimeout ) * time.Second)
	defer cancel()

	// Prepare aws services
	awsConfig, err := awsConfig.NewAWSConfig(ctx, appServer.InfoPod.AWSRegion)
	if err != nil {
		panic("error create new aws session " + err.Error())
	}

	coreDynamoDB := databaseDynamo.NewDatabaseDynamo(awsConfig)
	coreSecretManager := awsSecretManager.NewAwsSecretManager(awsConfig)

	appServer.RsaKey, err = loadKey(ctx, 
									appServer.AwsService.SecretName, 
									coreSecretManager)
	if err != nil {
		panic("error get keys" + err.Error())
	}

	// wire	
	workerService, err := service.NewWorkerService(	coreDynamoDB, 
													appServer.AwsService, 
													appServer.RsaKey,
													service.TokenValidationRSA,
													service.CreatedTokenRSA)
	if err != nil {
		panic("error create a workerservice " + err.Error())
	}
	httpRouters := api.NewHttpRouters(workerService)
	httpServer := server.NewHttpAppServer(appServer.Server)

	// start server
	httpServer.StartHttpAppServer(ctx, &httpRouters, &appServer)
}