package model

import "time"

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DriverInput struct {
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

// Demo test data
type CategoryInput struct {
	Title string `json:"title"`
}

type TransactionInput struct {
	CategoryId string  `json:"categoryId"`
	Title      string  `json:"title"`
	Amount     float32 `json:"amount"`
	Currency   string  `json:"currency"`
	Type       string  `json:"type"`
}
