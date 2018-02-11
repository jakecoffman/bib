package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const (
	//key       = Key
	//pw        = Pw
	esv       = "eng-ESV"
	msg       = "eng-MSG"
	https     = "https://"
	biblesUrl = https + "bibles.org"
)

var versions = []string{esv, msg}

type Book struct {
	Id   string `json:"abbr"`
	Name string `json:"name"`
}

type Cache struct {
	sync.RWMutex
	Data map[string]string
}

func (c *Cache) Put(key string, verses string) {
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
}

func (c *Cache) Get(key string) string {
	c.RLock()
	defer c.RUnlock()
	return c.Data[key]
}

var cache = Cache{
	RWMutex: sync.RWMutex{},
	Data:    map[string]string{},
}

const cacheFile = "./bibleCache_2.json"

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
	req, err := http.NewRequest("GET", biblesUrl+"/versions/eng-ESV/books.json", nil)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "error talking to bibles.org"})
		return
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "error talking to bibles.org"})
		return
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "error talking to bibles.org"})
		return
	}
	c.String(200, string(data))
}

func ListVersions(c *gin.Context) {
	c.JSON(200, versions)
}

func ListChapters(c *gin.Context) {
	version := c.Param("version")
	book := c.Param("book")
	r, err := http.Get(biblesUrl + "/books/eng-" + version + ":" + book + "/chapters.json")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "error talking to bibles.org"})
		return
	}
	defer r.Body.Close()
	var response []struct {
		Chapter string `json:"chapter"`
	}
	if err = json.NewDecoder(r.Body).Decode(&response); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "error decoding data from bibles.org"})
		return
	}
	var chapters []int
	for _, ch := range response {
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

	if verses := cache.Get(key); verses != "" {
		c.JSON(200, verses)
		return
	}
	r, err := http.Get(biblesUrl + "/chapters/eng-" + version + ":" + book + "." + chapter + ".json")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "error talking to bibles.org"})
		return
	}
	defer r.Body.Close()
	var response struct {
		Text string `json:"text"`
	}
	if err = json.NewDecoder(r.Body).Decode(&response); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "error decoding data from bibles.org"})
		return
	}

	cache.Put(key, response.Text)

	c.String(200, response.Text)
}
