package studioapi

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// AbilitiesAPIController binds http requests to an api service and writes the service results to the http response
type AbilitiesAPIController struct {
	service      AbilitiesAPIServicer
	errorHandler ErrorHandler
}

// AbilitiesAPIOption for how the controller is set up.
type AbilitiesAPIOption func(*AbilitiesAPIController)

// WithAbilitiesAPIErrorHandler inject ErrorHandler into controller
func WithAbilitiesAPIErrorHandler(h ErrorHandler) AbilitiesAPIOption {
	return func(c *AbilitiesAPIController) {
		c.errorHandler = h
	}
}

// NewAbilitiesAPIController creates a default api controller
func NewAbilitiesAPIController(s AbilitiesAPIServicer, opts ...AbilitiesAPIOption) *AbilitiesAPIController {
	controller := &AbilitiesAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the AbilitiesAPIController
func (c *AbilitiesAPIController) Routes() Routes {
	return Routes{
		"GetAbilityDetails": Route{
			strings.ToUpper("Get"),
			"/api/abilities/{symbol}",
			c.GetAbilityDetails,
		},
		"GetAbilities": Route{
			strings.ToUpper("Get"),
			"/api/abilities",
			c.GetAbilities,
		},
	}
}

// GetAbilityDetails - Get an ability details
func (c *AbilitiesAPIController) GetAbilityDetails(w http.ResponseWriter, r *http.Request) {
	var (
		params = r.URL.Query()
		lang   = params.Get("lang")
		symbol = chi.URLParam(r, "symbol")
	)

	result, err := c.service.GetAbilityDetails(r.Context(), symbol, lang)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}

	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetAbilities - Get a page of abilities
func (c *AbilitiesAPIController) GetAbilities(w http.ResponseWriter, r *http.Request) {
	var (
		params = r.URL.Query()
		lang   = params.Get("lang")
	)

	result, err := c.service.GetAbilities(r.Context(), lang)
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}

	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
