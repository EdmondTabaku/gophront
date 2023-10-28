package gophront

import "fmt"

// SelfClosingElement is an HTML element with a self-closing tag
type SelfClosingElement struct {
	TagName string
	Props   map[string]string
	Style   Style
}

// NewSelfClosingElement creates a new SelfClosingElement
func NewSelfClosingElement(tagName string, props map[string]string, style Style) SelfClosingElement {
	return SelfClosingElement{tagName, props, style}
}

func (e SelfClosingElement) Render() (string, error) {

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

	return element + " />", nil
}

func Br(props map[string]string, style Style) SelfClosingElement {
	return NewSelfClosingElement("br", props, style)
}

func Col(props map[string]string, style Style) SelfClosingElement {
	return NewSelfClosingElement("col", props, style)
}

func Hr(props map[string]string, style Style) SelfClosingElement {
	return NewSelfClosingElement("hr", props, style)
}

func Img(props map[string]string, style Style) SelfClosingElement {
	return NewSelfClosingElement("img", props, style)
}

func Input(props map[string]string, style Style) SelfClosingElement {
	return NewSelfClosingElement("input", props, style)
}

func Link(props map[string]string, style Style) SelfClosingElement {
	return NewSelfClosingElement("link", props, style)
}

func Meta(props map[string]string, style Style) SelfClosingElement {
	return NewSelfClosingElement("meta", props, style)
}
