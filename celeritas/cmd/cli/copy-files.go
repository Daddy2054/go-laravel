package main

import (
	"embed"
	"os"
)

//go:embed templates/migrations/migration.postgres.down.sql
//go:embed templates/migrations/migration.postgres.up.sql
//go:embed templates/migrations/auth_tables.postgres.sql
var templateFS embed.FS

func copyFilefromTemplate(templatePath, targetFile string) error {
	// TODO: check to ensure file does not already exist

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
