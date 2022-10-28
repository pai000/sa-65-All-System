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
			router.GET("/genders", controller.ListGenders)
			router.GET("/gender/:id", controller.GetGender)
			router.POST("/genders", controller.CreateGender)
			router.PATCH("/genders", controller.UpdateGender)
			router.DELETE("/genders/:id", controller.DeleteGender)

			//job_position routes
			router.GET("/job_positions", controller.ListJob_Position)
			router.GET("/job_position/:id", controller.GetJob_Position)
			router.POST("/job_positions", controller.CreateJob_Position)
			router.PATCH("/job_positions", controller.UpdateJob_Position)
			router.DELETE("/job_positions/:id", controller.DeleteJob_Position)

			//province routes
			router.GET("/provinces", controller.ListProvince)
			router.GET("/province/:id", controller.GetProvince)
			router.POST("/provinces", controller.CreateProvince)
			router.PATCH("/provinces", controller.UpdateProvince)
			router.DELETE("/provinces/:id", controller.DeleteProvince)

			//role routes
			router.GET("/employees", controller.ListEmployee)
			router.GET("/employee/:id", controller.GetEmployee)
			router.POST("/employees", controller.CreateEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)
			router.DELETE("/employees/:id", controller.DeleteEmployee)

			// ============================================================================================================
			// ============================================================================================================
			// ============================================ Student ======================================================
			// ============================================================================================================
			// ============================================================================================================

			// Role Routes
			router.GET("/roles", controller.ListRole)
			router.GET("/role/:id", controller.GetRole)
			router.POST("/roles", controller.CreateRole)

			// program Routes
			router.GET("programs", controller.ListProgram)
			router.GET("/program/:id", controller.GetProgram)
			router.POST("/programs", controller.CreateProgram)

			// student Routes
			router.GET("/students", controller.ListStudent)
			router.GET("/student/:id", controller.GetStudent)
			router.POST("/students", controller.CreateStudent)

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
