package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Author struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Place     string `json:"place"`
	AvatarUrl string `json:"avatar_url"`
}

type Post struct {
	Id          int    `json:"id"`
	AuthorId    int    `json:"author_id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	ImageUrl    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
	DisplayTime string
}

func main() {

	var authors []Author
	authorsMap := map[int]Author{}

	var posts []Post
	err := getJson("http://maqe.github.io/json/authors.json", &authors)
	if err != nil {
		log.Fatal(err)
	}
	err = getJson("http://maqe.github.io/json/posts.json", &posts)
	if err != nil {
		log.Fatal(err)
	}
	processPosts(posts)

	for _, a := range authors {
		authorsMap[a.Id] = a
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "MAQE Forums",
			"posts":   posts[:8],
			"authors": authorsMap,
			"ModTwo":  ModTwo,
			"N":       N,
		})
	})
	router.Static("/css", "templates/css")

	router.Run(":8080")
}

func processPosts(posts []Post) {

	layout := "2006-01-02 15:04:05"
	for i := 0; i < len(posts); i++ {
		t, err := time.Parse(layout, posts[i].CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		diffMonth := math.Round(time.Now().Sub(t).Hours() / 24 / 30)
		posts[i].DisplayTime = fmt.Sprintf("%.0f months ago", diffMonth)
	}
}

func getJson(url string, target interface{}) error {
	var client = &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func ModTwo(i int) bool {
	return i%2 == 1
}

func N(n int) []int {
	return make([]int, n)
}
