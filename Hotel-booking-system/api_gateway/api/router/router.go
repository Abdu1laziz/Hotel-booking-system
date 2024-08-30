package router

import (
	"api-gateway/config"
	"api-gateway/internal/connections"
	_ "api-gateway/internal/docs"
	token "api-gateway/utils/jwt"
	"fmt"
	"log"
	"net/http"

	swag "github.com/swaggo/http-swagger"
)

// @title           Booking Hotel API
// @version         2.0
// @description     This is an API for booking Hotels.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8085
// @BasePath /

func NewRouter() {
	c := config.Configuration()

	r := http.NewServeMux()
	handler := connections.NewHandler()

	// Users

	r.HandleFunc("POST /users/register", handler.Register)
	r.HandleFunc("POST /users/verify", handler.Verify)
	r.HandleFunc("POST /users/login", handler.LogIn)
	r.HandleFunc("GET /users/{id}", token.JWTMiddleware(handler.GetUser))
	r.HandleFunc("PUT /users/{id}", token.JWTMiddleware(handler.UpdateUser))
	r.HandleFunc("DELETE /users/{id}", token.JWTMiddleware(handler.DeleteUser))
	r.HandleFunc("POST /users/logout/{id}", token.JWTMiddleware(handler.LogOut))
	r.Handle("/swagger/", swag.WrapHandler)

	// Hotel

	r.HandleFunc("POST /hotels/create", token.JWTMiddleware(handler.CreateHotel))
	r.HandleFunc("POST /hotels/rooms/create", token.JWTMiddleware(handler.CreateRoom))
	r.HandleFunc("GET /hotels", token.JWTMiddleware(handler.GetHotels))
	r.HandleFunc("GET /hotels/{id}", token.JWTMiddleware(handler.GetHotel))
	r.HandleFunc("GET /hotels/rooms/{id}", token.JWTMiddleware(handler.GetRooms))
	r.HandleFunc("GET /hotels/room", token.JWTMiddleware(handler.GetRoom))
	r.HandleFunc("PUT /hotels/{id}", token.JWTMiddleware(handler.UpdateHotel))
	r.HandleFunc("PUT /hotels/rooms/{id}", token.JWTMiddleware(handler.UpdateRoom))
	r.HandleFunc("DELETE /hotels/{id}", token.JWTMiddleware(handler.DeleteHotel))
	r.HandleFunc("DELETE /hotels/rooms/{id}", token.JWTMiddleware(handler.DeleteRoom))

	//Booking

	r.HandleFunc("POST /bookings", token.JWTMiddleware(handler.CreateBooking))
	r.HandleFunc("POST /waitinglists", token.JWTMiddleware(handler.CreateWaiting))
	r.HandleFunc("GET /bookings/{id}", token.JWTMiddleware(handler.GetBooking))
	r.HandleFunc("GET /waitinglists/{id}", token.JWTMiddleware(handler.GetWaiting))
	r.HandleFunc("PUT /bookings/{id}", token.JWTMiddleware(handler.UpdateBooking))
	r.HandleFunc("PUT /waitinglists/{id}", token.JWTMiddleware(handler.UpdateWaiting))
	r.HandleFunc("DELETE /bookings/{id}", token.JWTMiddleware(handler.DeleteBooking))
	r.HandleFunc("DELETE /waitinglists/{id}", token.JWTMiddleware(handler.DeleteWaiting))

	certfile := "./cert/api.pem"
	keyfile := "./cert/api-key.pem"
	fmt.Printf("Server started on port %s\n", c.User.Port)
	if err := http.ListenAndServeTLS(c.User.Port, certfile, keyfile, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
