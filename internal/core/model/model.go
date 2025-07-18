package model

import(
	"time"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v4"
	go_core_observ "github.com/eliezerraj/go-core/observability"
)

type AppServer struct {
	InfoPod 		*InfoPod 		`json:"info_pod"`
	Server     		*Server     	`json:"server"`
	ConfigOTEL		*go_core_observ.ConfigOTEL	`json:"otel_config"`
	AwsService		*AwsService		`json:"aws_service"`
	RsaKey			*RsaKey			`json:"rsa_key"`
}

type InfoPod struct {
	PodName				string `json:"pod_name"`
	ApiVersion			string `json:"version"`
	AWSRegion			string `json:"aws_region,omitempty"`
	OSPID				string `json:"os_pid"`
	IPAddress			string `json:"ip_address"`
	AvailabilityZone 	string `json:"availabilityZone"`
	IsAZ				bool	`json:"is_az"`
	Env					string `json:"enviroment,omitempty"`
	AccountID			string `json:"account_id,omitempty"`
}

type Server struct {
	Port 			int `json:"port"`
	ReadTimeout		int `json:"readTimeout"`
	WriteTimeout	int `json:"writeTimeout"`
	IdleTimeout		int `json:"idleTimeout"`
	CtxTimeout		int `json:"ctxTimeout"`
}

type MessageRouter struct {
	Message			string `json:"message"`
}

type Authentication struct {
	Token			string	`json:"token,omitempty"`
	TokenEncrypted	string	`json:"token_encrypted,omitempty"`
	ExpirationTime	time.Time `json:"expiration_time,omitempty"`
	ApiKey			string	`json:"api_key,omitempty"`
}

type Credential struct {
	ID				string	`json:"ID,omitempty"`
	SK				string	`json:"SK,omitempty"`
	User			string	`json:"user,omitempty"`
	Password		string	`json:"password,omitempty"`
	Token			string 	`json:"token,omitempty"`
	UsagePlan		string 	`json:"usage_plan,omitempty"`
	ApiKey			string 	`json:"apikey,omitempty"`
	Updated_at  	time.Time 	`json:"updated_at,omitempty"`
	CredentialScope	*CredentialScope `json:"credential_scope,omitempty"`
	JwtKeySign		interface{}
	JwtKeyCreation	interface{}
}

type CredentialScope struct {
	ID				string		`json:"ID"`
	SK				string		`json:"SK"`
	User			string		`json:"user,omitempty"`
	Scope			[]string	`json:"scope,omitempty"`
	Updated_at  	time.Time 	`json:"updated_at,omitempty"`
}

type JwtData struct {
	TokenUse	string 	`json:"token_use"`
	ISS			string 	`json:"iss"`
	Version		string 	`json:"version"`
	JwtId		string 	`json:"jwt_id"`
	Username	string 	`json:"username"`
	Scope	  	[]string `json:"scope"`
	jwt.RegisteredClaims
}

type AwsService struct {
	AwsRegion			string `json:"aws_region"`
	DynamoTableName		string `json:"dynamo_table_name"`
	SecretName			string `json:"secret_name"`
}

type RsaKey struct{
	JwtKey				string
	Key_rsa_priv_pem	string
	Key_rsa_pub_pem 	string
	Crl_pem 			string
	CaCert				string 	`json:"ca_cert,omitempty"`
	Key_rsa_priv 		*rsa.PrivateKey
	Key_rsa_pub 		*rsa.PublicKey	
}

type Jwks struct{
	Keys		[]JKey 	`json:"keys"`
}

type JKey struct{
	Type		string 	`json:"kty"`
	Algorithm	string 	`json:"alg"`
	JwtId		string 	`json:"kid"`
	NBase64		string 	`json:"n"`
}

type JwksData struct {
	Token			string 	`json:"token,omitempty"`
	JwtId			string 	`json:"kid,omitempty"`
	RSAPublicKeyB64	string 	`json:"rsa_public_key_b64"`
}