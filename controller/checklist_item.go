package controller

import (
	"net/http"
	"strconv"
	"sync"
	"todolist/model"

	"github.com/gin-gonic/gin"
)

type ChecklistItemController struct {
	ChecklistItems []model.ChecklistItem
	mu             sync.Mutex
}

func NewChecklistItemController() *ChecklistItemController {
	return &ChecklistItemController{}
}

func (cic *ChecklistItemController) CreateChecklistItem(c *gin.Context) {
	checklistID, err := strconv.Atoi(c.Param("checklistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist ID"})
		return
	}

	var req struct {
		ItemName string `json:"itemName" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := model.ChecklistItem{
		ID:          len(cic.ChecklistItems) + 1,
		ChecklistID: checklistID,
		ItemName:    req.ItemName,
	}

	cic.mu.Lock()
	defer cic.mu.Unlock()

	cic.ChecklistItems = append(cic.ChecklistItems, item)

	c.JSON(http.StatusCreated, gin.H{"message": "Checklist item created successfully", "data": item})
}

func (cic *ChecklistItemController) GetAllChecklistItemsByChecklistID(c *gin.Context) {
	checklistID, err := strconv.Atoi(c.Param("checklistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist ID"})
		return
	}

	var items []model.ChecklistItem

	for _, item := range cic.ChecklistItems {
		if item.ChecklistID == checklistID {
			items = append(items, item)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (cic *ChecklistItemController) GetChecklistItemByChecklistIDAndItemID(c *gin.Context) {
	checklistID, err := strconv.Atoi(c.Param("checklistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist ID"})
		return
	}

	itemID, err := strconv.Atoi(c.Param("checklistItemId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist item ID"})
		return
	}

	cic.mu.Lock()
	defer cic.mu.Unlock()

	var checklistItem model.ChecklistItem
	for _, item := range cic.ChecklistItems {
		if item.ID == itemID && item.ChecklistID == checklistID {
			checklistItem = item
			break
		}
	}

	if checklistItem.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checklist item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": checklistItem})
}

func (cic *ChecklistItemController) UpdateChecklistItemStatus(c *gin.Context) {
	checklistID, err := strconv.Atoi(c.Param("checklistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist ID"})
		return
	}

	itemID, err := strconv.Atoi(c.Param("checklistItemId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist item ID"})
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cic.mu.Lock()
	defer cic.mu.Unlock()

	var updated bool
	for i := range cic.ChecklistItems {
		if cic.ChecklistItems[i].ID == itemID && cic.ChecklistItems[i].ChecklistID == checklistID {
			cic.ChecklistItems[i].Status = req.Status
			updated = true
			break
		}
	}

	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checklist item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checklist item status updated successfully"})
}

func (cic *ChecklistItemController) RenameChecklistItem(c *gin.Context) {
	checklistID, err := strconv.Atoi(c.Param("checklistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist ID"})
		return
	}

	itemID, err := strconv.Atoi(c.Param("checklistItemId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checklist item ID"})
		return
	}

	var req struct {
		ItemName string `json:"itemName" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cic.mu.Lock()
	defer cic.mu.Unlock()

	var updated bool
	for i := range cic.ChecklistItems {
		if cic.ChecklistItems[i].ID == itemID && cic.ChecklistItems[i].ChecklistID == checklistID {
			cic.ChecklistItems[i].ItemName = req.ItemName
			updated = true
			break
		}
	}

	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checklist item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checklist item renamed successfully"})
}
