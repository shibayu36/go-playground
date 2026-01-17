package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/slack-go/slack"
)

func TestSlackFileSearch(t *testing.T) {
	token := os.Getenv("SLACK_USER_TOKEN")
	query := os.Getenv("SLACK_FILE_SEARCH_QUERY")
	if query == "" {
		log.Fatal("SLACK_FILE_SEARCH_QUERY is required")
	}
	query = query + " type:canvases"

	api := slack.New(token)

	// ファイルを検索
	searchParams := slack.SearchParameters{
		Sort:          "timestamp",
		SortDirection: "desc",
		Count:         100,
	}
	result, err := api.SearchFiles(query, searchParams)
	if err != nil {
		fmt.Printf("search error: %+v\n", err)
		log.Fatal(err)
	}

	fmt.Printf("Found %d files for query: %s\n", result.Total, query)

	// ダウンロードディレクトリを作成（タイムスタンプ付き）
	timestamp := time.Now().Format("20060102_150405")
	downloadDir := fmt.Sprintf("slack_downloads_%s", timestamp)
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		log.Fatalf("failed to create download directory: %+v", err)
	}

	// 検索結果のファイルをすべてダウンロード
	for _, f := range result.Matches {
		fmt.Printf("Downloading: %s (ID: %s)\n", f.Name, f.ID)

		downloadURL := f.URLPrivateDownload
		if downloadURL == "" {
			downloadURL = f.URLPrivate
		}

		filePath := filepath.Join(downloadDir, f.Name+".html")
		out, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("create file error for %s: %+v\n", f.Name, err)
			continue
		}

		if err := api.GetFile(downloadURL, out); err != nil {
			fmt.Printf("download error for %s: %+v\n", f.Name, err)
			out.Close()
			continue
		}
		out.Close()
		fmt.Printf("Downloaded: %s\n", filePath)
	}
}
