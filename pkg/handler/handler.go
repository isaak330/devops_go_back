package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		characters := api.Group("/characters")
		{
			characters.POST("/", h.CreateCharacter)
			characters.GET("/", h.GetCharacters)
			characters.GET("/:id", h.GetCharacterById)
			characters.PUT("/:id", h.UpdateCharacter)
			characters.DELETE("/:id", h.DeleteCharacter)
		}
	}
	return router
}
