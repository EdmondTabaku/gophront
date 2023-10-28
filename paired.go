package gophront

import (
	"fmt"
)

// PairedElement is an HTML element with a start and end tag
type PairedElement struct {
	TagName string
	Props   map[string]string
	Content []Element
	Style   Style
}

// NewPairedElement creates a new paired element
func NewPairedElement(tagName string, props map[string]string, style Style, content []Element) PairedElement {
	return PairedElement{tagName, props, content, style}
}

func (e PairedElement) Render() (string, error) {

	if e.TagName == "" {
		return "", fmt.Errorf("tag name is empty")
	}

	if e.Props == nil {
		e.Props = make(map[string]string)
	}

	if e.Style != (Style{}) {
		e.Props["style"] = e.Props["style"] + e.Style.ToString()
	}

	element := "<" + e.TagName

	for key, value := range e.Props {
		element += " " + key + "=\"" + value + "\""
	}

	element += ">"

	for _, value := range e.Content {
		renderedContent, err := value.Render()
		if err != nil {
			return "", err
		}
		element += renderedContent
	}

	return element + "</" + e.TagName + ">", nil
}

func A(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("a", props, style, content)
}

func B(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("b", props, style, content)
}

func Body(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("body", props, style, content)
}

func Button(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("button", props, style, content)
}

func Canvas(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("canvas", props, style, content)
}

func Code(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("code", props, style, content)
}

func Div(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("div", props, style, content)
}

func Em(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("em", props, style, content)
}

func Form(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("form", props, style, content)
}

func Span(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("span", props, style, content)
}

func H1(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("h1", props, style, content)
}

func H2(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("h2", props, style, content)
}

func H3(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("h3", props, style, content)
}

func H4(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("h4", props, style, content)
}

func H5(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("h5", props, style, content)
}

func H6(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("h6", props, style, content)
}

func Head(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("head", props, style, content)
}

func Header(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("header", props, style, content)
}

func Html(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("html", props, style, content)
}

func Iframe(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("iframe", props, style, content)
}

func Label(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("label", props, style, content)
}

func Li(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("li", props, style, content)
}

func Nav(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("nav", props, style, content)
}

func NoScript(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("noscript", props, style, content)
}

func Ol(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("ol", props, style, content)
}

func Option(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("option", props, style, content)
}

func P(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("p", props, style, content)
}

func Script(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("script", props, style, content)
}

func Section(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("section", props, style, content)
}

func Select(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("select", props, style, content)
}

func Strong(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("strong", props, style, content)
}

func Table(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("table", props, style, content)
}

func Tbody(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("tbody", props, style, content)
}

func Td(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("td", props, style, content)
}

func Textarea(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("textarea", props, style, content)
}

func Tfoot(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("tfoot", props, style, content)
}

func Thead(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("thead", props, style, content)
}

func Title(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("title", props, style, content)
}

func Tr(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("tr", props, style, content)
}

func Ul(props map[string]string, style Style, content []Element) PairedElement {
	return NewPairedElement("ul", props, style, content)
}

// DynamicLink is a link that changes the page body dynamically without reloading it
func DynamicLink(route string, props map[string]string, style Style, content []Element) PairedElement {
	if props == nil {
		props = make(map[string]string)
	}

	props["onclick"] = fmt.Sprintf("handleDynamicLinkClick('%s', 'app')", route)

	return NewPairedElement("a", props, style, content)
}

// DynamicContentLink is a link that changes the specified content dynamically without reloading the page
func DynamicContentLink(route string, contentId string, props map[string]string, style Style, content []Element) PairedElement {
	if props == nil {
		props = make(map[string]string)
	}

	props["onclick"] = fmt.Sprintf("handleDynamicLinkClick('%s', '%s')", route, contentId)

	return NewPairedElement("a", props, style, content)
}
