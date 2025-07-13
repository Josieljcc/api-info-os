package controller

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/Josieljcc/api-info-os/utils"
	"github.com/gin-gonic/gin"
)

func BackupDBController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)

	if !isAuthorized {
		return
	}

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := "db"

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("backup_%s.sql", timestamp)

	cmd := exec.Command("mysqldump",
		"-h", dbHost,
		"-u", dbUser,
		fmt.Sprintf("-p%s", dbPassword),
		"--ssl-mode=DISABLED",
		dbName,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.String(500, fmt.Sprintf("Erro ao gerar backup:\n%s\nDetalhes: %v", string(output), err))
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Data(200, "application/octet-stream", output)
}
