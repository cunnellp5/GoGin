package main

import (
  "github.com/gin-gonic/gin"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

	"net/http"
	"os"
)


type Butt struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string
	Firm        string
	Style       string
}

// this is the function that runs the program
func main() {
    lab := os.Getenv("MONGOLAB_URI")
  	db := os.Getenv("myButts")


    port := os.Getenv("PORT")
    if port == "" {
      port = "3000"
    }

  // initialize a gin server
	r := gin.Default()
	r.LoadHTMLGlob("*.html")
	r.Static("/public", "public")

  session, err := mgo.Dial(lab)
  col := session.DB(db).C("myButts")
  if err != nil {
    panic(err)
  }

  r.GET("/", func(c *gin.Context) {
		var butts []Butt
		col.Find(nil).All(&butts)
		c.JSON(http.StatusOK, gin.H{
			"butts": butts,
		})
	})


	r.POST("/", func(c *gin.Context) {
		name := c.PostForm("name")
		firm := c.PostForm("firm")
		style := c.PostForm("style")

		err = col.Insert(&Butt{Name: name, Firm: firm, Style: style})
		if err != nil {
			panic(err)
		}
    c.JSON(http.StatusOK, gin.H{
			"plesa": "work",
		})
    // c.Redirect(http.StatusMovedPermanently, "/")
	})

  // start the server
	r.Run(":" + port)
}
