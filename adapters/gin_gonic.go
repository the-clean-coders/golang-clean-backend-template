package adapters

import (
	"my_module_path/my_module_name/adapters/endpoints"
	"my_module_path/my_module_name/entities"
	"my_module_path/my_module_name/services/example"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinHandler represents a gin gonic handler.
type GinHandler struct {
	Usecase example.UseCase
}

func NewGinHandler(usecase example.UseCase, jwtSecret string) *gin.Engine {
	h := &GinHandler{
		Usecase: usecase,
	}

	r := gin.Default()
	r.GET(endpoints.GetExample, h.GetExample)
	return r
}

type getExampleResponse entities.Example

// GetExample is the gin HandlerFunc for GetExample endpoint.
func (h *GinHandler) GetExample(ctx *gin.Context) {
	exampleID := ctx.Param("example_id")

	ex, err := h.Usecase.GetExample(exampleID)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, getExampleResponse(ex))
}
