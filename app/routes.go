package app

import (
	"github.com/crowdeco/demo-todo-service/controllers"
)

func mapUrls() {
	v1 := router.Group("api/v1/")
	{
		kycGroup := v1.Group("users")
		{
			kycGroup.GET("/:id", user.GetUserProfile)
			kycGroup.POST("/", user.CreateUser)
		}
	}
}
