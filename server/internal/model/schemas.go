package model

import "time"

type User struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	DriverId  string     `json:"driverId"`
	BuyerId   string     `json:"buyerId"`
}

type Driver struct {
	Id                string     `json:"id"`
	UserId            string     `json:"userId"`
	FirstName         string     `json:"firstName"`
	LastName          string     `json:"lastName"`
	PhoneNumber       string     `json:"phoneNumber"`
	PolicyNumber      string     `json:"policyNumber"`
	InsProvider       string     `json:"insProvider"`
	PolicyStartDate   time.Time  `db:"policy_start_date" json:"policyStartDate"`
	PolicyEndDate     time.Time  `db:"policy_end_date" json:"policyEndDate"`
	LicenseNumber     string     `json:"licenseNumber"`
	LicenseExpireDate time.Time  `db:"license_expire_date" json:"licenseExpireDate"`
	CreatedAt         time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt         *time.Time `db:"updated_at" json:"updatedAt"`
}

type Buyer struct {
	Id              string     `json:"id"`
	UserId          string     `json:"userId"`
	FirstName       string     `json:"firstName"`
	LastName        string     `json:"lastName"`
	PhoneNumber     string     `json:"phoneNumber"`
	TransactionID   string     `json:"ratingId"`
	PriorDeliveries int64      `json:"priorDeliveries"`
	CreatedAt       time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updatedAt"`
}

type Rating struct {
	Id             string     `json:"id"`
	DriverId       string     `json:"driverId"`
	PastDeliveries string     `json:"PastDeliveries"`
	AverageRating  string     `json:"averageRating"`
	CreatedAt      time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updatedAt"`
}

// Todo: remove demo structs
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
