package models

import "time"

// Reservation holds Reservation data
type Reservation struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Users is the Users model
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Rooms is the Rooms model
type Rooms struct {
	ID        int
	RoomName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Restrictions is the Restrictions model
type Restrictions struct {
	ID              int
	RestrictionName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Reservations is the Reservations model
type Reservations struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	startDate time.Time
	endDate   time.Time
	RoomID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// RoomRestrictions is the RoomRestrictions model
type RoomRestrictions struct {
	ID            int
	RoomID        int
	ReservationID int
	RestrictionID int
	Room          Rooms
	Reservation   Reservations
	Restriction   Restrictions
	startDate     time.Time
	endDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
