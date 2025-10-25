package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/anvarmx/tranzify-api/services"
	"github.com/anvarmx/tranzify-api/types"
	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	response := types.APIResponse{
		Status: "success",
		Result: types.RootResult{
			Message:  "Tranzify API is working! 🚀",
			Version:  "v1.0.0",
			DateTime: time.Now().Format(time.RFC3339),
		},
	}

	c.JSON(http.StatusOK, response)
}

func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, types.APIResponse{
		Status: "error",
		Result: types.ErrorResponse{
			Message: "Invalid request!",
		},
	})
}

func MorseToTextHandler(c *gin.Context) {
	var req types.MorseRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Status: "error",
			Result: types.ErrorResponse{
				Message: "Invalid request body",
			},
		})

		return
	}

	if strings.TrimSpace(req.Input) == "" {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Status: "error",
			Result: types.ErrorResponse{
				Message: "Input cannot be empty",
			},
		})

		return
	}

	if len(req.Input) > 1024 {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Status: "error",
			Result: types.ErrorResponse{
				Message: "Input text exceeds 1024 character limit",
			},
		})

		return
	}

	input := req.Input
	output := services.MorseToText(input)

	c.JSON(http.StatusOK, types.APIResponse{
		Status: "success",
		Result: types.ConversionResult{
			Input:  input,
			Output: output,
			Length: len([]rune(input)),
		},
	})
}

func TextToMorseHandler(c *gin.Context) {
	var req types.MorseRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Status: "error",
			Result: gin.H{"message": "Invalid request body"},
		})

		return
	}

	if strings.TrimSpace(req.Input) == "" {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Status: "error",
			Result: types.ErrorResponse{
				Message: "Input cannot be empty",
			},
		})

		return
	}

	input := req.Input

	if len(input) > 1024 {
		c.JSON(http.StatusBadRequest, types.APIResponse{
			Status: "error",
			Result: gin.H{"message": "Input text exceeds 1024 character limit"},
		})

		return
	}

	output := services.TextToMorse(input)

	c.JSON(http.StatusOK, types.APIResponse{
		Status: "success",
		Result: types.ConversionResult{
			Input:  input,
			Output: output,
			Length: len([]rune(input)),
		},
	})
}
