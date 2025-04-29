package render

import "fmt"

func GenerateHTML(code, theme string) string {
	darkCSS := `
      body{background:#1e1e1e;color:#61dafb}
      pre {background:#282c34;color:#61dafb}`
	lightCSS := `
      body{background:#ffffff;color:#333}
      pre {background:#f6f8fa;color:#111}`

	css := darkCSS
	if theme == "light" {
		css = lightCSS
	}

	return fmt.Sprintf(`
<!DOCTYPE html><html><head><meta charset="utf-8"><title>SnapCode</title>
<style>%s
pre{padding:16px;border-radius:8px;font-size:16px;max-width:90%%;overflow:auto;}
body,html{display:flex;justify-content:center;align-items:center;height:100vh;margin:0;}
</style></head><body><pre><code>%s</code></pre></body></html>`, css, code)
}
