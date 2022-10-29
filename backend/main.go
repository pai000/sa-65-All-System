package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/controller"
	"github.com/pai000/sa-65-project/entity"
	"github.com/pai000/sa-65-project/middlewares"
)

//const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// ============================================================================================================
			// ============================================================================================================
			// ============================================ Employee ======================================================
			// ============================================================================================================
			// ============================================================================================================

			//gender routes
			r.GET("/genders", controller.ListGenders)
			r.GET("/gender/:id", controller.GetGender)
			r.POST("/genders", controller.CreateGender)
			r.PATCH("/genders", controller.UpdateGender)
			r.DELETE("/genders/:id", controller.DeleteGender)

			//job_position routes
			r.GET("/job_positions", controller.ListJob_Position)
			r.GET("/job_position/:id", controller.GetJob_Position)
			r.POST("/job_positions", controller.CreateJob_Position)
			r.PATCH("/job_positions", controller.UpdateJob_Position)
			r.DELETE("/job_positions/:id", controller.DeleteJob_Position)

			//province routes
			r.GET("/provinces", controller.ListProvince)
			r.GET("/province/:id", controller.GetProvince)
			r.POST("/provinces", controller.CreateProvince)
			r.PATCH("/provinces", controller.UpdateProvince)
			r.DELETE("/provinces/:id", controller.DeleteProvince)

			//role routes
			r.GET("/employees", controller.ListEmployee)
			r.GET("/employee/:id", controller.GetEmployee)
			r.POST("/employees", controller.CreateEmployee)
			r.PATCH("/employees", controller.UpdateEmployee)
			r.DELETE("/employees/:id", controller.DeleteEmployee)

			// ============================================================================================================
			// ============================================================================================================
			// ============================================ Student ======================================================
			// ============================================================================================================
			// ============================================================================================================

			// Role Routes
			r.GET("/roles", controller.ListRole)
			r.GET("/role/:id", controller.GetRole)
			r.POST("/roles", controller.CreateRole)

			// program Router
			r.GET("programs", controller.ListProgram)
			r.GET("/program/:id", controller.GetProgram)
			r.POST("/programs", controller.CreateProgram)

			// student Routes
			r.GET("/students", controller.ListStudent)
			r.GET("/student/:id", controller.GetStudent)
			r.POST("/students", controller.CreateStudent)

			// ============================================================================================================
			// ============================================================================================================
			// ============================================== Room ========================================================
			// ============================================================================================================
			// ============================================================================================================
			r.GET("/Room_types", controller.ListRoom_types)
			r.GET("/Room_type/:id", controller.GetRoom_type)
			r.POST("/Room_types", controller.CreateRoom_type)

			r.GET("/Room_prices", controller.ListRoom_prices)
			r.GET("/Room_price/:id", controller.GetRoom_price)
			r.POST("/Room_prices", controller.CreateRoom_price)

			r.GET("/Set_of_furnitures", controller.ListSet_of_furnitures)
			r.GET("/Set_of_furniture/:id", controller.GetSet_of_furniture)
			r.POST("/Set_of_furnitures", controller.CreateSet_of_furniture)

			r.GET("/Furnitures", controller.ListFurniture)
			r.GET("/Furniture/:id", controller.GetFurniture)
			r.POST("/Furnitures", controller.CreateFurniture)

			r.GET("/Rooms", controller.ListRoom)
			r.GET("/Room/:id", controller.GetRoom)
			r.POST("/Rooms", controller.CreateRoom)

			// ============================================================================================================
			// ============================================================================================================
			// ============================================== Repair =======================================================
			// ============================================================================================================
			// ============================================================================================================

			r.GET("/Repairs", controller.ListRepairs)
			r.GET("/Repair/:id", controller.GetRepair)
			r.POST("/Repairs", controller.CreateRepair)

			// ============================================================================================================
			// ============================================================================================================
			// =========================================== Book ==========================================================
			// ============================================================================================================
			// ============================================================================================================

			r.GET("/Times", controller.ListTimes)
			r.GET("/Time/:id", controller.GetTime)
			r.POST("/Times", controller.CreateTime)

			r.GET("/Bookings", controller.ListBookings)
			r.GET("/Booking/:id", controller.GetBooking)
			r.POST("/Bookings", controller.CreateBooking)

			// ============================================================================================================
			// ============================================================================================================
			// ========================================== Borrow ==========================================================
			// ============================================================================================================
			// ============================================================================================================

			r.GET("/Equipments", controller.ListEquipments)
			r.GET("/Equipment/:id", controller.GetEquipment)
			r.POST("/Equipments", controller.CreateEquipment)

			r.GET("/Borrow_cards", controller.ListBorrow_cards)
			r.GET("/Borrow_card/:id", controller.GetBorrow_card)
			r.POST("/Borrow_cards", controller.CreateBorrow_card)

			r.GET("/List_datas", controller.ListList_data)
			r.GET("/List_data/:id", controller.GetList_data)
			r.POST("/List_datas", controller.CreateList_data)

			// ============================================================================================================
			// ============================================================================================================
			// ========================================= Payment ==========================================================
			// ============================================================================================================
			// ============================================================================================================

			// Semester Routes
			router.POST("/semesters", controller.CreateSemester)
			router.GET("/semester/:id", controller.GetSemester)
			router.GET("/semesters", controller.ListSemesters)

			// Payment_Bill Routes
			router.POST("/payment_bills", controller.CreatePayment_Bill)
			router.GET("/payment_bills", controller.ListPayment_Bills)
			router.GET("/payment_bill/:id", controller.GetPayment_Bill)

		}
	}

	//Signup User Route
	r.POST("/signup_employee", controller.CreateEmployee)

	r.POST("/signup_student", controller.CreateLoginStudent)
	// login User Route
	r.POST("/login_employee", controller.LoginEmployee)

	r.POST("/login_student", controller.LoginStudent)

	// Run the server go run main.go
	r.Run("0.0.0.0:8080")
}

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
