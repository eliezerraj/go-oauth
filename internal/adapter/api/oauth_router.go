package api

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"github.com/rs/zerolog/log"
	"github.com/go-oauth/internal/core/service"
	"github.com/go-oauth/internal/core/model"
	"github.com/go-oauth/internal/core/erro"
	go_core_observ "github.com/eliezerraj/go-core/observability"
	go_core_tools "github.com/eliezerraj/go-core/tools"
	"github.com/eliezerraj/go-core/coreJson"
	"github.com/gorilla/mux"
)

var childLogger = log.With().Str("component", "go-outhh").Str("package", "internal.adapter.api").Logger()

var core_json coreJson.CoreJson
var core_apiError coreJson.APIError
var core_tools go_core_tools.ToolsCore
var tracerProvider go_core_observ.TracerProvider

type HttpRouters struct {
	workerService 	*service.WorkerService
}

// Above setup the type model of jwt key signature
func (h *HttpRouters) setSignModel(model string, credential *model.Credential){
	childLogger.Info().Str("func","setSignModel").Send()

	if model == "HS256" {
		credential.JwtKeySign = h.workerService.Keys.JwtKey
		credential.JwtKeyCreation = h.workerService.Keys.JwtKey
		h.workerService.TokenSignedValidation = service.TokenValidationHS256
		h.workerService.CreatedToken = service.CreatedTokenHS256
	} else {
		credential.JwtKeySign = h.workerService.Keys.Key_rsa_pub
		credential.JwtKeyCreation = h.workerService.Keys.Key_rsa_priv
		h.workerService.TokenSignedValidation = service.TokenValidationRSA
		h.workerService.CreatedToken = service.CreatedTokenRSA
	}
}

func NewHttpRouters(workerService *service.WorkerService) HttpRouters {
	childLogger.Info().Str("func","NewHttpRouters").Send()

	return HttpRouters{
		workerService: workerService,
	}
}

// About return a health
func (h *HttpRouters) Health(rw http.ResponseWriter, req *http.Request) {
	childLogger.Info().Str("func","Health").Interface("trace-resquest-id", req.Context().Value("trace-request-id")).Send()

	json.NewEncoder(rw).Encode(model.MessageRouter{Message: "true"})
}

// About return a live
func (h *HttpRouters) Live(rw http.ResponseWriter, req *http.Request) {
	childLogger.Debug().Msg("Live")

	json.NewEncoder(rw).Encode(model.MessageRouter{Message: "true"})
}

// About show all header received
func (h *HttpRouters) Header(rw http.ResponseWriter, req *http.Request) {
	childLogger.Info().Str("func","Header").Interface("trace-resquest-id", req.Context().Value("trace-request-id")).Send()

	json.NewEncoder(rw).Encode(req.Header)
}

// About login using symetric method  
func (h *HttpRouters) OAUTHCredential(rw http.ResponseWriter, req *http.Request) error {
	childLogger.Info().Str("func","OAUTHCredential").Interface("trace-resquest-id", req.Context().Value("trace-request-id")).Send()

	//trace
	span := tracerProvider.Span(req.Context(), "adapter.api.OAUTHCredential")
	defer span.End()

	trace_id := fmt.Sprintf("%v",req.Context().Value("trace-request-id"))

	// prepare body
	credential := model.Credential{}
	err := json.NewDecoder(req.Body).Decode(&credential)
    if err != nil {
		core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusBadRequest)
		return &core_apiError
    }
	defer req.Body.Close()

	// Check which type of authentication method 
	if req.RequestURI == "/oauth_credential_hs256" {
		h.setSignModel("HS256", &credential)
	} else {
		h.setSignModel("RSA", &credential)
	}

	//call service
	res, err := h.workerService.OAUTHCredential(req.Context(), credential)
	if err != nil {
		switch err {
		case erro.ErrNotFound:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusNotFound)
		default:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusInternalServerError)
		}
		return &core_apiError
	}
	
	return core_json.WriteJSON(rw, http.StatusOK, res)
}

// About check a token expitation date
func (h *HttpRouters) TokenValidation(rw http.ResponseWriter, req *http.Request) error {
	childLogger.Info().Str("func","TokenValidation").Interface("trace-resquest-id", req.Context().Value("trace-request-id")).Send()

	//trace
	span := tracerProvider.Span(req.Context(), "adapter.api.TokenValidation")
	defer span.End()
	trace_id := fmt.Sprintf("%v",req.Context().Value("trace-request-id"))

	//parameters
	vars := mux.Vars(req)
	varID := vars["id"]

	credential := model.Credential{}
	credential.Token = varID

	// Check which type of authentication method and insert the rigth function
	if strings.Contains(req.RequestURI, "/tokenValidation_hs256/") {
		h.setSignModel("HS256", &credential)
	} else {
		h.setSignModel("RSA", &credential)
	}

	//call service
	res, err := h.workerService.TokenValidation(req.Context(), credential)
	if err != nil {
		switch err {
		case erro.ErrTokenExpired:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusUnauthorized)
		case erro.ErrStatusUnauthorized:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusUnauthorized)
		default:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusInternalServerError)
		}
		return &core_apiError
	}
	
	return core_json.WriteJSON(rw, http.StatusOK, res)
}

