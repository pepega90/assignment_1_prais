// package router mengatur rute untuk aplikasi
package router

import (
	"assignment_1/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter menginisialisasi dan mengatur rute untuk aplikasi
func SetupRouter(r *gin.Engine,
	userHandler handler.IUserHandler,
	submissionsHandler handler.ISubmissionHandler,
) {
	usersPublicEndpoint := r.Group("/users")
	usersPublicEndpoint.GET("/:id", userHandler.GetUser)
	usersPublicEndpoint.GET("", userHandler.GetAllUsers)
	usersPublicEndpoint.GET("/", userHandler.GetAllUsers)
	usersPublicEndpoint.POST("", userHandler.CreateUser)
	usersPublicEndpoint.POST("/", userHandler.CreateUser)
	usersPublicEndpoint.PUT("/:id", userHandler.UpdateUser)
	usersPublicEndpoint.DELETE("/:id", userHandler.DeleteUser)

	submissionsPublicEndpoint := r.Group("/submissions")
	submissionsPublicEndpoint.GET("/:id", submissionsHandler.GetSubmission)
	submissionsPublicEndpoint.GET("", submissionsHandler.GetAllSubmissions)
	submissionsPublicEndpoint.GET("/", submissionsHandler.GetAllSubmissions)
	submissionsPublicEndpoint.POST("", submissionsHandler.CreateSubmission)
	submissionsPublicEndpoint.POST("/", submissionsHandler.CreateSubmission)
	submissionsPublicEndpoint.DELETE("/:id", submissionsHandler.DeleteSubmission)
}
