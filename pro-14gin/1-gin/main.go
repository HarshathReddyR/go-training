// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type j struct {
// 	Name string
// }

// func main() {
// 	router := gin.Default()

// 	router.GET("/home", Home)
// 	router.Run(":8080")
// }

// // func(*Context)
//
//	func Home(c *gin.Context) {
//		j1 := j{
//			Name: "Abhi",
//		}
//		//c.String(http.StatusOK, "this is my home page")
//		//using the map to send the json response
//		c.JSON(http.StatusOK, gin.H{"msg": "this is my home page"})
//		c.JSON(http.StatusOK, j1)//here we are passing the structure
//		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{})
//	}
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/home/:name", Home)
	router.Run(":8080")
}

// func(*Context)
func Home(c *gin.Context) {

	name := c.Param("name")
	fmt.Println(name)
	//c.String(http.StatusOK, "this is my home page")
	//using the map to send the json response
	c.JSON(http.StatusOK, gin.H{"msg": "this is my home page"})
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
}
