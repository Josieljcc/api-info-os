package controller

import (
	"github.com/Josieljcc/api-info-os/service"
	"github.com/Josieljcc/api-info-os/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Login
// @Description Login
// @Tags Login
// @Accept  json
// @Produce  json
// @Param body body schemas.ClientLogin true "Client Login"
// @Success 200 {object} schemas.ClientLoginResponse
// @Router /login [post]
func LoginController(c *gin.Context) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if creds.Email == "" || creds.Password == "" {
		c.JSON(400, gin.H{
			"error": "Email and password are required",
		})
		return
	}

	tecnitian, tecerr := service.GetTechnicianByEmail(creds.Email)
	client, clienterr := service.GetClientByEmail(creds.Email)
	if tecerr != nil && clienterr != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if !utils.CheckPassword(creds.Password, tecnitian.Password) && !utils.CheckPassword(creds.Password, client.Password) {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	var idToCreateToken uint
	var role string
	if tecerr == nil {
		idToCreateToken = tecnitian.ID
		role = "technician"
	}

	if clienterr == nil {
		idToCreateToken = client.ID
		role = "client"
	}

	token, err := utils.GenerateToken(idToCreateToken, role)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "error generating token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"role":  role,
		"email": creds.Email
	})

}
