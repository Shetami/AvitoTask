package response

import (
	"AvitoTask/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	admin_api := router.Group("/api/admin")
	{
		admin_api.POST("/confirmation", h.confirmation)
	}

	api := router.Group("/api")
	{
		api.POST("/pay", h.pay)
		api.POST("/transfer", h.transferMoney)
		api.POST("/add", h.addMoney)
		api.GET("/balance/:id", h.getBalance)
	}
	return router
}
