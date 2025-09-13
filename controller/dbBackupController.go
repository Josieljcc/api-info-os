package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/Josieljcc/api-info-os/utils"
	"github.com/gin-gonic/gin"
)

// BackupDBController godoc
// @Summary Backup a database
// @Description Backup a database
// @Tags Database
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Router /backup [get]
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
		"--no-tablespaces",
		"--ssl-mode=DISABLED",
		"-h", dbHost,
		"-u", dbUser,
		fmt.Sprintf("-p%s", dbPassword),
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

// RestoreDBController godoc
// @Summary Restore a database
// @Description Restore a database
// @Tags Database
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param file formData file true "File"
// @Router /restore [post]
func RestoreDBController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Arquivo não encontrado na requisição")
		return
	}

	tempFile, err := file.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao abrir o arquivo")
		return
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(tempFile)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao ler o conteúdo do arquivo")
		return
	}

	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := "db"

	cmd := exec.Command("mysql",
		"--ssl-mode=DISABLED",
		"-h", dbHost,
		"-u", dbUser,
		fmt.Sprintf("-p%s", dbPassword),
		dbName,
	)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao criar stdin pipe")
		return
	}

	go func() {
		defer stdin.Close()
		stdin.Write(fileBytes)
	}()

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.String(500, fmt.Sprintf("Erro ao restaurar backup:\n%s\nDetalhes: %v", string(output), err))
		return
	}

	c.String(200, "Backup restaurado com sucesso!")
}
