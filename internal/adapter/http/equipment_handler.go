package http

import (
	"net/http"
	"strconv"

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
	router.GET("/equipments", handler.getEquipments)
	router.GET("/equipments/:id", handler.getEquipment)
	router.PUT("/equipments/:id", handler.updateEquipment)
	router.DELETE("/users/:id", handler.deleteEquipment)
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

func (h *equipmentHandler) getEquipments(c *gin.Context) {
	users, err := h.equipmentUsecase.GetEquipments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *equipmentHandler) getEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	equipment, err := h.equipmentUsecase.GetEquipmentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if equipment == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipment not found"})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

func (h *equipmentHandler) updateEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	var equipment domain.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	equipment.ID = id

	if err := h.equipmentUsecase.UpdateEquipment(&equipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

func (h *equipmentHandler) deleteEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	if err := h.equipmentUsecase.DeleteEquipment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipment deleted successfully"})
}
