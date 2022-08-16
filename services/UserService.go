package services

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/callummclu/Gocial-Media-Platform/auth"
	"github.com/callummclu/Gocial-Media-Platform/configs"
	"github.com/callummclu/Gocial-Media-Platform/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.ReturnedUser
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

	rows, err := db.Query("select username, name, surname, display_image, description from users where strpos(username, $1) > 0 OR strpos(name, $1) > 0 OR strpos(surname, $1) > 0 LIMIT $2 OFFSET $3", query, limit, offset)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	var results int

	results_sum, err := db.Query("select count(*) from users where strpos(username, $1) > 0 OR strpos(name, $1) > 0 OR strpos(surname, $1) > 0", query)

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
			Username     string
			Name         string
			Surname      string
			DisplayImage string
			Description  string
		)

		if err := rows.Scan(&Username, &Name, &Surname, &DisplayImage, &Description); err != nil {
			fmt.Print(err)
		}

		users = append(users, models.ReturnedUser{
			Username:     Username,
			Surname:      Surname,
			Name:         Name,
			DisplayImage: DisplayImage,
			Description:  Description,
		})
	}

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot get users"})
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

	c.JSON(200, gin.H{"data": users, "results": results, "pages": pages})
}

func GetUserByUsername(c *gin.Context) {
	var user models.ReturnedUser

	var username string = c.Param("username")
	err := user.GetUserByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": "Cannot find user"})
		return
	}

	c.JSON(200, gin.H{"data": user})

}

func CreateNewUser(c *gin.Context) {
	user := models.NewUser()
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if strings.TrimSpace(user.Name) == "" && strings.TrimSpace(user.Surname) == "" && strings.TrimSpace(user.Username) == "" && strings.TrimSpace(user.Email) == "" && strings.TrimSpace(user.Password) == "" {
		c.JSON(400, gin.H{
			"error": "Fill in all the require fields",
		})
		return
	}
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if checkmail.ValidateFormat(user.Email) != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email format",
		})
		return
	}
	err := user.SaveUser()
	if err != nil {
		if err.Error() == "User already exists" {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "User registered successfully",
	})
}

func DeleteOneUser(c *gin.Context) {
	var user models.LogInUser
	c.BindJSON(&user)

	db, err := configs.GetDB()

	if err != nil {
		c.JSON(400, gin.H{"error": "DB Failed to Connect"})
		return
	}

	stmt, err := db.Prepare("SELECT username,password FROM users WHERE username = $1 OR email = $1")
	if err != nil {
		c.JSON(400, gin.H{"error": "Could not find user"})
		return
	}
	defer stmt.Close()
	var (
		username string
		password string
	)
	err = stmt.QueryRow(user.User).Scan(&username, &password)

	if err != nil {
		c.JSON(400, gin.H{"error": "couldnt find user"})
	}

	err = auth.CheckPassword(password, user.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong password"})
		return
	} else {
		row, err := db.Query("delete from users where username = $1", user.User)

		if err != nil {
			c.JSON(400, gin.H{"error": "user doesnt exist"})
			return
		}

		defer row.Close()
		c.JSON(200, gin.H{"message": "user successfully deleted"})
		return
	}

	c.JSON(400, gin.H{"error": err})
	return
}

func EditOneUser(c *gin.Context) {}

// May go into invitationsService.go
func GetAllSentInvitations(c *gin.Context) {

	var username string = c.Param("username")
	received_invitations, err := models.GetAllSentInvitations(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"data": received_invitations})
}

func GetAllReceivedInvitations(c *gin.Context) {
	var username string = c.Param("username")
	sent_invitations, err := models.GetRecievedInvitationsByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"data": sent_invitations})
}

func GetAllInvitations(c *gin.Context) {
	var username string = c.Param("username")
	sent_invitations, err := models.GetRecievedInvitationsByUsername(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	received_invitations, err := models.GetAllSentInvitations(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"sent": sent_invitations, "received": received_invitations})
}

func SendInvitation(c *gin.Context) {

	// Check sentUsername is not already in username.friends

	token := c.Param("token")
	username := c.Param("username")
	sentUsername := c.Param("sentUsername")

	err := models.SendUserInvitation(username, sentUsername, token)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"message": "request sent successfully"})

}

func AcceptFriendRequest(c *gin.Context) {

	// Check token is valid

	// GET SENT reuquests from username
	// GET Received requests from sentUsername

	// sent_requests = remove sentUsername
	// received_requests = remove username

	// UPDATE users SET sent_requests = sent_requests WHERE username=username
	// UPDATE users SET received_requests = received_requests WHERE username=sentUsername

	// GET friends from username
	// GET friends from sentUsername

	// friends_username = append(friends_username, sentUsername)
	// friends_sentUsername = append(friends_sentUsername, username)

	// UPDATE users SET friends = friends_username WHERE username=username
	// UPDATE users SET friends = friends_sentUsername WHERE username=sentUsername

}

func DeleteInvitation(c *gin.Context) {

	// Check token is valid

	// GET SENT reuquests from username
	// GET Received requests from sentUsername

	// sent_requests = remove sentUsername
	// received_requests = remove username

	// UPDATE users SET sent_requests = sent_requests WHERE username=username
	// UPDATE users SET received_requests = received_requests WHERE username=sentUsername
}

func GetUsersFriends(c *gin.Context) {
	var username string = c.Param("username")
	friends, err := models.GetAllFriends(username)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"friends": friends})

}

func DeleteFriend(c *gin.Context) {
	// Check token is valid

	// GET friends from username
	// GET friends from sentUsername

	// friends_username = remove sentUsername
	// friends_sentUsername = remove username

	// UPDATE users SET friends = friends_username WHERE username=username
	// UPDATE users SET friends = friends_sentUsername WHERE username=sentUsername
}
