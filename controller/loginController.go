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
// @Param   email    query  string  true "Email"
// @Param   password    query  string  true "Password"
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
	if tecerr == nil {
		idToCreateToken = tecnitian.ID
	}

	if clienterr == nil {
		idToCreateToken = client.ID
	}

	token, err := utils.GenerateToken(idToCreateToken)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "error generating token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}
