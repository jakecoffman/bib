package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"os"
	"sync"

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

type Verse struct {
	Verse string `db:"verse"`
	Text  string `db:"text"`
	// TODO: Next and Prev
}

type Cache struct {
	sync.RWMutex
	Data map[string][]Verse
}

func (c *Cache) Put(key string, verses []Verse) {
	c.Lock()
	defer c.Unlock()
	c.Data[key] = verses

	f, err := os.OpenFile(cacheFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}

	defer f.Close()
	if err = json.NewEncoder(f).Encode(c.Data); err != nil {
		log.Println(err)
	}

	log.Println("OK!")
}

func (c *Cache) Get(key string) []Verse {
	c.RLock()
	defer c.RUnlock()
	return c.Data[key]
}

var cache = Cache{
	RWMutex: sync.RWMutex{},
	Data:    map[string][]Verse{},
}

const cacheFile = "./bibleCache.json"

func init() {
	f, err := os.Open(cacheFile)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&cache.Data); err != nil {
		log.Println(err)
	}
}

func main() {
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "not found"})
	})
	r.GET("/books", ListBooks)
	r.GET("/versions", ListVersions)
	r.GET("/versions/:version/:book", ListChapters)
	r.GET("/versions/:version/:book/:chapter", ListVerses)
	port := "8111"
	log.Println("Serving http://localhost:" + port)
	r.Run("0.0.0.0:" + port)
}

func ListBooks(c *gin.Context) {
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
}

func ListVersions(c *gin.Context) {
	c.JSON(200, versions)
}

func ListChapters(c *gin.Context) {
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
			c.JSON(400, gin.H{"error": "chapter is not a number"})
			return
		}
		chapters = append(chapters, chap)
	}
	c.JSON(200, chapters)
}

func ListVerses(c *gin.Context) {
	version := c.Param("version")
	chapter := c.Param("chapter")
	book := c.Param("book")
	key := version + "/" + book + "/" + chapter

	if verses := cache.Get(key); verses != nil {
		c.JSON(200, verses)
		return
	}

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

	cache.Put(key, response.Response.Verses)

	c.JSON(200, response.Response.Verses)
}
