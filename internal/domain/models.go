package domain

type FileStats struct {
	Name      string
	Words     int
	Chars     int
}

type AnalysisResult struct {
	Files      []FileStats
	TotalWords int
	TotalChars int
}