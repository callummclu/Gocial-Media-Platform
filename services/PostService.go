package services

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		c.JSON(400, gin.H{"error": err})
	}
	var query string = c.Query("searchParams")

	limit, err := strconv.Atoi(c.Query("itemsPerPage"))

	if err != nil {
		limit = 20
	}

	offset, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		offset = 1
	}

	if offset < 1 {
		offset = 1
	}

	offset = (offset - 1) * limit

	rows, err := db.Query("select username,title,content,created_at,Id from posts where strpos(username, $1) > 0 OR strpos(title, $1) > 0 OR strpos(content, $1) > 0 ORDER BY created_at DESC LIMIT $2 OFFSET $3 ", query, limit, offset)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	var results int

	results_sum, err := db.Query("select count(*) from posts where strpos(username, $1) > 0 OR strpos(title, $1) > 0 OR strpos(content, $1) > 0", query)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer results_sum.Close()

	for results_sum.Next() {
		if err := results_sum.Scan(&results); err != nil {
			log.Fatal(err)
		}
	}

	for rows.Next() {
		var (
			Username  string
			Title     string
			Content   string
			CreatedAt string
			ID        int64
		)

		if err := rows.Scan(&Username, &Title, &Content, &CreatedAt, &ID); err != nil {
			fmt.Print(err)
		}

		posts = append(posts, models.Post{
			Username:  Username,
			Title:     Title,
			Content:   Content,
			CreatedAt: CreatedAt,
			ID:        ID,
		})
	}

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot get posts"})
		return
	}

	fmt.Print(results, limit, results/limit)
	if limit < 1 {
		limit = 1
	}
	pages := math.Ceil(float64(results / limit))

	if pages == 0 {
		pages = 1
	}

	c.JSON(200, gin.H{"data": posts, "results": results, "pages": pages})
}

func GetPostByUsername(c *gin.Context) {

	var username string = c.Param("username")

	fmt.Println(username)

	posts, err := models.GetPostsByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find user"})
		return
	}

	c.JSON(200, gin.H{"data": posts})
}

func CreateNewPost(c *gin.Context) {
	post := models.NewPost()

	token := c.Param("token")

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err := post.SavePost(token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "Post made successfully",
	})
}

func DeleteOnePost(c *gin.Context) {

	str_id := c.Param("id")

	token := c.Param("token")

	username := c.Param("username")

	id, err := strconv.ParseInt(str_id, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "user does not exist"})
	}

	err = models.DeletePostById(id, token, username)

	if err != nil {
		c.JSON(400, gin.H{"error": "cannot find post"})
	}

	c.JSON(200, gin.H{"message": "post succesfully deleted"})
}
func EditOnePost(c *gin.Context) {}
