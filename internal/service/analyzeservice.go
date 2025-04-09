package service

import (
	"path/filepath"
	"strings"
	"sync"

	"fileasync/internal/domain"
	"fileasync/internal/repository"
)

type AnalyzerService interface {
	AnalyzeFiles(dir string) (*domain.AnalysisResult, error)
}

type analyzerService struct {
	repo repository.FileRepository
}

func NewAnalyzerService(repo repository.FileRepository) AnalyzerService {
	return &analyzerService{repo: repo}
}

func (s *analyzerService) AnalyzeFiles(dir string) (*domain.AnalysisResult, error) {
	files, err := s.repo.GetAllFiles(dir)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	results := make(chan *domain.FileStats, len(files))
	errChan := make(chan error, 1)

	for _, file := range files {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			content, err := s.repo.ReadFileContent(path)
			if err != nil {
				select {
					case errChan <- err:
					default:
				}
				return
			}

			words := len(strings.Fields(content))
			chars := len(content)

			results <- &domain.FileStats{
				Name:  filepath.Base(path),
				Words: words,
				Chars: chars,
			}
		}(file)
	}

	go func() {
		wg.Wait()
		close(results)
		close(errChan)
	}()

	var stats []domain.FileStats
	var totalWords, totalChars int

	for result := range results {
		stats = append(stats, *result)
		totalWords += result.Words
		totalChars += result.Chars
	}

	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return &domain.AnalysisResult{
		Files:      stats,
		TotalWords: totalWords,
		TotalChars: totalChars,
	}, nil
}