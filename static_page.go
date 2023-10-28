package gophront

import "time"

// StaticPage is a static HTML page with no dynamic content changes
type StaticPage struct {
	Name    string
	Header  []Element
	Content []Element
	Style   Style
	Js      string
	Id      string
}

// NewStaticPage creates a new static page
func NewStaticPage(name string, style Style, header []Element, content []Element, js string) StaticPage {
	return StaticPage{name, header, content, style, js, "static-page" + string(time.Now().UnixNano())}
}

func (e StaticPage) Render() (string, error) {

	element := "<!DOCTYPE html><html><head><meta charset=\"utf-8\" />"

	element += "<title>" + e.Name + "</title>"

	if e.Header != nil {
		for _, value := range e.Header {
			renderedContent, err := value.Render()
			if err != nil {
				return "", err
			}
			element += renderedContent
		}
	}

	element += "</head><body id=\"" + e.Id + "\" " + "style=\"" + e.Style.ToString() + "\">"

	element += "<div id=\"app\">"

	if e.Content != nil {
		for _, value := range e.Content {
			renderedContent, err := value.Render()
			if err != nil {
				return "", err
			}
			element += renderedContent
		}
	}

	element += "</div>"

	if e.Js != "" {
		element += "<script>" + e.Js + " </script>"
	}

	return element + "</body></html>", nil
}
