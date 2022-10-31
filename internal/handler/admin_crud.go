package response

import (
	"AvitoTask/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) confirmation(c *gin.Context) {
	var input models.ConfirmationStruct
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Admin.Confirmation(input.Id, input.Value); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}
