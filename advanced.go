package gomarkdown

import (
	"fmt"
	"strings"
)

// Table generates a table
func (mw *markdownWriter) Table(headers []string, rows [][]string) {
	mw.writeLine(strings.Join(headers, " | "))
	mw.writeLine(strings.Repeat("--- | ", len(headers)))
	for _, row := range rows {
		mw.writeLine(strings.Join(row, " | "))
	}
}

// TaskList generates a task list
func (mw *markdownWriter) TaskList(items []*TaskModel) {
	for _, item := range items {
		status := " "
		if item.Done {
			status = "x"
		}
		mw.writeLine(fmt.Sprintf("- [%s] %s", status, item.Text))
	}
}

// AlertBox generates a GitHub-style alert box
func (mw *markdownWriter) AlertBox(alertType, text string) error {
	switch alertType {
	case "note":
		return mw.writeLine(fmt.Sprintf("> **Note**: %s", text))
	case "warning":
		return mw.writeLine(fmt.Sprintf("> **Warning**: %s", text))
	case "danger":
		return mw.writeLine(fmt.Sprintf("> **Danger**: %s", text))
	default:
		return mw.writeLine(fmt.Sprintf("> %s", text))
	}
}

// Image generates an image
func (mw *markdownWriter) Image(altText, url string) {
	mw.writeLine(fmt.Sprintf("![%s](%s)", altText, url))
}

// ImageWithSize generates an image with specified width and height
func (mw *markdownWriter) ImageWithSize(altText, url string, width, height int) error {
	// If both width and height are provided, use both; otherwise, use only the specified dimension
	if width > 0 && height > 0 {
		return mw.writeLine(fmt.Sprintf("<img src=\"%s\" alt=\"%s\" width=\"%d\" height=\"%d\" />", url, altText, width, height))
	} else if width > 0 {
		return mw.writeLine(fmt.Sprintf("<img src=\"%s\" alt=\"%s\" width=\"%d\" />", url, altText, width))
	} else if height > 0 {
		return mw.writeLine(fmt.Sprintf("<img src=\"%s\" alt=\"%s\" height=\"%d\" />", url, altText, height))
	}
	// Default to the original Image function if no size is specified
	return mw.writeLine(fmt.Sprintf("![%s](%s)", altText, url))
}

// ImageWithScale generates an image with a specified zoom percentage
func (mw *markdownWriter) ImageWithScale(altText, url string, scale int) error {
	// Ensure scale is a positive integer
	if scale <= 0 {
		return fmt.Errorf("scale must be a positive integer")
	}

	// Generate the image with a specified zoom percentage
	return mw.writeLine(fmt.Sprintf("<img src=\"%s\" alt=\"%s\" style=\"zoom:%d%%;\" />", url, altText, scale))
}

// SequenceDiagram generates a sequence diagram using Mermaid syntax
func (mw *markdownWriter) SequenceDiagram(content string) error {
	return mw.writeLine(fmt.Sprintf("```mermaid\nsequenceDiagram\n%s\n```", content))
}

// Flowchart generates a flowchart using Mermaid syntax
func (mw *markdownWriter) Flowchart(content string) error {
	return mw.writeLine(fmt.Sprintf("```mermaid\nflowchart TD\n%s\n```", content))
}

// MermaidDiagram generates a custom Mermaid diagram
func (mw *markdownWriter) MermaidDiagram(content string) error {
	return mw.writeLine(fmt.Sprintf("```mermaid\n%s\n```", content))
}

// CodeBlockMath generates a code block containing a mathematical expression using LaTeX syntax
func (mw *markdownWriter) CodeBlockMath(expression string) error {
	return mw.writeLine(fmt.Sprintf("```math\n%s\n```", expression))
}

// InlineMath generates inline mathematical expressions using LaTeX syntax
func (mw *markdownWriter) InlineMath(expression string) error {
	return mw.writeLine(fmt.Sprintf("$%s$", expression))
}
