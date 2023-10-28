package gophront

// Element is an interface for all HTML elements
type Element interface {

	// Render returns a string representation of the element
	Render() (string, error)
}
