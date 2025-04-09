package main

import (
	"os"
	"path/filepath"

	"fileasync/internal/console"
	"fileasync/internal/repository"
	"fileasync/internal/service"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic("Не могу определить рабочую директорию")
	}

	filesDir := filepath.Join(wd, "files")


	fileRepo := repository.NewFileRepository()
	analyzerService := service.NewAnalyzerService(fileRepo)
	cliService := service.NewCLIService()

	
	cliHandler := console.NewCLIHandler(analyzerService, cliService, filesDir)
	cliHandler.Start()
}