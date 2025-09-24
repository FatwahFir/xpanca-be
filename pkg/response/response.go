// pkg/response/response.go
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
}

type ErrorPayload struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type Envelope struct {
	Success    bool          `json:"success"`
	Data       interface{}   `json:"data,omitempty"`
	Error      *ErrorPayload `json:"error,omitempty"`
	Pagination *Pagination   `json:"pagination,omitempty"`
	Meta       gin.H         `json:"meta,omitempty"`
}

func meta(c *gin.Context) gin.H {
	if id, ok := c.Get("request_id"); ok {
		return gin.H{"request_id": id}
	}
	return nil
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Envelope{Success: true, Data: data, Meta: meta(c)})
}

func Paginated(c *gin.Context, data interface{}, p Pagination) {
	c.JSON(http.StatusOK, Envelope{Success: true, Data: data, Pagination: &p, Meta: meta(c)})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Envelope{Success: true, Data: data, Meta: meta(c)})
}

func Err(c *gin.Context, status int, code, msg string, details interface{}) {
	c.JSON(status, Envelope{
		Success: false,
		Error:   &ErrorPayload{Code: code, Message: msg, Details: details},
		Meta:    meta(c),
	})
}
