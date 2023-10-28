package gophront

import "time"

// DynamicPage is a dynamic HTML page with dynamic content changes
type DynamicPage struct {
	Name    string
	Header  []Element
	Content []Element
	Style   Style
	Js      string
	Id      string
}

// NewDynamicPage creates a new dynamic page
func NewDynamicPage(name string, style Style, header []Element, content []Element, js string) DynamicPage {
	return DynamicPage{name, header, content, style, js, "dynamic-page" + string(time.Now().UnixNano())}
}

func (e DynamicPage) Render() (string, error) {

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

	// The following script is used to handle dynamic links
	element += `<script>
function handleDynamicLinkClick(route, id) {
	fetchContent(route, id);
}

function fetchContent(route, id) {
	const baseUri = window.location.protocol + '//' + window.location.host;
    fetch(baseUri + route)
    .then(response => response.text())
    .then(htmlContent => {
        updateContent(htmlContent, id);
    });
}

function updateContent(htmlContent, id) {
    const appDiv = document.getElementById(id);
    appDiv.innerHTML = htmlContent;
}
`
	if e.Js != "" {
		element += e.Js
	}

	element += " </script>"

	return element + "</body></html>", nil
}
