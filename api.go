package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	key       = Key
	pw        = Pw
	esv       = "eng-ESV"
	msg       = "eng-MSG"
	https     = "https://"
	biblesUrl = https + key + ":" + pw + "@" + "bibles.org/v2"
)

var versions = []string{esv, msg}

type Book struct {
	Id   string `json:"abbr"`
	Name string `json:"name"`
}

type Chapter struct {
	Chapter string `json:"chapter"`
}

func main() {

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "not found"})
	})
	// list books
	r.GET("/books", func(c *gin.Context) {
		r, err := http.Get(biblesUrl + "/versions/eng-ESV/books.js")
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": "error talking to bibles.org"})
			return
		}
		defer r.Body.Close()
		var response struct {
			Response struct {
				Books []Book `json:"books"`
			} `json:"response"`
		}
		if err = json.NewDecoder(r.Body).Decode(&response); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": "error decoding data from bibles.org"})
			return
		}
		c.JSON(200, response.Response.Books)
	})
	r.GET("/versions", func(c *gin.Context) {
		c.JSON(200, versions)
	})
	// list chapters
	r.GET("/versions/:version/:book", func(c *gin.Context) {
		version := c.Param("version")
		book := c.Param("book")
		r, err := http.Get(biblesUrl + "/books/eng-" + version + ":" + book + "/chapters.js")
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": "error talking to bibles.org"})
			return
		}
		var response struct {
			Response struct {
				Chapters []Chapter `json:"chapters"`
			} `json:"response"`
		}
		if err = json.NewDecoder(r.Body).Decode(&response); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": "error decoding data from bibles.org"})
			return
		}
		var chapters []int
		for _, ch := range response.Response.Chapters {
			chap, err := strconv.Atoi(ch.Chapter)
			if err != nil {
				log.Println(err)
				c.JSON(500, gin.H{"error": "WAT"})
				return
			}
			chapters = append(chapters, chap)
		}
		c.JSON(200, chapters)
	})
	// list verses of chapter
	r.GET("/versions/:version/:book/:chapter", func(c *gin.Context) {
		version := c.Param("version")
		chapter := c.Param("chapter")
		book := c.Param("book")
		r, err := http.Get(biblesUrl + "/chapters/eng-" + version + ":" + book + "." + chapter + "/verses.js")
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": "error talking to bibles.org"})
			return
		}
		var response struct {
			Response struct {
				Verses []Verse `json:"verses"`
			} `json:"response"`
		}
		if err = json.NewDecoder(r.Body).Decode(&response); err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"error": "error decoding data from bibles.org"})
			return
		}
		c.JSON(200, response.Response.Verses)
	})
	port := "8111"
	log.Println("Serving http://localhost:" + port)
	r.Run("0.0.0.0:" + port)
}

type Verse struct {
	Verse string `db:"verse"`
	Text  string `db:"text"`
	// TODO: Next and Prev
}
