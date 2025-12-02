package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/joho/godotenv"
)

var (
	session_cookie = loadSessionCookie()
	lang_templates = map[string]string{
		"go": "templates/main_go.tmpl",
	}
)

//go:embed templates/*
var templatesFS embed.FS

func main() {
	year := flag.Int("year", 0, "Advent of Code year")
	day := flag.Int("day", 0, "Advent of Code day")
	lang := flag.String("lang", "go", "Language for generated main")
	flag.Parse()

	if *year == 0 || *day == 0 {
		log.Fatalf("both --year and --day must be provided. year=%d day=%d", *year, *day)
	}

	// Directory name: cmd/<year>_<day>
	dirName := fmt.Sprintf("%d_%02d", *year, *day)
	dirPath := filepath.Join("cmd", dirName)

	createDayFiles(dirName, dirPath, *lang)
	downloadAOCInput(*year, *day, session_cookie, dirPath)
}

func loadSessionCookie() string {
	// Load .env file (silently ignore if already loaded)
	if err := godotenv.Load(); err != nil {
		// If .env isn't present, continue — env vars might already be set
		log.Printf("warning: could not load .env file: %v", err)
	}

	cookie := os.Getenv("SESSION_COOKIE")
	if cookie == "" {
		log.Println("warning: SESSION_COOKIE is not set")
	}
	return cookie
}

func createDayFiles(dirName, dirPath, lang string) {
	if err := os.MkdirAll(dirPath, 0o755); err != nil {
		log.Fatalf("failed to create directory %s: %v", dirPath, err)
	}

	// Output file
	outPath := filepath.Join(dirPath, "main.go")

	// Parse template file
	tmpl, err := template.ParseFS(templatesFS, lang_templates[lang])
	if err != nil {
		log.Fatalf("failed to parse template file: %v", err)
	}

	// Template data (expand this later if you want day/year in template)
	data := map[string]any{
		"Dir": dirName,
	}

	// Write main.go
	outFile, err := os.Create(outPath)
	if err != nil {
		log.Fatalf("failed to create %s: %v", outPath, err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, data); err != nil {
		log.Fatalf("failed to render template: %v", err)
	}

	fmt.Printf("Created %s\n", outPath)

}

func downloadAOCInput(year, day int, session, dirPath string) error {
	if session == "" {
		return fmt.Errorf("AOC session token is empty — set the AOC_SESSION environment variable")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// AoC requires a "session" cookie
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	// Optional: Identify yourself (recommended but not required)
	req.Header.Set("User-Agent", "github.com/yourusername/aocgen tool")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf(
			"unexpected status %d\nURL: %s\nBody: %s",
			resp.StatusCode, url, string(body),
		)
	}

	outPath := filepath.Join(dirPath, "input.txt")
	outFile, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create input.txt: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
