
# 🎭 Gophront: Frontend Development with Go

GoPhront: Make Web Pages that Shine with Go
Welcome to GoPhront! Think of it as a magic box where you can make web pages that stay the same or change when you want them to. Want a page that never changes? Easy. Want a part of your page to change without refreshing? You got it. With GoPhront, making cool web pages is a breeze. Dive in and have fun! 


## ⁉️ How does it work?

### Your Control, Your Code

With GoPhront, you're the maestro of your web show. 
Every little bit, every element on your webpage, is under your command. 

It's like building with digital lego bricks, where each brick has a purpose, and you decide where and how it fits.

### Live from Your Go Server
Instead of having your frontend and backend talk across a vast digital divide, GoPhront serves everything straight from your Go server. 

It's like having a live concert right in your living room!

### Two Kinds of Web Pages: Static & Dynamic
#### Static Pages:
**Lightweight:** Since they have minimal JavaScript, these pages are like paper planes – light and fast.

**Refresh to Change:** They're set in their ways. If you want to see something new, give it a quick refresh.


#### Dynamic Pages:

**Feature-packed:** A bit heavier due to the extra JavaScript, but they come with superpowers!

**Live Changes:** Click a special link, and watch parts of the page change without the whole page reloading. It's like magic!

### Be the Boss of Your Page's State
Worried about managing the state of your web app? Don't be! With GoPhront, you get the freedom to choose. Use sessions or any other method you like. You're in the driver's seat.

### Perfect for All-In-One Applications
If you love having your backend and frontend working closely, like best friends, GoPhront is your match made in heaven. It bridges the two seamlessly, making them work together in perfect harmony.
## 🎥 Demo

This demo demonstrates the capabilities of the ```gophront``` library in crafting dynamic web pages, integrated with the ```gin-gonic/gin``` web framework for routing and session management.

### Overview
The demonstration entails a simple webpage that displays a counter. Users can press an "INCREMENT" button, and the new incremented counter will be fetchet from the go server and updated without refreshing the page

### Prerequisites
Go installed on your system.

An understanding of the ```gin-gonic/gin``` web framework.

### Dependencies
``` go
import (
	"fmt"
	"github.com/EdmondTabaku/gophront"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

```

Install these using:

``` bash
go get github.com/EdmondTabaku/gophront
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/sessions
go get github.com/gin-contrib/sessions/cookie

```

### Web Server and Session Configuration

``` go
func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret-key"))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 60 * 60 * 24 * 7, Path: "/"})
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", Home)
	r.GET("/increment", Increment)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

```

### Home Page ('/')
``` go
func Home(c *gin.Context) {
	session := sessions.Default(c)
	// initialize session
	if session.Get("count") == nil {
		session.Set("count", 0)
		err := session.Save()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
	// get current count
	currentCount := session.Get("count")

	homePage := gophront.NewDynamicPage(
		"Hello",
		gophront.Style{
			BackgroundColor: "#2D4356",
			FontFamily:      "Arial",
			FontSize:        "20px",
			FontWeight:      "bold",
			Color:           "#ffffff",
		},
		nil,
		[]gophront.Element{
			gophront.Div(nil, gophront.Style{
				Width:  "150px",
				Margin: "0 auto",
				Height: "400px",
			}, []gophront.Element{

				gophront.H1(
					map[string]string{"id": "count"},
					gophront.Style{
						FontSize:  "50px",
						TextAlign: "center",
					}, []gophront.Element{
						gophront.Plain{Content: fmt.Sprintf("%v", currentCount)},
					}),
				gophront.Div(map[string]string{"id": "container"}, gophront.Style{}, []gophront.Element{
					gophront.DynamicContentLink(
						"/increment",
						"count",
						nil,
						gophront.Style{
							BackgroundColor: "#E57C23",
							Padding:         "10px 20px",
							BorderRadius:    "5px",
							Cursor:          "pointer",
						}, []gophront.Element{
							gophront.Plain{Content: "INCREMENT"},
						}),
				}),
			}),
		}, "")

	renderedHomePage, err := homePage.Render()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "html", []byte(renderedHomePage))
	return
}
```