// About refresh token
func (h *HttpRouters) RefreshToken(rw http.ResponseWriter, req *http.Request) error {
	childLogger.Info().Str("func","RefreshToken").Interface("trace-resquest-id", req.Context().Value("trace-request-id")).Send()

	//trace
	span := tracerProvider.Span(req.Context(), "adapter.api.RefreshToken")
	defer span.End()

	trace_id := fmt.Sprintf("%v",req.Context().Value("trace-request-id"))

	// prepare body
	credential := model.Credential{}
	err := json.NewDecoder(req.Body).Decode(&credential)
    if err != nil {
		core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusBadRequest)
		return &core_apiError
    }
	defer req.Body.Close()

	// Check which type of authentication method and insert the rigth function
	if strings.Contains(req.RequestURI, "/refresh_token_hs256") {
		h.setSignModel("HS256", &credential)
	} else {
		h.setSignModel("RSA", &credential)
	}

	//call service
	res, err := h.workerService.RefreshToken(req.Context(), credential)
	if err != nil {
		switch err {
		case erro.ErrNotFound:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusNotFound)
		default:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusInternalServerError)
		}
		return &core_apiError
	}
	
	return core_json.WriteJSON(rw, http.StatusOK, res)
}

// About wellknow service
func (h *HttpRouters) WellKnown(rw http.ResponseWriter, req *http.Request) error {
	childLogger.Info().Str("func","WellKnown").Interface("trace-resquest-id", req.Context().Value("trace-request-id")).Send()

	//trace
	span := tracerProvider.Span(req.Context(), "adapter.api.WellKnown")
	defer span.End()

	trace_id := fmt.Sprintf("%v",req.Context().Value("trace-request-id"))

	//call service
	res, err := h.workerService.WellKnown(req.Context())
	if err != nil {
		switch err {
		case erro.ErrNotFound:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusNotFound)
		default:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusInternalServerError)
		}
		return &core_apiError
	}
	
	return core_json.WriteJSON(rw, http.StatusOK, res)
}

// About validate token was signed with a pubkey
func (h *HttpRouters) ValidationTokenSignedPubKey(rw http.ResponseWriter, req *http.Request) error {
	childLogger.Info().Str("func","ValidationTokenSignedPubKey").Interface("trace-resquest-id", req.Context().Value("trace-request-id")).Send()

	//trace
	span := tracerProvider.Span(req.Context(), "adapter.api.ValidationTokenSignedPubKey")
	defer span.End()

	trace_id := fmt.Sprintf("%v",req.Context().Value("trace-request-id"))

	jwksData := model.JwksData{}
	err := json.NewDecoder(req.Body).Decode(&jwksData)
    if err != nil {
		core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusBadRequest)
		return &core_apiError
    }
	defer req.Body.Close()

	//call service
	res, err := h.workerService.ValidationTokenSignedPubKey(req.Context(), jwksData)
	if err != nil {
		switch err {
		case erro.ErrNotFound:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusNotFound)
		default:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusInternalServerError)
		}
		return &core_apiError
	}
	
	return core_json.WriteJSON(rw, http.StatusOK, res)
}

// About checj a crl list
func (h *HttpRouters) VerifyCertCRL(rw http.ResponseWriter, req *http.Request) error {
	childLogger.Info().Str("func","VerifyCertCRL").Interface("trace-resquest-id", req.Context().Value("trace-request-id")).Send()

	//trace
	span := tracerProvider.Span(req.Context(), "adapter.api.VerifyCertCRL")
	defer span.End()

	trace_id := fmt.Sprintf("%v",req.Context().Value("trace-request-id"))

	caCert := model.RsaKey{}
	err := json.NewDecoder(req.Body).Decode(&caCert)
    if err != nil {
		core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusBadRequest)
		return &core_apiError
    }
	defer req.Body.Close()

	//call service
	res, err := h.workerService.VerifyCertCRL(req.Context(), caCert.CaCert)
	if err != nil {
		switch err {
		case erro.ErrNotFound:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusNotFound)
		default:
			core_apiError = core_apiError.NewAPIError(err, trace_id, http.StatusInternalServerError)
		}
		return &core_apiError
	}
	
	return core_json.WriteJSON(rw, http.StatusOK, res)
}