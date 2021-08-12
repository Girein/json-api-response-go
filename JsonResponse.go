package pretty

// JSONResponse format
type JSONResponse struct {
	StatusCode string         `json:"status_code"`
	StatusText string         `json:"status_text"`
	Message    string         `json:"message"`
	Data       interface{}    `json:"data"`
	Paginator  map[string]int `json:"paginator"`
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

// WithData Response
func (r *JSONResponse) WithPagination(statusCode string, statusText string, message string, data interface{}, paginator map[string]int) {
	r.StatusCode = statusCode
	r.StatusText = statusText
	r.Message = message
	r.Data = data
	r.Paginator = map[string]int{
		"per_page":      paginator["PageSize"],
		"current_page":  paginator["Page"],
		"previous_page": paginator["Page"] - 1,
		"next_page":     paginator["Page"] + 1,
	}
}
