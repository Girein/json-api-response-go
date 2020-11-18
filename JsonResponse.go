package pretty

// JSONResponse format
type JSONResponse struct {
	StatusCode string      `json:"status_code"`
	StatusText string      `json:"status_text"`
	Message    string      `json:"message"`
	Error      error       `json:"error"`
	Data       interface{} `json:"data"`
}

// Simple Response
func (r *JSONResponse) Simple(statusCode string, statusText string, message string) {
	r.StatusCode = statusCode
	r.StatusText = statusText
	r.Message = message
}

// WithData Response
func (r *JSONResponse) WithData(statusCode string, statusText string, message string, data interface{}) {
	r.StatusCode = statusCode
	r.StatusText = statusText
	r.Message = message
	r.Data = data
}
