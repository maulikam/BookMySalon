package salon

import (
	"bookmysalon/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SalonHandler struct {
	service SalonService
}

func NewSalonHandler(s SalonService) *SalonHandler {
	return &SalonHandler{service: s}
}

func (h *SalonHandler) CreateSalon(c *gin.Context) {
	var salon models.Salon

	if err := c.ShouldBindJSON(&salon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	salonID, err := h.service.AddSalon(salon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"salon_id": salonID})
}

func (h *SalonHandler) UpdateSalonDetails(c *gin.Context) {
	var salon models.Salon

	if err := c.ShouldBindJSON(&salon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	if err := h.service.UpdateSalon(salon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte("YourSecretKey"), nil // Replace with your secret key
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	service, _ := NewSalonService() // Add proper error handling here
	handler := NewSalonHandler(service)

	r := gin.Default()
	r.Use(Authenticate())
	r.POST("/salon", handler.CreateSalon)
	r.PUT("/salon/update", handler.UpdateSalonDetails)
	r.Run(":8080")
}