Upon visiting the home page, the session's current count is retrieved. If it hasn't been set, it's initialized to zero. 

Then, a dynamic page layout is constructed using gophront. This layout presents the current count and offers an "INCREMENT" **DynamicContentLink**.

This is where the magic happens. In this DynamicContentLink we specify the endpoint where we will get the new content, and the id of the element whose content will be changed.

### Increment Endpoint ('/increment')

``` go
func Increment(c *gin.Context) {
	session := sessions.Default(c)
	currentCount := session.Get("count")
	i := currentCount.(int)
	i++
	session.Set("count", i)
	err := session.Save()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	incrementedCount := gophront.Plain{Content: fmt.Sprintf("%v", i)}
	content, err := incrementedCount.Render()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Data(http.StatusOK, "html", []byte(content))
	return
}

```

Clicking on the "INCREMENT" DynamicContentLink triggers a request to this endpoint. 

It retrieves the current count, increases it by one, and then saves the updated count to the session. 

The incremented count is returned as HTML content.

### The result

![demo](https://github.com/EdmondTabaku/gophront/assets/54292924/513d11b1-c7c1-421b-bf34-86848d29f243)



## 📘 Documentation

### PairedElement structure

**PairedElement**, represents an HTML element with a start and end tag.

### Ready to use paired elements

- **Anchor**: `A(props, style, content)`
- **Bold**: `B(props, style, content)`
- **Body**: `Body(props, style, content)`
- **Button**: `Button(props, style, content)`
- **Canvas**: `Canvas(props, style, content)`
- **Code**: `Code(props, style, content)`
- **Division**: `Div(props, style, content)`
- **Emphasis**: `Em(props, style, content)`
- **Form**: `Form(props, style, content)`
- **Span**: `Span(props, style, content)`
- **Header 1**: `H1(props, style, content)`
- **Header 2**: `H2(props, style, content)`
- **Header 3**: `H3(props, style, content)`
- **Header 4**: `H4(props, style, content)`
- **Header 5**: `H5(props, style, content)`
- **Header 6**: `H6(props, style, content)`
- **Head**: `Head(props, style, content)`
- **Header Section**: `Header(props, style, content)`
- **HTML**: `Html(props, style, content)`
- **Iframe**: `Iframe(props, style, content)`
- **Label**: `Label(props, style, content)`
- **List Item**: `Li(props, style, content)`
- **Navigation**: `Nav(props, style, content)`
- **No Script**: `NoScript(props, style, content)`
- **Ordered List**: `Ol(props, style, content)`
- **Option**: `Option(props, style, content)`
- **Paragraph**: `P(props, style, content)`
- **Script**: `Script(props, style, content)`
- **Section**: `Section(props, style, content)`
- **Select**: `Select(props, style, content)`
- **Strong**: `Strong(props, style, content)`
- **Table**: `Table(props, style, content)`
- **Table Body**: `Tbody(props, style, content)`
- **Table Data**: `Td(props, style, content)`
- **Textarea**: `Textarea(props, style, content)`
- **Table Footer**: `Tfoot(props, style, content)`
- **Table Header**: `Thead(props, style, content)`
- **Title**: `Title(props, style, content)`
- **Table Row**: `Tr(props, style, content)`
- **Unordered List**: `Ul(props, style, content)`

Each function takes:

- `props`: Properties of the element.
- `style`: Styling information.
- `content`: Nested elements or content.

### Special Elements

- **Dynamic Link**: `DynamicLink(route, props, style, content)` - A link that dynamically updates the page body without reloading.
- **Dynamic Content Link**: `DynamicContentLink(route, contentId, props, style, content)` - A link that dynamically updates specified content without reloading the entire page.


### SelfClosingElement structure

**SelfClosingElement**, represents an HTML element with only one tag.

### Ready to use self closing elements


- **Break**: `Br(props, style)`
- **Column**: `Col(props, style)`
- **Horizontal Rule**: `Hr(props, style)`
- **Image**: `Img(props, style)`
- **Input**: `Input(props, style)`
- **Link**: `Link(props, style)`
- **Meta**: `Meta(props, style)`

Each function takes:

- `props`: Properties of the element.
- `style`: Styling information.

## DynamicPage Structure

- `Name`: Name of the page.
- `Header`: List of HTML elements to be rendered in the header section of the page.
- `Content`: List of HTML elements to be rendered in the main content area.
- `Style`: Styling information for the body of the page.
- `Js`: Custom JavaScript to be added to the page.
- `Id`: A unique identifier for the body tag (generated based on the current time).

The `NewDynamicPage` function can be used to create a new dynamic page with a specified name, style, header elements, content elements, and custom JavaScript.

### Rendering

The `Render` method returns the full HTML string representation of the dynamic page. If there are errors rendering any of the header or content elements, it will return an error.

The rendered page includes:

1. A doctype and a head section containing charset meta tag and the provided header elements.
2. A body section with a unique id and provided style. Inside the body, there's a div with id "app" to hold the main content.
3. JavaScript functions (`handleDynamicLinkClick`, `fetchContent`, and `updateContent`) to handle dynamic link clicks and content fetching. Custom provided JavaScript is also included.

### Dynamic Content Loading

The provided JavaScript functions enable the dynamic page to:

1. Handle click events on dynamic links.
2. Fetch content based on the provided route.
3. Update the main content area (with id "app") with fetched content.

## StaticPage Structure

- `Name`: Name of the page.
- `Header`: List of HTML elements to be rendered in the header section of the page.
- `Content`: List of HTML elements to be rendered in the main content area.
- `Style`: Styling information for the body of the page.
- `Js`: Custom JavaScript to be added to the page.
- `Id`: A unique identifier for the body tag (generated based on the current time).

The `NewStaticPage` function can be used to create a new static page with a specified name, style, header elements, content elements, and custom JavaScript.

### Rendering

The `Render` method returns the full HTML string representation of the static page. If there are errors rendering any of the header or content elements, it will return an error.

The rendered page includes:

1. A doctype and a head section containing charset meta tag and the provided header elements.
2. A body section with a unique id and provided style. Inside the body, there's a div with id "app" to hold the main content.
3. Custom provided JavaScript (if any).

### No Dynamic Content Loading

Unlike the `DynamicPage`, the `StaticPage` does not have built-in support for dynamically loading or updating content. Once the page is rendered and loaded, its content remains static and does not change.

Use the `StaticPage` structure and related functions for simple and straightforward static HTML pages without any dynamic content updates.


## Style Structure

The `Style` struct encompasses the following fields:

- `BackgroundColor`: Sets the background color.
- `BackgroundImage`: Defines a background image.
- `Border`: Specifies the border styling.
- `BorderRadius`: Used for rounded corners.
- `Bottom`: Position from the bottom.
- `BoxShadow`: Adds shadow effects around an element's frame.
- `Color`: Sets the color of text.
- `Cursor`: Defines the type of cursor to be displayed.
- `Display`: Specifies the display value.
- `Float`: Specifies the floating behavior.
- `FontFamily`: Sets the font for text.
- `FontSize`: Sets the size of the font.
- `FontWeight`: Sets the thickness of characters in text.
- `Height`: Specifies the height.
- `Left`: Position from the left.
- `LineHeight`: Sets the height of line boxes.
- `Margin`: Specifies the margins of an element.
- `Opacity`: Sets the opacity.
- `Padding`: Specifies the padding of an element.
- `Position`: Specifies the type of positioning.
- `Right`: Position from the right.
- `TextAlign`: Aligns the inner content.
- `TextDecoration`: Specifies the decoration of text.
- `TextTransform`: Controls the capitalization of text.
- `Top`: Position from the top.
- `Transition`: Sets the transition effects.
- `Width`: Specifies the width.
- `ZIndex`: Specifies the z-order of an element.

**Conversion to String**

The `ToString` method converts the `Style` instance into its CSS string representation. Any field with a non-empty value gets appended to the resulting CSS string in the format: `property:value;`.







## 🤝 Contributing
Contributions are welcomed, whether it's bug fixes, new features, or improvements to the documentation.



