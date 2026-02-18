package userhandler

import (
	"gym/internal/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service user.Service
}

func NewHandler(service user.Service) *Handler {
	return &Handler{service: service}
}

// Registra todas as rotas de user no router
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	// Rotas públicas — sem middleware
	public := users.Group("/")
	{
		public.GET("/", func(ctx *gin.Context) {
			println("testeteste")
		})
		public.POST("/login", h.Create)
		public.POST("/register", h.Create)
	}

	// Rotas privadas — com auth
	private := users.Group("/")
	//private.Use(AuthMiddleware())
	{
		private.GET("/:id", h.GetByID)
		private.DELETE("/:id", h.Delete)
	}
}

func (h *Handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	u, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) Create(c *gin.Context) {
	var body struct {
		Name  string `json:"name"  binding:"required"`
		Email string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*u, err := h.service.Create(c.Request.Context(), body.Name, body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)*/
}

func (h *Handler) Delete(c *gin.Context) {
	//id := c.Param("id")

	/*if err := h.service.(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}*/

	c.JSON(http.StatusNoContent, nil)
}
