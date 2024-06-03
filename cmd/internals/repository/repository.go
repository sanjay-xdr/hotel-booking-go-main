package repository

import "github.com/sanjay-xdr/cmd/internals/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
