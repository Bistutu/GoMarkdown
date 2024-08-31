package gomarkdown

import (
	"testing"
)

// TestMarkdownWriter tests all the functions of the gomarkdown package
func TestMarkdownWriter(t *testing.T) {
	// Create a new markdownWriter instance

	mw, err := NewMarkdownWriter("example", false)
	if err != nil {
		t.Fatal(err)
	}

	// Testing different Markdown elements
	mw.H1("Heading Level 1")
	mw.H2("Heading Level 2")
	mw.H3("Heading Level 3")
	mw.H4("Heading Level 4")
	mw.H5("Heading Level 5")
	mw.H6("Heading Level 6")

	mw.Hr()

	mw.Bold("Bold Text Example")
	mw.Italic("Italic Text Example")
	mw.Strikethrough("Strikethrough Text Example")
	mw.InlineCode("Inline Code Example")

	mw.Footnote("1", "This is a footnote example.")
	mw.Highlight("Highlighted Text Example")
	mw.CustomHTML("<p>This is custom HTML content.</p>")

	mw.Code("go", `fmt.Println("Hello, Go!")`)

	mw.List([]string{"Item 1", "Item 2", "Item 3"})
	mw.OrderedList([]string{"First", "Second", "Third"})
	mw.Link("Go to Google", "https://www.google.com")
	mw.Image("Go Logo", "https://golang.org/doc/gopher/frontpage.png")
	mw.ImageWithSize("Go Logo", "https://golang.org/doc/gopher/frontpage.png", 100, 100)
	mw.ImageWithScale("Go Logo", "https://golang.org/doc/gopher/frontpage.png", 30)
	mw.Quote("This is a blockquote example.")
	mw.Subscript("Subscript Example")
	mw.Superscript("Superscript Example")

	definitions := []*DefinitionList{
		{Term: "Term 1", Definition: "Definition for Term 1"},
		{Term: "Term 2", Definition: "Definition for Term 2"},
	}
	mw.DefinitionList(definitions)

	mw.TableOfContents([]string{"Heading Level 1", "Heading Level 2", "Heading Level 3"})

	mw.Table([]string{"Header 1", "Header 2"}, [][]string{
		{"Row 1, Col 1", "Row 1, Col 2"},
		{"Row 2, Col 1", "Row 2, Col 2"},
	})

	taskItems := []*TaskModel{
		{Text: "Task 1", Done: true},
		{Text: "Task 2", Done: false},
	}
	mw.TaskList(taskItems)

	mw.AlertBox("note", "This is a note alert.")
	mw.AlertBox("warning", "This is a warning alert.")
	mw.AlertBox("danger", "This is a danger alert.")

	mw.SequenceDiagram("Alice -> Bob: Hello Bob!")
	mw.Flowchart("A-->B")
	mw.MermaidDiagram("graph TD; A-->B;")

	mw.CodeBlockMath("E = mc^2")
	mw.InlineMath("a^2 + b^2 = c^2")
}
