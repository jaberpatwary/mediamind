package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the index.html file
	content, err := os.ReadFile("frontend/index.html")
	if err != nil {
		panic(err)
	}

	html := string(content)

	// CSS-based MediaMind Logo HTML
	mediaMindLogoHTML := `<div style="display: flex; align-items: center; gap: 12px;">
						<div style="position: relative; width: 45px; height: 45px; background: linear-gradient(135deg, #714DFF 0%, #9C83FF 50%, #E151FF 100%); border-radius: 12px; display: flex; align-items: center; justify-content: center; box-shadow: 0 0 20px rgba(113, 77, 255, 0.4); animation: logoGlow 3s ease-in-out infinite;">
							<span style="font-family: 'Outfit', sans-serif; font-size: 28px; font-weight: 700; color: white; text-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);">M</span>
						</div>
						<div style="font-family: 'Outfit', sans-serif; font-size: 24px; font-weight: 700; background: linear-gradient(80.72deg, #714DFF 0%, #9C83FF 31.28%, #E151FF 95.64%); -webkit-background-clip: text; -webkit-text-fill-color: transparent; background-clip: text; letter-spacing: -0.5px;">MediaMind</div>
					</div>`

	// Small icon version for sticky header
	mediaMindIconHTML := `<div style="position: relative; width: 45px; height: 45px; background: linear-gradient(135deg, #714DFF 0%, #9C83FF 50%, #E151FF 100%); border-radius: 12px; display: flex; align-items: center; justify-content: center; box-shadow: 0 0 20px rgba(113, 77, 255, 0.4);">
							<span style="font-family: 'Outfit', sans-serif; font-size: 28px; font-weight: 700; color: white; text-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);">M</span>
						</div>`

	// Replace AIT main logo
	html = strings.Replace(html, 
		`<a href="https://ait.inc">
			<img width="800" height="339" src="https://ait.inc/wp-content/uploads/2025/01/AIT-new-logo-1024x434.png" class="attachment-large size-large wp-image-2870" alt="" srcset="https://ait.inc/wp-content/uploads/2025/01/AIT-new-logo-1024x434.png 1024w, https://ait.inc/wp-content/uploads/2025/01/AIT-new-logo-300x127.png 300w, https://ait.inc/wp-content/uploads/2025/01/AIT-new-logo-768x325.png 768w, https://ait.inc/wp-content/uploads/2025/01/AIT-new-logo-1536x651.png 1536w, https://ait.inc/wp-content/uploads/2025/01/AIT-new-logo-2048x868.png 2048w" sizes="(max-width: 800px) 100vw, 800px" />				</a>`,
		`<a href="/">` + mediaMindLogoHTML + `</a>`, -1)

	// Replace AIT sticky logo
	html = strings.Replace(html,
		`<a href="https://ait.inc">
			<img width="200" height="108" src="https://ait.inc/wp-content/uploads/2025/04/AIT-symbol.png" class="attachment-large size-large wp-image-2878" alt="" />				</a>`,
		`<a href="/">` + mediaMindIconHTML + `</a>`, -1)

	// Replace title and meta tags
	html = strings.Replace(html, `<title>Portfolio | AIT</title>`, `<title>Portfolio | MediaMind</title>`, -1)
	html = strings.Replace(html, `content="Portfolio | AIT"`, `content="Portfolio | MediaMind"`, -1)
	html = strings.Replace(html, `content="AIT"`, `content="MediaMind"`, -1)

	// Add logo animation CSS if not exists
	if !strings.Contains(html, "@keyframes logoGlow") {
		logoCSS := `
<style>
@keyframes logoGlow {
	0%, 100% { box-shadow: 0 0 20px rgba(113, 77, 255, 0.4); }
	50% { box-shadow: 0 0 30px rgba(113, 77, 255, 0.8); }
}
</style>`
		html = strings.Replace(html, `</head>`, logoCSS+`</head>`, 1)
	}

	// Write updated content
	err = os.WriteFile("frontend/index.html", []byte(html), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Successfully updated branding to MediaMind!")
	fmt.Println("✅ Logo replaced with CSS-based animated MediaMind logo")
	fmt.Println("✅ Title and meta tags updated")
}
