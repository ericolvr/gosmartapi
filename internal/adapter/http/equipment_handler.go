package http

import (
	"net/http"

	"github.com/ericolvr/goapi/internal/domain"
	"github.com/ericolvr/goapi/internal/usecase"
	"github.com/gin-gonic/gin"
)

type equipmentHandler struct {
	equipmentUsecase usecase.EquipmentUsecase
}

func NewEquipmentHandler(router *gin.Engine, equipmentUsecase usecase.EquipmentUsecase) {
	handler := &equipmentHandler{
		equipmentUsecase: equipmentUsecase,
	}

	router.POST("/equipments", handler.createEquipment)
}

func (h *equipmentHandler) createEquipment(c *gin.Context) {
	var equipment domain.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.equipmentUsecase.CreateEquipment(&equipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, equipment)
}
