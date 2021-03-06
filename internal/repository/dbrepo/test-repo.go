package dbrepo

import (
	"github.com/gdguesser/gobookings/internal/models"
	"time"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservations a reservation into the database
func (m *testDBRepo) InsertReservations(res models.Reservation) (int, error) {

	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into the database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if the availability exists for roomID, and false if no availability
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of availability rooms, if any, for a given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets a room by ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	return room, nil
}

// GetUserByID gets a user in the database by ID
func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

// UpdateUser updates a user in the database
func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

// Authenticate authenticates a user
func (m *testDBRepo) Authenticate(email, tespPassword string) (int, string, error) {
	return 1, "", nil
}

// AllReservations return a slice of all reservations
func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

// AllNewReservations return a slice of all reservations
func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

// GetReservationByID returns one reservation by ID
func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation
	return res, nil
}

// UpdateReservation updates a user in the database
func (m *testDBRepo) UpdateReservation(u models.Reservation) error {
	return nil
}

// DeleteReservation deletes one reservation by ID
func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

// UpdateProcessedForReservation updates processed for a reservation by ID
func (m *testDBRepo) UpdateProcessedForReservation(id, processed int) error {
	return nil
}