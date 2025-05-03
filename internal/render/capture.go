package render

import (
	"encoding/base64"
	"time"

	"github.com/go-rod/rod/lib/proto"
)

// MakePNG returns the screenshot bytes for given code+theme.
func MakePNG(code, theme string) ([]byte, error) {
	br := getBrowser()
	html := GenerateHTML(code, theme)
	dataURL := "data:text/html;base64," + base64.StdEncoding.EncodeToString([]byte(html))

	page, err := br.Page(proto.TargetCreateTarget{})
	if err != nil {
		return nil, err
	}
	defer page.Close()
	if err := page.Navigate(dataURL); err != nil {
		return nil, err
	}
	page.WaitLoad()

	time.Sleep(1 * time.Second)

	pre := page.MustElement("pre")
	return pre.Screenshot(proto.PageCaptureScreenshotFormatPng, -1)
}
