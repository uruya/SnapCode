package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Define command line options
	output := flag.String("o", "generated_code.png", "Output file name")
	theme := flag.String("theme", "dark", "Theme: dark or light")
	flag.Parse()

	// Get remaining arguments (non-optional)
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: snapcode [-o output.png] [-theme dark|light] '<code string>'")
		os.Exit(1)
	}

	code := args[0]

	// Browser Launch
	launchURL := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(launchURL).MustConnect()
	defer browser.MustClose()

	// Dynamically generate HTML from code Browser startup
	htmlContent := generateHTML(code, *theme)

	// Convert HTML to Data URL and open
	dataURL := "data:text/html;base64," + base64.StdEncoding.EncodeToString([]byte(htmlContent))
	page := browser.MustPage(dataURL)

	// page load wait
	time.Sleep(1 * time.Second)

	// Only the pre element is scrubbed.
	pre := page.MustElement("pre")
	pre.MustScreenshot(*output)

	fmt.Println("Saved screenshot to:", *output)
}

// Generate an HTML template from the given code string
func generateHTML(code string, theme string) string {
	darkCSS := `
	  body { background: #1e1e1e;color: #61dafb; }
	  pre { background: #282c34; color: #61dafb; }	
	  `
	lightCSS := `
	  body { background: #ffffff; color: #333; }
	  pre { background: #f6f8fa; color: #111; }
	  `
	css := darkCSS
	if theme == "light" {
		css = lightCSS
	}

	return fmt.Sprintf(`
<!DOCTYPE html><html><head><meta charset="utf-8">
<title>SnapCode</title><style>%s
pre{padding:16px;border-radius:8px;font-size:16px;max-width:90%%;overflow:auto;}
body,html{display:flex;justify-content:center;align-items:center;height:100vh;margin:0;}
</style></head><body><pre><code>%s</code></pre></body></html>
`, css, code)
}
