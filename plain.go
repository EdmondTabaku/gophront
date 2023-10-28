package gophront

// Plain is a plain text element without any HTML tags (can be used for anything)
type Plain struct {
	Content string
}

func (e Plain) Render() (string, error) {
	return e.Content, nil
}
