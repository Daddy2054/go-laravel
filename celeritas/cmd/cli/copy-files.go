package main

import (
	"embed"
	"errors"
	"os"
)
//go:embed templates/*

		//go:embed templates/data/token.go.txt
		//go:embed templates/migrations/migration.postgres.down.sql
		//go:embed templates/migrations/migration.postgres.up.sql
		//go:embed templates/migrations/auth_tables.postgres.sql
		//go:embed templates/data/user.go.txt
var templateFS embed.FS

func copyFilefromTemplate(templatePath, targetFile string) error {
	if fileExists(targetFile) {
		return errors.New(targetFile + " already exists!")
	}

	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile(data, targetFile)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)

	if err != nil {
		return err
	}
	return nil
}

func fileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}