package render

import (
	"encoding/base64"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// MakePNG returns the screenshot bytes for given code+theme.
func MakePNG(code, theme string) ([]byte, error) {
	url := launcher.New().Headless(true).MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	html := GenerateHTML(code, theme)
	dataURL := "data:text/html;base64," + base64.StdEncoding.EncodeToString([]byte(html))

	page := browser.MustPage(dataURL)
	time.Sleep(1 * time.Second)

	pre := page.MustElement("pre")
	return pre.Screenshot(proto.PageCaptureScreenshotFormatPng, -1)
}
