package utilities

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// LoadPayload loads the payload from the request into an interface
func LoadPayload(c *gin.Context) (interface{}, error) {
	var payload interface{}
	// Load the request's body into the payload
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
