package service

import (
	"fmt"

	"fileasync/internal/domain"
)

type CLIService interface {
	PrintWelcome()
	PrintResults(result *domain.AnalysisResult)
	PrintError(err error)
}

type cliService struct{}

func NewCLIService() CLIService {
	return &cliService{}
}

func (s *cliService) PrintWelcome() {
	fmt.Println("Анализатор файлов")
	fmt.Println("---------------------")
	fmt.Println("Доступные команды:")
	fmt.Println("• run - анализ файлов в папке files")
	fmt.Println("• exit - выход из программы")
	fmt.Println()
}

func (s *cliService) PrintResults(result *domain.AnalysisResult) {
	fmt.Println("\nРезультаты анализа:")
	for i, file := range result.Files {
		fmt.Printf("%d. %s - %d слов, %d символов\n", i+1, file.Name, file.Words, file.Chars)
	}
	fmt.Printf("\nИтого: %d слов, %d символов\n\n", result.TotalWords, result.TotalChars)
}

func (s *cliService) PrintError(err error) {
	fmt.Printf("\nОшибка: %v\n\n", err)
}