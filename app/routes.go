package app

import (
	"github.com/crowdeco/demo-todo-service/controllers"
)

func mapUrls() {
	v1 := router.Group("api/v1/")
	{
		todoGroup := v1.Group("todo")
		{
			todoGroup.GET("/:id", controllers.GetUserProfile)
		}
	}
}
