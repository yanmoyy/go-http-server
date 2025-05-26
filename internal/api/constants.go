package api

// Endpoints defines the API endpoint paths
const (
	// Endpoints

	EndpointReset   = "/admin/reset"
	EndpointUsers   = "/api/users"
	EndpointChirps  = "/api/chirps"
	EndpointMetrics = "/admin/metrics"
	EndpointHealthz = "/api/healthz"
	EndpointLogin   = "/api/login"
	EndpointRefresh = "/api/refresh"
	EndpointRevoke  = "/api/revoke"

	// Params
	ChirpIDParam = "chirpID"
	BearerPrefix = "Bearer "

	// Header
	HeaderContentType     = "Content-Type"
	HeaderAuthorization   = "Authorization"
	HeaderApplicationJson = "application/json"
)
