# GoMarkdown

[中文](README_CN.md)

GoMarkdown is a Go library for generating Markdown documents programmatically. It offers comprehensive support for various Markdown elements, from basic formatting to advanced features like Mermaid diagrams, mathematical expressions, and GitHub-style alert boxes.

## Features

- **Basic Markdown Elements:**
  - Headings (H1-H6)
  - Bold, Italic, Strikethrough text
  - Inline code and code blocks
  - Horizontal rules
  - Links and Images (with customizable size and scale)
  - Blockquotes and Footnotes
  - Unordered and ordered lists, task lists
  - Definition lists and Table of contents

- **Advanced Features:**
  - Mermaid diagrams (sequence diagrams, flowcharts, and custom diagrams)
  - Mathematical expressions using LaTeX syntax (inline and block)
  - GitHub-style alert boxes (note, warning, danger)
  - Custom HTML content embedding

## Installation

To install GoMarkdown, use the following `go get` command:

```bash
go get github.com/Bistutu/GoMarkdown
```

## Usage

Here's an example of how to use GoMarkdown to create a Markdown document:

```go
package main

import (
	"fmt"
	"os"
	"github.com/Bistutu/GoMarkdown"
)

func main() {
	// Create a MarkdownWriter instance
	writer, err := gomarkdown.NewMarkdownWriter("example.md", true)
	if err != nil {
		fmt.Printf("Failed to create MarkdownWriter: %v\n", err)
		return
	}
	defer writer.Close()

	// Generate Markdown content
	writer.H1("GoMarkdown Example")
	writer.Bold("This is bold text")
	writer.Italic("This is italic text")
	writer.Code("go", `fmt.Println("Hello, GoMarkdown!")`)
	writer.ImageWithScale("Sample Image", "https://example.com/image.png", 50)

	// Additional Markdown generation examples...
}
```

## Documentation

For detailed usage and all available functions, please refer to the [GoDoc documentation](https://pkg.go.dev/github.com/Bistutu/GoMarkdown).

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Make your changes and commit them (`git commit -m 'Add YourFeature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a new Pull Request.

Please adhere to the [Go style guidelines](https://golang.org/doc/effective_go) when contributing.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Go](https://golang.org/) - The Go programming language
- [Markdown Guide](https://www.markdownguide.org/) - A comprehensive guide to Markdown syntax
