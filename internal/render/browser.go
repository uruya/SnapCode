package render

import (
	"sync"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

var (
	browser *rod.Browser
	once    sync.Once
)

func getBrowser() *rod.Browser {
	once.Do(func() {
		url := launcher.New().Headless(true).MustLaunch()
		browser = rod.New().ControlURL(url).MustConnect()
	})
	return browser
}
