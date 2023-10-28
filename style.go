package gophront

// Style contains CSS style properties
type Style struct {
	BackgroundColor string
	BackgroundImage string
	Border          string
	BorderRadius    string
	Bottom          string
	BoxShadow       string
	Color           string
	Cursor          string
	Display         string
	Float           string
	FontFamily      string
	FontSize        string
	FontWeight      string
	Height          string
	Left            string
	LineHeight      string
	Margin          string
	Opacity         string
	Padding         string
	Position        string
	Right           string
	TextAlign       string
	TextDecoration  string
	TextTransform   string
	Top             string
	Transition      string
	Width           string
	ZIndex          string
}

// ToString converts the Style to a string
func (s Style) ToString() string {
	style := ""
	if s.BackgroundColor != "" {
		style += "background-color:" + s.BackgroundColor + ";"
	}
	if s.BackgroundImage != "" {
		style += "background-image:" + s.BackgroundImage + ";"
	}
	if s.Border != "" {
		style += "border:" + s.Border + ";"
	}
	if s.BorderRadius != "" {
		style += "border-radius:" + s.BorderRadius + ";"
	}
	if s.Bottom != "" {
		style += "bottom:" + s.Bottom + ";"
	}
	if s.BoxShadow != "" {
		style += "box-shadow:" + s.BoxShadow + ";"
	}
	if s.Color != "" {
		style += "color:" + s.Color + ";"
	}
	if s.Cursor != "" {
		style += "cursor:" + s.Cursor + ";"
	}
	if s.Display != "" {
		style += "display:" + s.Display + ";"
	}
	if s.Float != "" {
		style += "float:" + s.Float + ";"
	}
	if s.FontFamily != "" {
		style += "font-family:" + s.FontFamily + ";"
	}
	if s.FontSize != "" {
		style += "font-size:" + s.FontSize + ";"
	}
	if s.FontWeight != "" {
		style += "font-weight:" + s.FontWeight + ";"
	}
	if s.Height != "" {
		style += "height:" + s.Height + ";"
	}
	if s.Left != "" {
		style += "left:" + s.Left + ";"
	}
	if s.LineHeight != "" {
		style += "line-height:" + s.LineHeight + ";"
	}
	if s.Margin != "" {
		style += "margin:" + s.Margin + ";"
	}
	if s.Opacity != "" {
		style += "opacity:" + s.Opacity + ";"
	}
	if s.Padding != "" {
		style += "padding:" + s.Padding + ";"
	}
	if s.Position != "" {
		style += "position:" + s.Position + ";"
	}
	if s.Right != "" {
		style += "right:" + s.Right + ";"
	}
	if s.TextAlign != "" {
		style += "text-align:" + s.TextAlign + ";"
	}
	if s.TextDecoration != "" {
		style += "text-decoration:" + s.TextDecoration + ";"
	}
	if s.TextTransform != "" {
		style += "text-transform:" + s.TextTransform + ";"
	}
	if s.Top != "" {
		style += "top:" + s.Top + ";"
	}
	if s.Transition != "" {
		style += "transition:" + s.Transition + ";"
	}
	if s.Width != "" {
		style += "width:" + s.Width + ";"
	}
	if s.ZIndex != "" {
		style += "z-index:" + s.ZIndex + ";"
	}
	return style
}
