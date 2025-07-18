package server

import (
	"time"
	"encoding/json"
	"net/http"
	"strconv"
	"os"
	"os/signal"
	"syscall"
	"context"

	"github.com/go-oauth/internal/adapter/api"
	"github.com/go-oauth/internal/core/model"
	go_core_observ "github.com/eliezerraj/go-core/observability"  
	"github.com/eliezerraj/go-core/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

var childLogger = log.With().Str("component","go-oauth").Str("package","internal.infra.server").Logger()

var core_middleware middleware.ToolsMiddleware
var tracerProvider go_core_observ.TracerProvider
var infoTrace go_core_observ.InfoTrace

type HttpServer struct {
	httpServer	*model.Server
}

// About create a new http server
func NewHttpAppServer(httpServer *model.Server) HttpServer {
	childLogger.Info().Str("func","NewHttpAppServer").Send()

	return HttpServer{httpServer: httpServer }
}

// About start http server
func (h HttpServer) StartHttpAppServer(	ctx context.Context, 
										httpRouters *api.HttpRouters,
										appServer *model.AppServer) {
childLogger.Info().Str("func","StartHttpAppServer").Send()
			
	// otel
	infoTrace.PodName = appServer.InfoPod.PodName
	infoTrace.PodVersion = appServer.InfoPod.ApiVersion
	infoTrace.ServiceType = "k8-workload"
	infoTrace.Env = appServer.InfoPod.Env
	infoTrace.AccountID = appServer.InfoPod.AccountID

	tp := tracerProvider.NewTracerProvider(	ctx, 
											appServer.ConfigOTEL, 
											&infoTrace)

	if tp != nil {
		otel.SetTextMapPropagator(xray.Propagator{})
		otel.SetTracerProvider(tp)
	}

	// handle defer
	defer func() { 
		if tp != nil {
			err := tp.Shutdown(ctx)
			if err != nil{
				childLogger.Error().Err(err).Send()
			}
		}
		childLogger.Info().Msg("stop done !!!")
	}()

	// router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(core_middleware.MiddleWareHandlerHeader)

	myRouter.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		childLogger.Info().Str("HandleFunc","/").Send()
		
		json.NewEncoder(rw).Encode(appServer)
	})

	health := myRouter.Methods(http.MethodGet, http.MethodOptions).Subrouter()
    health.HandleFunc("/health", httpRouters.Health)

	live := myRouter.Methods(http.MethodGet, http.MethodOptions).Subrouter()
    live.HandleFunc("/live", httpRouters.Live)

	header := myRouter.Methods(http.MethodGet, http.MethodOptions).Subrouter()
    header.HandleFunc("/header", httpRouters.Header)

	myRouter.HandleFunc("/info", func(rw http.ResponseWriter, req *http.Request) {
		childLogger.Info().Str("HandleFunc","/info").Send()

		rw.Header().Set("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(appServer)
	})
	
	oauthCredential := myRouter.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	oauthCredential.Handle("/oauth_credential",core_middleware.MiddleWareErrorHandler(httpRouters.OAUTHCredential),)
	oauthCredential.Use(otelmux.Middleware("go-oauth"))

	oauthCredentialHs256 := myRouter.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	oauthCredentialHs256.Handle("/oauth_credential_hs256",core_middleware.MiddleWareErrorHandler(httpRouters.OAUTHCredential),)
	oauthCredentialHs256.Use(otelmux.Middleware("go-oauth"))

	validToken := myRouter.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	validToken.Handle("/tokenValidation/{id}",core_middleware.MiddleWareErrorHandler(httpRouters.TokenValidation),)
	validToken.Use(otelmux.Middleware("go-oauth"))

	validTokenHs256 := myRouter.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	validTokenHs256.Handle("/tokenValidation_hs256/{id}",core_middleware.MiddleWareErrorHandler(httpRouters.TokenValidation),)
	validTokenHs256.Use(otelmux.Middleware("go-oauth"))

	refreshToken := myRouter.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	refreshToken.Handle("/refresh_token",core_middleware.MiddleWareErrorHandler(httpRouters.RefreshToken),)
	refreshToken.Use(otelmux.Middleware("go-oauth"))

	refreshTokenHs256 := myRouter.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	refreshTokenHs256.Handle("/refresh_token_hs256",core_middleware.MiddleWareErrorHandler(httpRouters.RefreshToken),)
	refreshTokenHs256.Use(otelmux.Middleware("go-oauth"))

	wellKnown := myRouter.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	wellKnown.Handle("/wellKnown/1",core_middleware.MiddleWareErrorHandler(httpRouters.WellKnown),)
	wellKnown.Use(otelmux.Middleware("go-oauth"))

	ValidationTokenSignedPubKey := myRouter.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	ValidationTokenSignedPubKey.Handle("/validationTokenSignedPubKey",core_middleware.MiddleWareErrorHandler(httpRouters.ValidationTokenSignedPubKey),)
	ValidationTokenSignedPubKey.Use(otelmux.Middleware("go-oauth"))	

	verifyCertCRL := myRouter.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	verifyCertCRL.Handle("/verifyCertCRL",core_middleware.MiddleWareErrorHandler(httpRouters.VerifyCertCRL),)
	verifyCertCRL.Use(otelmux.Middleware("go-oauth"))	

	signIn := myRouter.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	signIn.Handle("/signIn",core_middleware.MiddleWareErrorHandler(httpRouters.SignIn),)
	signIn.Use(otelmux.Middleware("go-oauth"))	

	addScope := myRouter.Methods(http.MethodPost, http.MethodOptions).Subrouter()
	addScope.Handle("/add/acope",core_middleware.MiddleWareErrorHandler(httpRouters.AddScope),)
	addScope.Use(otelmux.Middleware("go-oauth"))	

	getCredential := myRouter.Methods(http.MethodGet, http.MethodOptions).Subrouter()
	getCredential.Handle("/credential/{id}",core_middleware.MiddleWareErrorHandler(httpRouters.GetCredential),)
	getCredential.Use(otelmux.Middleware("go-oauth"))

	// setup http server
	srv := http.Server{
		Addr:         ":" +  strconv.Itoa(h.httpServer.Port),      	
		Handler:      myRouter,                	          
		ReadTimeout:  time.Duration(h.httpServer.ReadTimeout) * time.Second,   
		WriteTimeout: time.Duration(h.httpServer.WriteTimeout) * time.Second,  
		IdleTimeout:  time.Duration(h.httpServer.IdleTimeout) * time.Second, 
	}

	childLogger.Info().Str("Service Port", strconv.Itoa(h.httpServer.Port)).Send()

	// start http server
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			childLogger.Error().Err(err).Msg("canceling http mux server !!!")
		}
	}()

	// handle sig TERM
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	if err := srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		childLogger.Error().Err(err).Msg("warning dirty shutdown !!!")
		return
	}
}