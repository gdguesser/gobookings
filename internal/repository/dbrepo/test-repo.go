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
