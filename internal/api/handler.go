package api

import (
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nurcholisnanda/golang-assignment/internal/dto"
	"github.com/nurcholisnanda/golang-assignment/internal/service"
)

// handler represents the API handler.
type handler struct {
	svc service.ServiceInterface
}

// NewHandler creates a new API handler with the provided service interface.
func NewHandler(svc service.ServiceInterface) *handler {
	return &handler{
		svc: svc,
	}
}

// GetRecords handles the HTTP GET request to fetch records based on query parameters.
func (h *handler) GetRecords(ctx *gin.Context) {

	// Calculate today's date in YYYY-MM-DD format
	nowStr := time.Now().Format(time.DateOnly)

	// Default request parameters
	req := &dto.FetchRecordsRequest{
		StartDate: "1900-01-01",
		EndDate:   nowStr,
		MinCount:  0,
		MaxCount:  math.MaxInt,
	}

	// Override request parameters with query parameters if provided
	if startDate, ok := ctx.GetQuery("startDate"); ok {
		req.StartDate = startDate
	}
	if endDate, ok := ctx.GetQuery("endDate"); ok {
		req.EndDate = endDate
	}
	if minCount, ok := ctx.GetQuery("minCount"); ok {
		temp, err := strconv.Atoi(minCount)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.FetchRecordsResponse{
				Code: dto.ErrBadRequest,
				Msg:  err.Error(),
			})
			return
		}
		req.MinCount = temp
	}
	if maxCount, ok := ctx.GetQuery("maxCount"); ok {
		temp, err := strconv.Atoi(maxCount)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.FetchRecordsResponse{
				Code: dto.ErrBadRequest,
				Msg:  err.Error(),
			})
			return
		}
		req.MaxCount = temp
	}

	// Call service layer to fetch records based on the request
	res, err := h.svc.GetRecords(req)
	if err != nil {
		var statusCode int
		var errMsg string

		// Determine appropriate HTTP status code and error message based on the error
		if strings.Contains(err.Error(), "not found") {
			statusCode = http.StatusNotFound
			errMsg = "Records not found: " + err.Error()
		} else {
			statusCode = http.StatusBadRequest
			errMsg = "Bad request: " + err.Error()
		}

		// Return error response to the client
		ctx.AbortWithStatusJSON(statusCode, dto.FetchRecordsResponse{
			Code: dto.ErrUnknown,
			Msg:  errMsg,
		})
		return
	}

	// Return successful response with records to the client
	ctx.JSON(http.StatusOK, res)
}
