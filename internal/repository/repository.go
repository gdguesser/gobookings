package repository

import (
	"github.com/gdguesser/gobookings/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservations(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDate(start, end time.Time, roomID int) (bool, error)
}
