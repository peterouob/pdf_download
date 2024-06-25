package service

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/redis/go-redis/v9"
	"os"
)

type DownloadService struct {
	redis   *redis.Client
	context context.Context
}

func NewDownloadService(redis *redis.Client) *DownloadService {
	return &DownloadService{
		redis:   redis,
		context: context.Background(),
	}
}

func (s *DownloadService) DownloadAndSavePDF(url, outPath, name string) error {
	if err := downLoadFile(url, outPath); err != nil {
		return err
	}
	pdfData, err := os.ReadFile(outPath)
	if err != nil {
		return err
	}
	fileInfo, err := os.Stat(outPath)
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()
	fileMeta := map[string]interface{}{
		"name":     name,
		"url":      url,
		"fileData": pdfData,
		"fileSize": fileSize,
	}
	if err := s.redis.HMSet(s.context, url, fileMeta).Err(); err != nil {
		return err
	}
	return nil
}

func downLoadFile(url, outPath string) error {
	ctx, channel := chromedp.NewContext(context.Background())
	defer channel()

	var buf []byte
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	}); err != nil {
		return err
	}

	if err := os.WriteFile(outPath, buf, 0777); err != nil {
		return err
	}
	return nil
}
