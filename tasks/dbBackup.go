package tasks

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func RunBackup() {
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := "db"

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("backup_%s.sql", timestamp)

	cmd := exec.Command("mysqldump",
		"--ssl-mode=DISABLED",
		"-h", dbHost,
		"-u", dbUser,
		fmt.Sprintf("-p%s", dbPassword),
		dbName,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Erro ao gerar backup: %s\nDetalhes: %v\n", string(output), err)
		return
	}

	err = os.WriteFile(fmt.Sprintf("backups/%s", filename), output, 0644)
	if err != nil {
		fmt.Printf("Erro ao salvar o backup: %v\n", err)
		return
	}

	fmt.Println("Backup automático concluído:", filename)
}
