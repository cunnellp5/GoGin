// package main
//
// import "fmt"
//
// func main() {
//     fmt.Println("Hello world!")
// }

package main

// here we define our imports
import (
  // import in standard libraries
	"net/http"
	"os"

  // import the web framework
	"github.com/gin-gonic/gin"
)

// all go projects require a main function
// this is the function that runs the program
func main() {
  // initialize a gin server
	r := gin.Default()

  // load html files to be rendered
	r.LoadHTMLGlob("*.html")
  // parse through static files to be served
  // static files include our js and css
	r.Static("/public", "public")

  // define the port our server will be running on
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

  // define the route path and response
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"HelloMessage": "Phil is Awesome and i like the color rebecca!",
		})
	})

  // start the server
	r.Run(":" + port)
}
