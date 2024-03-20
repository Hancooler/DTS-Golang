package models

import (
	"time"

	"gorm.io/gorm"
)

type Borrowing struct {
	gorm.Model
	MemberID   uint       "json:member_id"
	PetugasID  uint       "json:petugas_id"
	BookID     uint       "json:book_id"
	BorrowDate time.Time  "json:borrow_date"
	ReturnDate *time.Time "json:return_date"
	Book       Book       "json:book"
	Member     Member     "json:member"
	Petugas    Petugas    "json:petugas"
}
