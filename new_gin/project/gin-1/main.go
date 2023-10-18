// package main

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type User struct {
// 	FName    string `json:"f_name"`
// 	Password string `json:"-"`
// 	LName    string `json:"l_name"`
// 	Email    string `json:"email"`
// }

// var users = map[uint64]User{
// 	123: {
// 		FName:    "Bob",
// 		LName:    "abc",
// 		Password: "someSecretPassword",
// 		Email:    "bob@email.com",
// 	},
// }

// func FetchUser(userId uint64) (User, error) {
// 	userdata, ok := users[userId]
// 	if !ok {
// 		return User{}, errors.New("user not found")
// 	}
// 	return userdata, nil
// }
// func Home(c *gin.Context) {

// 	id := c.Param("userId")
// 	fmt.Println(id, "*********8")
// 	//c.String(http.StatusOK, "this is my home page")
// 	//using the map to send the json response
// 	userid, err := strconv.ParseUint(id, 10, 64)
// 	if err != nil {
// 		log.Println(err)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
// 		return
// 	}
// 	userdetials, err := FetchUser(userid)
// 	if err != nil {
// 		log.Println(err)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
// 		return
// 	}
// 	fmt.Print(userdetials)

// 	c.JSON(http.StatusOK, userdetials)
// 	// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
// }
// func main() {
// 	router := gin.Default()

//		router.GET("/home/:userId", Home)
//		router.Run(":8081")
//	}
package main

import (
	"errors"

	"fmt"

	"log"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id uint64

	Name string
}

var students = map[uint64]Student{

	101: Student{

		Id: 101,

		Name: "jeevan",
	},

	102: Student{

		Id: 102,

		Name: "afthab",
	},
}

func main() {

	router := gin.Default()

	router.GET("/home/:user_id", Home)

	router.Run(":8081")

}

// func(*Context)

func Home(c *gin.Context) {

	Stringid := c.Param("user_id")

	fmt.Println(Stringid)

	uid, err := strconv.ParseUint(Stringid, 10, 64)

	if err != nil {

		log.Println("conversion string to int error", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error found at conversion.."})

		return

	}

	val, err := fetchStudent(uid)

	if err != nil {

		log.Println("student not found in map", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "students not found"})

		return

	}

	c.JSON(http.StatusOK, val)

	// fmt.Println(name)

	//c.String(http.StatusOK, "this is my home page")

	//using the map to send the json response

	// c.JSON(http.StatusOK, gin.H{"msg": "this is my home page"})

	// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})

}

func fetchStudent(uid uint64) (Student, error) {

	sData, ok := students[uid]

	if !ok {

		return Student{}, errors.New("data not there")

	}

	return sData, nil

}
