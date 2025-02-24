package app

import (
	"github.com/PainCodermax/to-do-list-api/pkg/errcode"
	"github.com/PainCodermax/to-do-list-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Validation(c *gin.Context, body any, response *Response) error {
	if err := c.ShouldBind(body); err != nil {
		logger.WithTrace(c).Errorf("bodys errs: %v", err)
		response.ToErrorResponse(errcode.InvalidRequestBody.WithDetails(err.Error()))
		return err
	}
	return nil
}
