package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	files := []string{"frontend/index.html", "frontend/admin.html", "frontend/login.html"}
	
	for _, filename := range files {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf("Skipping %s: %v\n", filename, err)
			continue
		}
		
		content := string(data)
		
		// Branding replacements
		content = strings.ReplaceAll(content, "AIT", "Media Mind")
		content = strings.ReplaceAll(content, "ait.inc", "mediamind.com")
		content = strings.ReplaceAll(content, "Portfolio | Media Mind", "Media Mind | Portfolio")
		
		// Logo replacements in index.html
		if filename == "frontend/index.html" {
			content = strings.ReplaceAll(content, "https://ait.inc/wp-content/uploads/2025/01/AIT-new-logo-1024x434.png", "/frontend/img/logo.svg")
			content = strings.ReplaceAll(content, "https://ait.inc/wp-content/uploads/2025/04/AIT-symbol.png", "/frontend/img/logo.svg")
            // Fix header title
            content = strings.ReplaceAll(content, "<title>Portfolio | Media Mind</title>", "<title>Media Mind | Premium Portfolio</title>")
		}

        // Admin replacements
        if filename == "frontend/admin.html" {
            content = strings.ReplaceAll(content, "Jaber Core", "Media Mind Admin")
        }
		
		err = ioutil.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error writing %s: %v\n", filename, err)
		} else {
			fmt.Printf("Updated branding in %s\n", filename)
		}
	}
}
