package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"fileasync/internal/service"
)

type CLIHandler struct {
	analyzer  service.AnalyzerService
	cli       service.CLIService
	filesDir  string
}

func NewCLIHandler(analyzer service.AnalyzerService, cli service.CLIService, filesDir string) *CLIHandler {
	return &CLIHandler{
		analyzer: analyzer,
		cli:      cli,
		filesDir: filesDir,
	}
}

func (h *CLIHandler) Start() {
	reader := bufio.NewReader(os.Stdin)
	h.cli.PrintWelcome()

	for {
		fmt.Print("Введите команду > ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		switch cmd {
		case "run":
			h.runAnalysis()
		case "exit":
			fmt.Println("Выход из программы...")
			return
		default:
			fmt.Println("Неправильная команда. Попробуйте 'run' или 'exit'")
		}
	}
}

func (h *CLIHandler) runAnalysis() {
	result, err := h.analyzer.AnalyzeFiles(h.filesDir)
	if err != nil {
		h.cli.PrintError(err)
		return
	}
	h.cli.PrintResults(result)
}