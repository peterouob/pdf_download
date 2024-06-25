package models

type Pdf struct {
	ID       int64
	URL      string
	FileName string
	FileData []string
	Size     int64
}
