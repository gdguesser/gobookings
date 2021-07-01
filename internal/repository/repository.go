package repository

import "github.com/gdguesser/gobookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservations(res models.Reservation) error
}
