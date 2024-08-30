package handler

import (
	broad "api-gateway/internal/broadcast"
	"api-gateway/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	B *broad.Adjust
}

// Register handles the user registration process.
// @Summary Register a new user
// @Description Register a new user by providing email and password. A verification code will be sent to the provided email.
// @Tags user
// @Accept json
// @Produce json
// @Param registerUserRequest body models.RegisterUserRequest true "User registration data"
// @Success 200 {string} string "Verification code is sent to your email"
// @Failure 500 {string} string "Internal Server Error"
// @Failure 403 {string} string "Already Exists"
// @Router /users/register [post]
func (u *Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.Register(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Verification code is sent to your email")
}

// Verify handles the user verification process.
// @Summary Verify a user account
// @Description Verify a user account by providing the verification code sent to the user's email.
// @Tags user
// @Accept json
// @Produce json
// @Param verifyRequest body models.VerifyRequest true "Verification code data"
// @Success 200 {string} string "You have verified your account and now you can log in"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/verify [post]
func (u *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.VerifyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := u.B.Verify(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("You have verified your account and now you can log in")
}

// LogIn handles the user login process.
// @Summary Log in a user
// @Description Log in a user by providing their email and password.
// @Tags user
// @Accept json
// @Produce json
// @Param logInRequest body models.LogInRequest true "User login data"
// @Success 200 {string} string "JWT token"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/login [post]
func (u *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.LogInRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := u.B.Login(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(token)
}

// GetUser retrieves user information by ID.
// @Summary Get user information
// @Description Retrieve user information by providing the user ID.
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.GetUserResponse "User information"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /users/{id} [get]
func (u *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetUser(&models.GetUserRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateUser updates user information by ID.
// @Summary Update user information
// @Description Update user information by providing the user ID in the path and the fields to update in the request body.
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param updateUserRequest body models.UpdateUserRequest true "User update data"
// @Success 200 {string} string "Your account is updating we'll notify you when it's updated"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /users/{id} [put]
func (u *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req models.UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.ID = int32(id)
	if err := u.B.UpdateUser(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Your account is updating we'll notify you when it's updated")
}

// DeleteUser deletes a user by ID.
// @Summary Delete a user
// @Description Delete a user by providing their user ID.
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "Your account is deleting"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /users/{id} [delete]
func (u *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.DeleteUser(&models.GetUserRequest{ID: int32(id)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Your account is deleting")
}

// LogOut logs out a user by ID.
// @Summary Log out a user
// @Description Log out a user by providing their user ID.
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string "You have logged out!"
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /users/logout/{id} [post]
func (u *Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.Logout(&models.GetUserRequest{ID: int32(id)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("You have logged out!")
}

// Hotel---Service

// CreateHotel godoc
// @Summary      Create a new hotel
// @Description  Create a new hotel with the given details
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        hotel  body      models.CreateHotelRequest  true  "Hotel details"
// @Success      200    {string}  string                     "Hotel created successfully"
// @Failure      500    {string}  string                     "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels/create [post]
func (u *Handler) CreateHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateHotelRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.CreateHotel(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Hotel created successfully")
}

// GetHotel godoc
// @Summary      Get hotel details
// @Description  Get details of a specific hotel by ID
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        id    path      int                     true  "Hotel ID"
// @Success      200   {object}  models.GetHotelResponse
// @Failure      500   {string}  string                  "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels/{id} [get]
func (u *Handler) GetHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetHotel(&models.GetHotelRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// GetHotels godoc
// @Summary      Get list of hotels
// @Description  Get a list of all available hotels
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Success      200   {array}   models.GetHotelResponse
// @Failure      500   {string}  string                  "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels [get]
func (u *Handler) GetHotels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := u.B.GetHotels(&models.GetsRequest{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateHotel godoc
// @Summary      Update hotel details
// @Description  Update the details of a specific hotel by ID
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        id     path      int                        true  "Hotel ID"
// @Param        hotel  body      models.UpdateHotelRequest  true  "Updated hotel details"
// @Success      200    {string}  string                     "Hotel details are updated"
// @Failure      500    {string}  string                     "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels/{id} [put]
func (u *Handler) UpdateHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req models.UpdateHotelRequest
	req.ID = int32(id)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.UpdateHotel(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Hotel deails are updated")
}

// DeleteHotel godoc
// @Summary      Delete a hotel
// @Description  Delete a specific hotel by ID
// @Tags         hotels
// @Accept       json
// @Produce      json
// @Param        id    path      int                     true  "Hotel ID"
// @Success      200   {string}  string                  "Hotel Deleted"
// @Failure      500   {string}  string                  "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels/{id} [delete]
func (u *Handler) DeleteHotel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.DeleteHotel(&models.GetHotelRequest{ID: int32(id)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Hotel Deleted")
}

// CreateRoom godoc
// @Summary      Create a new room
// @Description  Create a new room in a specific hotel
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        room   body      models.CreateRoomRequest  true  "Room details"
// @Success      200    {string}  string                    "Room created in hotel"
// @Failure      500    {string}  string                    "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels/rooms/create [post]
func (u *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateRoomRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.CreateRoom(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("room created in hotel")
}

// GetRoom godoc
// @Summary      Get room details
// @Description  Get details of a specific room by hotel and room IDs
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        hotel  query     int                     true  "Hotel ID"
// @Param        room   query     int                     true  "Room ID"
// @Success      200    {object}  models.GetRoomResponse
// @Failure      500    {string}  string                  "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels/room [get]
func (u *Handler) GetRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomid, err := strconv.Atoi(r.URL.Query().Get("room"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hotelid, err := strconv.Atoi(r.URL.Query().Get("hotel"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := u.B.GetRoom(&models.GetRoomRequest{HotelID: int32(hotelid), ID: int32(roomid)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// GetRooms godoc
// @Summary      Get list of rooms
// @Description  Get a list of all rooms in a specific hotel
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        id    path      int                     true  "Hotel ID"
// @Success      200   {array}   models.GetRoomResponse
// @Failure      500   {string}  string                  "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels/rooms/{id} [get]
func (u *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetRooms(&models.GetRoomRequest{HotelID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateRoom godoc
// @Summary      Update room details
// @Description  Update the details of a specific room
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        room   body      models.UpdateRoomRequest  true  "Updated room details"
// @Success      200    {string}  string                    "Room details are updated"
// @Failure      500    {string}  string                    "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels/rooms/{id} [put]
func (u *Handler) UpdateRoom(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := u.B.UpdateRoom(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Room details are updated")
}

// DeleteRoom godoc
// @Summary      Delete a room
// @Description  Delete a specific room by hotel and room IDs
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        hotel  query     int                     true  "Hotel ID"
// @Param        room   query     int                     true  "Room ID"
// @Success      200    {string}  string                  "Room is deleted"
// @Failure      500    {string}  string                  "Internal Server Error"
// @Security     BearerAuth
// @Router       /hotels/rooms/{id} [delete]
func (u *Handler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomid, err := strconv.Atoi(r.URL.Query().Get("room"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hotelid, err := strconv.Atoi(r.URL.Query().Get("hotel"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := u.B.DeleteRoom(&models.GetRoomRequest{HotelID: int32(hotelid), ID: int32(roomid)}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Room is deleted")
}

// CreateBooking creates a new hotel booking.
// @Summary Create a new hotel booking
// @Description Create a new hotel booking
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param request body models.BookHotelRequest true "Booking details"
// @Success 200 {object} models.GeneralResponse
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /bookings [post]
func (u *Handler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.BookHotelRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.CreateBooking(&req)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(res)
}

// GetBooking retrieves booking details by ID.
// @Summary Get booking details
// @Description Get booking details by ID
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param id path int true "Booking ID"
// @Success 200 {object} models.GetUsersBookResponse
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /bookings/{id} [get]
func (u *Handler) GetBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetBooking(&models.GetUsersBookRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateBooking updates an existing hotel booking.
// @Summary Update hotel booking
// @Description Update an existing hotel booking by ID
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param id path int true "Booking ID"
// @Param request body models.BookHotelUpdateRequest true "Updated booking details"
// @Success 200 {object} models.GeneralResponse
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /bookings/{id} [put]
func (u *Handler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var req models.BookHotelUpdateRequest
	req.ID = int32(id)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.UpdateBooking(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// DeleteBooking deletes a hotel booking by ID.
// @Summary Delete hotel booking
// @Description Delete a hotel booking by ID
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param id path int true "Booking ID"
// @Success 200 {object} models.GeneralResponse
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /bookings/{id} [delete]
func (u *Handler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.DeleteBooking(&models.CancelRoomRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// CreateWaiting adds a new entry to the waiting list.
// @Summary Add to waiting list
// @Description Add a new entry to the waiting list
// @Tags waitinglists
// @Accept  json
// @Produce  json
// @Param request body models.CreateWaitingList true "Waiting list details"
// @Success 200 {object} models.GeneralResponse
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /waitinglists [post]
func (u *Handler) CreateWaiting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateWaitingList

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.CreateWaitinglist(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)

}

// GetWaiting retrieves waiting list details by ID.
// @Summary Get waiting list details
// @Description Get waiting list details by ID
// @Tags waitinglists
// @Accept  json
// @Produce  json
// @Param id path int true "Waiting List ID"
// @Success 200 {object} models.GetWaitinglistResponse
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /waitinglists/{id} [get]
func (u *Handler) GetWaiting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := u.B.GetWaiting(&models.GetWaitinglistRequest{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// UpdateWaiting updates an entry in the waiting list.
// @Summary Update waiting list
// @Description Update an existing entry in the waiting list by ID
// @Tags waitinglists
// @Accept  json
// @Produce  json
// @Param id path int true "Waiting List ID"
// @Param request body models.UpdateWaitingListRequest true "Updated waiting list details"
// @Success 200 {object} models.GeneralResponse
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /waitinglists/{id} [put]
func (u *Handler) UpdateWaiting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var req models.UpdateWaitingListRequest
	req.ID = int32(id)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := u.B.UpdateWaiting(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// DeleteWaiting deletes an entry from the waiting list by ID.
// @Summary Delete waiting list entry
// @Description Delete an entry from the waiting list by ID
// @Tags waitinglists
// @Accept  json
// @Produce  json
// @Param id path int true "Waiting List ID"
// @Success 200 {object} models.GeneralResponse
// @Failure 500 {string} string "Internal Server Error"
// @Security BearerAuth
// @Router /waitinglists/{id} [delete]
func (u *Handler) DeleteWaiting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	res, err := u.B.DeleteWaiting(&models.DeleteWaitingList{ID: int32(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}
