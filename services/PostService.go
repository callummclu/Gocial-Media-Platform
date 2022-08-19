package services

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/callummclu/Gocial-Media-Platform/middleware"
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func GetPostsByUsername(query string) (posts []models.Post, e error) {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, title, content, created_at, id, likes FROM posts WHERE username = $1 ORDER BY created_at DESC", query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var p []models.Post

	for rows.Next() {
		var (
			Username  string
			Title     string
			Content   string
			CreatedAt string
			Id        int64
			Likes     []string
		)

		if err := rows.Scan(&Username, &Title, &Content, &CreatedAt, &Id, pq.Array(&Likes)); err != nil {
			fmt.Print(err)
		}

		p = append(p, models.Post{
			Username:  Username,
			Title:     Title,
			Content:   Content,
			CreatedAt: CreatedAt,
			ID:        Id,
			Likes:     Likes,
		})
	}

	return p, err
}

/*

	THIS EDIT POST FUNCTION IS UNFINISHED

*/

func EditPostById(id int64, token string, username string) error {
	err := auth.CheckJWT(token, &username)

	if err != nil {
		err = errors.New("invalid user")
		return err
	}

	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	// _, err = db.Query("DELETE FROM posts WHERE id = $1", id)

	return err

}

func DeletePostById(id int64, token string, username string) error {

	err := auth.CheckJWT(token, &username)

	if err != nil {
		err = errors.New("invalid user")
		return err
	}

	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	_, err = db.Query("DELETE FROM posts WHERE id = $1", id)

	return err
}

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

	rows, err := db.Query("select username,title,content,created_at,id,likes from posts where strpos(username, $1) > 0 OR strpos(title, $1) > 0 OR strpos(content, $1) > 0 ORDER BY created_at DESC LIMIT $2 OFFSET $3 ", query, limit, offset)

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
			Likes     []string
		)

		if err := rows.Scan(&Username, &Title, &Content, &CreatedAt, &ID, pq.Array(&Likes)); err != nil {
			fmt.Print(err)
		}

		posts = append(posts, models.Post{
			Username:  Username,
			Title:     Title,
			Content:   Content,
			CreatedAt: CreatedAt,
			ID:        ID,
			Likes:     Likes,
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

	posts, err := GetPostsByUsername(username)

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

	err = DeletePostById(id, token, username)

	if err != nil {
		c.JSON(400, gin.H{"error": "cannot find post"})
	}

	c.JSON(200, gin.H{"message": "post succesfully deleted"})
}

func EditOnePost(c *gin.Context) {
	str_id := c.Param("id")

	token := c.Param("token")

	username := c.Param("username")

	id, err := strconv.ParseInt(str_id, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "user does not exist"})
	}

	err = EditPostById(id, token, username)

	if err != nil {
		c.JSON(400, gin.H{"error": "cannot find post"})
	}

	c.JSON(200, gin.H{"message": "edited post successfully"})
}

func LikePostById(id int64, token string, username string) error {
	err := auth.CheckJWT(token, &username)

	if err != nil {
		err = errors.New("invalid user")
		return err
	}

	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		return err
	}
	defer db.Close()

	var liked_posts_ids []string

	get_likes_stmt := "SELECT likes FROM users WHERE username = $1"

	if err := db.QueryRow(get_likes_stmt, username).Scan(pq.Array(&liked_posts_ids)); err != nil {
		return err
	}

	var posts_likes []string

	get_posts_likes_stmt := "SELECT likes FROM posts WHERE id = $1"

	if err := db.QueryRow(get_posts_likes_stmt, id).Scan(pq.Array(posts_likes)); err != nil {
		return err
	}

	if posts_likes == nil {
		posts_likes = []string{}
	}

	if liked_posts_ids == nil {
		liked_posts_ids = []string{}
	}

	if middleware.Contains(liked_posts_ids, strconv.FormatInt(id, 10)) {
		liked_posts_ids = middleware.Remove(liked_posts_ids, strconv.FormatInt(id, 10))
		posts_likes = middleware.Remove(posts_likes, username)

	} else {
		liked_posts_ids = append(liked_posts_ids, strconv.FormatInt(id, 10))
		posts_likes = append(posts_likes, username)
	}

	save_users_likes := "UPDATE users SET likes = $1 WHERE username = $2"
	save_posts_likers := "UPDATE posts SET likes = $1 WHERE id = $2"

	_, err = db.Exec(save_users_likes, pq.Array(liked_posts_ids), username)

	if err != nil {
		fmt.Print("username")
		return err
	}

	_, err = db.Exec(save_posts_likers, pq.Array(posts_likes), id)

	if err != nil {
		fmt.Print("sentUsername")
		return err
	}

	return err
}

func LikePost(c *gin.Context) {

	str_id := c.Param("id")

	token := c.Param("token")

	username := c.Param("username")

	id, err := strconv.ParseInt(str_id, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"error": "user does not exist"})
	}

	err = LikePostById(id, token, username)

	if err != nil {
		c.JSON(400, gin.H{"error": "cannot find post"})
	}

	c.JSON(200, gin.H{"message": "liked post successfully"})
}

func getLikedPostsByUsernameHelper(username string) (l_posts []string, e error) {
	db, err := configs.GetDB()
	if err != nil {
		fmt.Println("DB")

		err = errors.New("DB connection error")
		return nil, err
	}
	defer db.Close()

	var liked_posts []string

	stmt := "SELECT likes FROM users WHERE username=$1"

	if err := db.QueryRow(stmt, username).Scan(pq.Array(&liked_posts)); err != nil {
		fmt.Println("cant find user")
		return nil, err
	}

	return liked_posts, err
}

func getPostsByIdList(ids []string) ([]models.Post, error) {
	db, err := configs.GetDB()
	if err != nil {
		err = errors.New("DB connection error")
		fmt.Println("DB")

		return nil, err
	}
	defer db.Close()

	likedPostsIds := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ", "), "")

	likedPostsIds = strings.ReplaceAll(likedPostsIds, "[", "")
	likedPostsIds = strings.ReplaceAll(likedPostsIds, "]", "")

	rows, err := db.Query("SELECT username,title,content,created_at,id,likes from posts where id = ANY($1::int[]) ORDER BY created_at DESC", ("{" + likedPostsIds + "}"))

	if err != nil {
		fmt.Println("Query")

		return nil, err
	}

	defer rows.Close()

	var p []models.Post

	for rows.Next() {
		var (
			Username  string
			Title     string
			Content   string
			CreatedAt string
			Id        int64
			Likes     []string
		)

		if err := rows.Scan(&Username, &Title, &Content, &CreatedAt, &Id, pq.Array(&Likes)); err != nil {
			fmt.Println("Scan")

			fmt.Print(err)
		}

		p = append(p, models.Post{
			Username:  Username,
			Title:     Title,
			Content:   Content,
			CreatedAt: CreatedAt,
			ID:        Id,
			Likes:     Likes,
		})
	}

	return p, err
}

func GetLikedPostsByUsername(c *gin.Context) {

	username := c.Param("username")

	likedPostsIds, err := getLikedPostsByUsernameHelper(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	posts, err := getPostsByIdList(likedPostsIds)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"posts": posts})

}
