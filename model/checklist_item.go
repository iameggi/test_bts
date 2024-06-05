package model

type ChecklistItem struct {
	ID          int    `json:"id"`
	ChecklistID int    `json:"checklist_id"`
	ItemName    string `json:"item_name"`
	Status      string `json:"status"`
}

func NewChecklistItem(id, checklistID int, itemName string) *ChecklistItem {
	return &ChecklistItem{
		ID:          id,
		ChecklistID: checklistID,
		ItemName:    itemName,
		Status:      "uncheck",
	}
}
