package routes

import (
	"etentnode-api/app/controller"
	"etentnode-api/app/repository"
	"etentnode-api/app/service"
	"etentnode-api/config"
	"etentnode-api/security"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func WebRouter(db config.Database) {
	// Repository Asset
	userRepo := repository.NewUserRepository(db)
	statusRepo := repository.NewStatusRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	eventCategoryRepo := repository.NewEventCategoryRepository(db)
	eventCategoryTypeRepo := repository.NewEventCategoryTypeRepository(db)
	eventCategoryFieldRepo := repository.NewEventCategoryFieldRepository(db)
	eventRepo := repository.NewEventRepository(db)
	eventHandlingRepo := repository.NewEventHandlingRepository(db)
	notificationRepo := repository.NewNotificitaionRepository(db)

	// Service Asset
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	statusService := service.NewStatusService(statusRepo)
	roleService := service.NewRoleService(roleRepo)
	eventCategoryService := service.NewEventCategoryService(eventCategoryRepo)
	eventCategoryTypeService := service.NewEventCategoryTypeService(eventCategoryTypeRepo)
	eventCategoryFieldService := service.NewEventCategoryFieldService(eventCategoryFieldRepo)
	eventService := service.NewEventService(eventRepo)
	eventHandlingService := service.NewEventHandlingService(eventHandlingRepo)
	notificationService := service.NewNotificationService(notificationRepo)

	//Controller Asset
	authController := controller.NewAuthController(userService, authService)
	userController := controller.NewUserConstroller(userService)
	statusController := controller.NewStatusController(statusService)
	roleController := controller.NewRoleController(roleService)
	eventCategoryController := controller.NewEventCategoryController(eventCategoryService)
	eventCategoryTypeController := controller.NewEventCategoryTypeController(eventCategoryTypeService)
	eventCategoryFieldController := controller.NewEventCategoryFieldService(eventCategoryFieldService)
	eventController := controller.NewEventController(eventService)
	eventHandlingController := controller.NewEventHandlingController(eventHandlingService)

	notificationController := controller.NewNotificationController(notificationService)

	// Route
	httpRouter := gin.Default()

	// Register routing
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Testing  connection
	httpRouter.GET("/status-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Service ✅ API Up and Running"})
	})

	httpRouter.POST("/api/v1/auth/login", authController.Login)
	httpRouter.POST("/api/v1/auth/register", userController.Store)
	httpRouter.Static("/uploads", "./uploads")
	httpRouter.GET("/ws", notificationController.HandleWebSocket)

	v1 := httpRouter.Group("/api/v1") // Grouping routes
	v1.Use(security.AuthMiddleware())
	v1.GET("/auth/logout", authController.Logout)

	v1.GET("/users", userController.Index)
	v1.POST("/users", userController.Store)
	v1.GET("/users/:id", userController.Show)
	v1.PUT("/users/:id", userController.Update)
	v1.DELETE("/users/:id", userController.Delete)

	v1.GET("/status", statusController.Index)
	v1.POST("/status", statusController.Store)
	v1.GET("/status/:id", statusController.Show)
	v1.PUT("/status/:id", statusController.Update)
	v1.DELETE("/status/:id", statusController.Delete)

	v1.GET("/roles", roleController.Index)
	v1.POST("/roles", roleController.Store)
	v1.GET("/roles/:id", roleController.Show)
	v1.PUT("/roles/:id", roleController.Update)
	v1.DELETE("/roles/:id", roleController.Delete)

	v1.GET("/event-categories", eventCategoryController.Index)
	v1.POST("/event-categories", eventCategoryController.Store)
	v1.GET("/event-categories/:id", eventCategoryController.Show)
	v1.PUT("/event-categories/:id", eventCategoryController.Update)
	v1.DELETE("/event-categories/:id", eventCategoryController.Delete)
	v1.POST("/event-categories/assign", eventCategoryController.AssignRole)

	v1.GET("/event-handling", eventHandlingController.Index)
	v1.POST("/event-handling", eventHandlingController.Store)
	v1.GET("/event-handling/:id", eventHandlingController.Show)
	v1.PUT("/event-handling/:id", eventHandlingController.Update)
	v1.DELETE("/event-handling/:id", eventHandlingController.Delete)

	v1.GET("/event-category-types", eventCategoryTypeController.Index)
	v1.POST("/event-category-types", eventCategoryTypeController.Store)
	v1.GET("/event-category-types/:id", eventCategoryTypeController.Show)
	v1.PUT("/event-category-types/:id", eventCategoryTypeController.Update)
	v1.DELETE("/event-category-types/:id", eventCategoryTypeController.Delete)

	v1.GET("/event-category-fields", eventCategoryFieldController.Index)
	v1.POST("/event-category-fields", eventCategoryFieldController.Store)
	v1.GET("/event-category-fields/:id", eventCategoryFieldController.Show)
	v1.PUT("/event-category-fields/:id", eventCategoryFieldController.Update)
	v1.DELETE("/event-category-fields/:id", eventCategoryFieldController.Delete)

	v1.GET("/events", eventController.Index)
	v1.POST("/events", eventController.Store)
	v1.GET("/events/:id", eventController.Show)
	v1.PUT("/events/:id", eventController.Update)
	v1.DELETE("/events/:id", eventController.Delete)
	v1.POST("/events/assign", eventController.AssignUser)
	v1.GET("/events/assign/:id", eventController.ListByUserAssign)
	v1.GET("/events/count", eventController.GetCountEvent)

	v1.POST("/send-notification", notificationController.SendNotificationHandler)
	v1.GET("/notification/:role", notificationController.Index)
	v1.PUT("/notification/:id/:role", notificationController.UpdateByEvent)

	httpRouter.Run(":" + os.Getenv("APP_PORT")) // Run Routes with PORT
}
