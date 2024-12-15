package model

import "time"

type Login struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}
type LoginReg struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	DriverId  *string    `db:"driver_id" json:"driverId"`
	BuyerId   *string    `db:"buyer_id"json:"buyerId"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}
type UserDriver struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	DriverId  string     `db:"driver_id" json:"driverId"`
}

type UserBuyer struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	BuyerId   string     `db:"buyer_id" json:"buyerId"`
}
type User struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	DriverId  string     `db:"driver_id" json:"driverId"`
	BuyerId   string     `db:"buyer_id" json:"buyerId"`
}

type Driver struct {
	Id                string     `json:"id"`
	UserId            string     `db:"user_id" json:"userId"`
	FirstName         string     `db:"first_name" json:"firstName"`
	LastName          string     `db:"last_name" json:"lastName"`
	PhoneNumber       string     `db:"phone_number" json:"phoneNumber"`
	PolicyNumber      string     `db:"policy_number" json:"policyNumber"`
	InsProvider       string     `db:"ins_provider" json:"insProvider"`
	PolicyStartDate   time.Time  `db:"policy_start_date" json:"policyStartDate"`
	PolicyEndDate     time.Time  `db:"policy_end_date" json:"policyEndDate"`
	LicenseNumber     string     `db:"license_number" json:"licenseNumber"`
	LicenseExpireDate time.Time  `db:"license_expire_date" json:"licenseExpireDate"`
	CreatedAt         time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt         *time.Time `db:"updated_at" json:"updatedAt"`
}

// first_name VARCHAR(255) NOT NULL,
//     last_name VARCHAR(255) NOT NULL,
//     phone_number VARCHAR(20) NOT NULL,
//     policy_number VARCHAR(255) NOT NULL,
//     ins_provider VARCHAR(255) NOT NULL,
//     policy_start_date TIMESTAMP WITH TIME ZONE NOT NULL,
//     policy_end_date TIMESTAMP WITH TIME ZONE NOT NULL,
//     license_number VARCHAR(255) NOT NULL,
//     license_expire_date TIMESTAMP WITH TIME ZONE NOT NULL,
//     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP WITH TIME ZONE
type Buyer struct {
	Id          string `json:"id"`
	UserId      string `db:"user_id" json:"userId"`
	FirstName   string `db:"first_name" json:"firstName"`
	LastName    string `db:"last_name" json:"lastName"`
	PhoneNumber string `db:"phone_number" json:"phoneNumber"`
	// TransactionID   string     `db:"" json:"ratingId"`
	PriorDeliveries int64      `db:"prior_deliveries" json:"priorDeliveries"`
	CreatedAt       time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updatedAt"`
}

type Rating struct {
	Id             string     `json:"id"`
	DriverId       string     `db:"driver_id" json:"driverId"`
	PastDeliveries string     `db:"past_deliveries" json:"PastDeliveries"`
	AverageRating  string     `db:"average_rating" json:"averageRating"`
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
