package commonsdtos

import "fmt"

type APIError struct {
	Code        string                 `json:"codigo"`
	Description string                 `json:"descripcion"`
	Details     map[string]interface{} `json:"-"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Description)
}
