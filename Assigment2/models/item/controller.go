package item

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"assigment2/helper" // Assuming helper package provides formatting functions
)

type itemController struct {
	service Service
}

func NewItemController(service Service) *itemController {
	return &itemController{service: service}
}

func (h *itemController) Get(c *gin.Context) {
	items, err := h.service.Get()
	if err != nil {
		// Handle error with proper status code and message
		h.handleError(c, http.StatusInternalServerError, "Error retrieving items", err)
		return
	}

	formattedItems := helper.ItemsFormatter(items) // Assuming ItemsFormatter is defined
	response := helper.APIResponse("Get Items Success", http.StatusOK, "success", formattedItems)
	c.JSON(http.StatusOK, response)
}

func (h *itemController) Create(c *gin.Context) {
	var input CreateItemInput // Use specific input struct
	err := c.ShouldBindJSON(&input)
	if err != nil {
		// Handle validation error with proper status code and message
		h.handleError(c, http.StatusBadRequest, "Invalid item data", err)
		return
	}

	item, err := h.service.Create(input)
	if err != nil {
		// Handle creation error with proper status code and message
		h.handleError(c, http.StatusInternalServerError, "Error creating item", err)
		return
	}

	formattedItem := helper.ItemFormatter(item) // Assuming ItemFormatter is defined
	response := helper.APIResponse("Create Item Success", http.StatusOK, "success", formattedItem)
	c.JSON(http.StatusOK, response)
}

// handleError provides a centralized way to handle errors with proper responses
func (h *itemController) handleError(c *gin.Context, statusCode int, message string, err error) {
	errorMessage := gin.H{"errors": helper.FormatValidationError(err)} // Assuming helper function exists
	response := helper.APIResponse(message, statusCode, "error", errorMessage)
	c.JSON(statusCode, response)
}
