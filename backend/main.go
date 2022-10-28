package main

import (
	"github.com/PatiphatPanchad/sa-65-example/controller"

	"github.com/PatiphatPanchad/sa-65-example/entity"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())
	// User Routes

	// r.GET("/users", controller.ListUsers)

	// r.GET("/user/:id", controller.GetUser)

	// r.POST("/users", controller.CreateUser)

	// r.PATCH("/users", controller.UpdateUser)

	// r.DELETE("/users/:id", controller.DeleteUser)
	/////
	r.GET("/Equipments", controller.ListEquipments)
	r.GET("/Equipment/:id", controller.GetEquipment)
	r.POST("/Equipments", controller.CreateEquipment)
	r.PATCH("/Equipments", controller.UpdateEquipment)
	r.DELETE("/Equipments/:id", controller.DeleteEquipment)
	////
	r.GET("/Borrow_cards", controller.ListBorrow_cards)
	r.GET("/Borrow_card/:id", controller.GetBorrow_card)
	r.POST("/Borrow_cards", controller.CreateBorrow_card)
	r.PATCH("/Borrow_cards", controller.UpdateBorrow_card)
	r.DELETE("/Borrow_cards/:id", controller.DeleteBorrow_card)
	////
	// r.GET("/Set_of_furnitures", controller.ListSet_of_furnitures)
	// r.GET("/Set_of_furniture/:id", controller.GetSet_of_furniture)
	// r.POST("/Set_of_furnitures", controller.CreateSet_of_furniture)
	// r.PATCH("/Set_of_furnitures", controller.UpdateSet_of_furniture)
	// r.DELETE("/Set_of_furnitures/:id", controller.DeleteSet_of_furniture)
	////
	r.GET("/Rooms", controller.ListRoom)
	r.GET("/Room/:id", controller.GetRoom)
	r.POST("/Rooms", controller.CreateRoom)
	r.PATCH("/Rooms", controller.UpdateRoom)
	r.DELETE("/Rooms/:id", controller.DeleteRoom)

	r.GET("/List_datas", controller.ListList_data)
	r.GET("/List_data/:id", controller.GetList_data)
	r.POST("/List_datas", controller.CreateList_data)
	r.PATCH("/List_datas", controller.UpdateList_data)
	r.DELETE("/List_datas/:id", controller.DeleteList_data)
	// Run the server
	r.POST("/signup", controller.CreateLoginEmployee)
	// login User Route
	r.POST("/login", controller.Login)

	r.Run("0.0.0.0:8080")

}
