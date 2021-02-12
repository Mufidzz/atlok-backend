package main

import (
	"./config"
	"./controller"
	sec "./security"
	j "./security/jwt"
	"crypto/sha1"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func getCookieStore() []byte {
	h := sha1.New()
	return h.Sum([]byte(time.Now().String()))
}

func main() {
	db := config.DBInit()
	InDB := &controller.InDB{DB: db}
	router := gin.Default()
	router.Use(sessions.Sessions("U-ACCESS-SESSX", cookie.NewStore(getCookieStore())))
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
	}))

	gin.SetMode(gin.ReleaseMode)

	auth := router.Group("/auth")
	{
		auth.POST("/login", InDB.Login)
		auth.POST("/register", InDB.CreateUser)
	}

	user := router.Group("/user")
	user.Use(sec.JWTBarrier(1))
	{
		user.GET("/unverified/:start/:count", InDB.GetUnverifiedUsers)
		user.GET("/verify/:id", InDB.VerifyUser)
		user.DELETE("/:id", InDB.DeleteUser)
	}

	customer := router.Group("/customer")
	{
		customer.Use(sec.JWTBarrier(0)).GET("/:q1", InDB.GetCustomerUsingID)
		customer.Use(sec.JWTBarrier(0)).GET("/:q1/search/:start/:count", InDB.SearchCustomers)

		customer.Use(sec.JWTBarrier(1)).POST("/", InDB.CreateCustomer)
		customer.Use(sec.JWTBarrier(1)).PUT("/:q1", InDB.UpdateCustomer)
		customer.Use(sec.JWTBarrier(1)).DELETE("/:q1", InDB.DeleteCustomer)
	}

	customerChange := router.Group("/customer-change")
	{
		customerChange.Use(sec.JWTBarrier(0)).POST("/:q1", InDB.CreateCustomerChange)
		customerChange.Use(sec.JWTBarrier(1)).PATCH("/:q1/accept", InDB.AcceptCustomerChange)
		customerChange.Use(sec.JWTBarrier(1)).PATCH("/:q1/deny", InDB.DenyCustomerChange)

		customerChange.Use(sec.JWTBarrier(1)).GET("/:q1/:q2", InDB.GetCustomerChanges)
		customerChange.Use(sec.JWTBarrier(1)).GET("/:q1", InDB.GetCustomerChangeUsingID)

	}

	powerRating := router.Group("/power-rating")
	{
		powerRating.Use(sec.JWTBarrier(0)).GET("/", InDB.GetPowerRatings)
		powerRating.Use(sec.JWTBarrier(0)).GET("/:q1/search/:start/:count", InDB.SearchPowerRates)
		powerRating.Use(sec.JWTBarrier(1)).POST("/", InDB.CreatePowerRating)
		powerRating.Use(sec.JWTBarrier(1)).PUT("/:q1", InDB.UpdatePowerRating)
		powerRating.Use(sec.JWTBarrier(1)).DELETE("/:q1", InDB.DeletePowerRating)
	}

	substation := router.Group("/substation")
	{
		substation.Use(sec.JWTBarrier(0)).GET("/:q1", InDB.GetSubstationUsingID)
		substation.Use(sec.JWTBarrier(0)).GET("/:q1/search/:start/:count", InDB.SearchSubstations)

		substation.Use(sec.JWTBarrier(1)).GET("/", InDB.GetSubstations)
		substation.Use(sec.JWTBarrier(1)).POST("/", InDB.CreateSubstation)
		substation.Use(sec.JWTBarrier(1)).PUT("/:q1", InDB.UpdateSubstation)
		substation.Use(sec.JWTBarrier(1)).DELETE("/:q1", InDB.DeleteSubstation)
	}

	debug := router.Group("/token/generate")
	{
		debug.GET("/:acs", func(context *gin.Context) {

			a, _ := strconv.Atoi(context.Param("acs"))

			tok, _ := j.Generate(j.Body{
				Uid: 1,
				Iat: time.Now().Unix(),
				Acs: uint(a),
			})

			context.JSON(200, tok)
		})
	}

	err := router.Run(":50006")
	if err != nil {
		panic(err)
	}
}
