// router/route.go

package router

import (
	"todolist/controller"
	"todolist/handlers"
	"todolist/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	checklistController := controller.NewChecklistController()
	checklistItemController := controller.NewChecklistItemController()

	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)

	private := r.Group("/api")
	private.Use(utils.AuthMiddleware())
	{
		private.POST("/checklist", checklistController.CreateChecklist)
		private.GET("/checklist", checklistController.GetAllChecklists)
		private.DELETE("/checklist/:id", checklistController.DeleteChecklist)
		private.POST("/checklist/:checklistId/item", checklistItemController.CreateChecklistItem)
		private.GET("/checklist/:checklistId/item", checklistItemController.GetAllChecklistItemsByChecklistID)
		private.GET("/checklist/:checklistId/item/:checklistItemId", checklistItemController.GetChecklistItemByChecklistIDAndItemID)
		private.PUT("/checklist/:checklistId/item/:checklistItemId", checklistItemController.UpdateChecklistItemStatus)
		private.PUT("/checklist/:checklistId/item/rename/:checklistItemId", checklistItemController.RenameChecklistItem)
	}

	return r
}
