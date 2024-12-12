package model

import "time"

type User struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
}

type Account struct {
	Id          string     `json:"id"`
	UserId      string     `json:"userId"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	CreatedAt   time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
	PhoneNumber string     `json:"phoneNumber"`
}

type Driver struct {
	Account // Embed the User struct to inherit its fields

	YearsOfExperience int       `json:"yearsOfExperience"`
	PolicyNumber      string    `json:"policyNumber"`
	InsProvider       string    `json:"insProvider"`
	PolicyStartDate   time.Time `db:"policy_start_date" json:"policyStartDate"`
	PolicyEndDate     time.Time `db:"policy_end_date" json:"policyEndDate"`
	LicenseNumber     string    `json:"licenseNumber"`
	LicenseExpireDate time.Time `db:"license_expire_date" json:"licenseExpireDate"`
	RatingID          string    `json:"ratingId"`
}

type Buyer struct {
	Account // Embed the User struct to inherit its fields

	TransactionID   string `json:"ratingId"`
	PriorDeliveries int64  `json:"priorDeliveries"`
}

type Rating struct {
	UserId         string `json:"userId"`
	PastDeliveries string `json:"PastDeliveries"`
	AverageRating  string `json:"averageRating"`
}

type Category struct {
	Id        string     `json:"id"`
	UserId    string     `db:"user_id" json:"user_id"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type Transaction struct {
	Id         string     `json:"id"`
	UserId     string     `db:"user_id" json:"userId"`
	CategoryId string     `db:"category_id" json:"categoryId"`
	Title      string     `json:"title"`
	Amount     float32    `json:"amount"`
	Currency   string     `json:"currency"`
	Type       string     `json:"type"`
	CreatedAt  time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updatedAt"`
}
