package controller

import (
	"net/http"
	"strconv"
	"sync"
	"todolist/model"

	"github.com/gin-gonic/gin"
)

type ChecklistController struct {
	Checklists []model.Checklist
	mu         sync.Mutex
}

func NewChecklistController() *ChecklistController {
	return &ChecklistController{}
}

func (cc *ChecklistController) CreateChecklist(c *gin.Context) {
	var req model.Checklist
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.ID = len(cc.Checklists) + 1

	cc.mu.Lock()
	defer cc.mu.Unlock()

	cc.Checklists = append(cc.Checklists, req)

	c.JSON(http.StatusCreated, gin.H{"message": "Checklist created successfully", "data": req})
}

func (cc *ChecklistController) GetAllChecklists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": cc.Checklists})
}

func (cc *ChecklistController) DeleteChecklist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist ID"})
		return
	}

	cc.mu.Lock()
	defer cc.mu.Unlock()

	if id <= 0 || id > len(cc.Checklists) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checklist not found"})
		return
	}

	cc.Checklists = append(cc.Checklists[:id-1], cc.Checklists[id:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Checklist deleted successfully"})
}
