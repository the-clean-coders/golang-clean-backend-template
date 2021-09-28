package adapters

import (
	"my_module_path/my_module_name/adapters/endpoints"
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
	r.POST(endpoints.AddExample, h.GetExample)
	r.GET(endpoints.GetExample, h.GetExample)
	return r
}

// AddExample is the gin HandlerFunc for AddExample endpoint.
func (h *GinHandler) AddExample(ctx *gin.Context) {
	var params example.AddExampleParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Error: "Cannot parse new example data",
		})
		return
	}

	response, err := h.Usecase.AddExample(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			Error: "Cannot add new example",
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// GetExample is the gin HandlerFunc for GetExample endpoint.
func (h *GinHandler) GetExample(ctx *gin.Context) {
	params := example.GetExampleParams{
		ExampleID: ctx.Param("example_id"),
	}

	ex, err := h.Usecase.GetExample(params)
	if ex == nil {
		ctx.JSON(http.StatusNotFound, &ErrorResponse{
			Error: "Not found",
		})
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			Error: "Internal Server Error",
		})
	}

	ctx.JSON(http.StatusOK, ex)
}
