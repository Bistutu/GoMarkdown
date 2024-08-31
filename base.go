package gomarkdown

import (
	"fmt"
	"strings"
)

// H1 generates a level 1 heading
func (mw *markdownWriter) H1(text string) error {
	return mw.writeLine(fmt.Sprintf("# %s", text))
}

// H2 generates a level 2 heading
func (mw *markdownWriter) H2(text string) error {
	return mw.writeLine(fmt.Sprintf("## %s", text))
}

// H3 generates a level 3 heading
func (mw *markdownWriter) H3(text string) error {
	return mw.writeLine(fmt.Sprintf("### %s", text))
}

// H4 generates a level 4 heading
func (mw *markdownWriter) H4(text string) error {
	return mw.writeLine(fmt.Sprintf("#### %s", text))
}

// H5 generates a level 5 heading
func (mw *markdownWriter) H5(text string) error {
	return mw.writeLine(fmt.Sprintf("##### %s", text))
}

// H6 generates a level 6 heading
func (mw *markdownWriter) H6(text string) error {
	return mw.writeLine(fmt.Sprintf("###### %s", text))
}

// Hr generates a horizontal rule
func (mw *markdownWriter) Hr() error {
	return mw.writeLine("---")
}

// Bold generates bold text
func (mw *markdownWriter) Bold(text string) error {
	return mw.writeLine(fmt.Sprintf("**%s**", text))
}

// Italic generates italic text
func (mw *markdownWriter) Italic(text string) error {
	return mw.writeLine(fmt.Sprintf("*%s*", text))
}

// Strikethrough generates strikethrough text
func (mw *markdownWriter) Strikethrough(text string) error {
	return mw.writeLine(fmt.Sprintf("~~%s~~", text))
}

// InlineCode generates inline code
func (mw *markdownWriter) InlineCode(code string) error {
	return mw.writeLine(fmt.Sprintf("`%s`", code))
}

// Footnote generates a footnote
func (mw *markdownWriter) Footnote(identifier, text string) error {
	mw.writeLine(fmt.Sprintf("[%s]: %s", identifier, text))
	return nil
}

// Highlight generates highlighted text
func (mw *markdownWriter) Highlight(text string) error {
	return mw.writeLine(fmt.Sprintf("==%s==", text))
}

// CustomHTML inserts custom HTML content
func (mw *markdownWriter) CustomHTML(html string) error {
	return mw.writeLine(html)
}

// Code generates a code block
func (mw *markdownWriter) Code(language, code string) {
	mw.writeLine(fmt.Sprintf("```%s\n%s\n```", language, code))
}

// List generates an unordered list
func (mw *markdownWriter) List(items []string) {
	for _, item := range items {
		mw.writeLine(fmt.Sprintf("- %s", item))
	}
}

// Link generates a hyperlink
func (mw *markdownWriter) Link(text, url string) {
	mw.writeLine(fmt.Sprintf("[%s](%s)", text, url))
}

// Quote generates a blockquote
func (mw *markdownWriter) Quote(text string) {
	mw.writeLine(fmt.Sprintf("> %s", text))
}

// Subscript generates subscript text
func (mw *markdownWriter) Subscript(text string) error {
	return mw.writeLine(fmt.Sprintf("<sub>%s</sub>", text))
}

// Superscript generates superscript text
func (mw *markdownWriter) Superscript(text string) error {
	return mw.writeLine(fmt.Sprintf("<sup>%s</sup>", text))
}

// DefinitionList generates a definition list
func (mw *markdownWriter) DefinitionList(items []*DefinitionList) {
	for _, item := range items {
		mw.writeLine(fmt.Sprintf("%s\n: %s", item.Term, item.Definition))
	}
}

// OrderedList generates an ordered list
func (mw *markdownWriter) OrderedList(items []string) {
	for i, item := range items {
		mw.writeLine(fmt.Sprintf("%d. %s", i+1, item))
	}
}

// TableOfContents generates a table of contents
func (mw *markdownWriter) TableOfContents(headers []string) {
	for _, header := range headers {
		mw.writeLine(fmt.Sprintf("- [%s](#%s)", header, strings.ToLower(strings.ReplaceAll(header, " ", "-"))))
	}
}

// writeLine will write a line of text to the file
func (mw *markdownWriter) writeLine(text string) error {
	_, err := mw.file.WriteString(text + "\n\n")
	if err != nil {
		fmt.Errorf("write error: %v\n", err)
		return err
	}
	return nil
}
