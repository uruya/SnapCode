package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Get code from command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go '<code string>'")
		os.Exit(1)
	}

	code := os.Args[1]

	// Browser Launch
	launchURL := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(launchURL).MustConnect()
	defer browser.MustClose()

	// Dynamically generate HTML from code Browser startup
	htmlContent := generateHTML(code)

	// Convert HTML to Data URL and open
	dataURL := "data:text/html;base64," + base64.StdEncoding.EncodeToString([]byte(htmlContent))
	page := browser.MustPage(dataURL)

	// page load wait
	time.Sleep(1 * time.Second)

	// Only the pre element is scrubbed.
	pre := page.MustElement("pre")
	outputFile := "generated_code.png"
	pre.MustScreenshot(outputFile)

	fmt.Println("Screenshot saved:", outputFile)
}

// Generate an HTML template from the given code string
func generateHTML(code string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head><meta charset="UTF-8"><title>SnapCode</title>
  <style>
    body { background: #1e1e1e; display:flex;justify-content:center;align-items:center;height:100vh;margin:0; }
    pre  { background:#282c34;color:#61dafb;padding:16px;border-radius:8px;font-size:16px;overflow:auto; }
  </style>
</head>
<body>
  <pre><code>%s</code></pre>
</body>
</html>
`, code)
}
