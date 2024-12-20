package model

import "time"

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User and LoginInput structs
type LoginInput struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type LoginRegInput struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	DriverId  *string    `db:"driver_id" json:"driverId"`
	OwnerId   *string    `db:"owner_id" json:"ownerId"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type UserDriverInput struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	DriverId  string     `db:"driver_id" json:"driverId"`
}

type UserOwnerInput struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	OwnerId   string     `db:"owner_id" json:"ownerId"`
}

type UserInput struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	DriverId  string     `db:"driver_id" json:"driverId"`
	OwnerId   string     `db:"owner_id" json:"ownerId"`
}

// Account associated structs
type ContactInfoInput struct {
	Id            string     `json:"id"`
	DriverId      string     `db:"driver_id" json:"driverId"`
	PhoneNumber   string     `db:"phone_number" json:"phoneNumber"`
	StreetAddress string     `db:"street_address" json:"streetAddress"`
	City          string     `db:"city" json:"city"`
	State         string     `db:"state" json:"state"`
	ZipCode       string     `db:"zip_code" json:"zipCode"`
	Country       string     `db:"country" json:"country"`
	CreatedAt     time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updated_at" json:"updatedAt"`
}

// Driver & Driver Associated structs
type DriverInput struct {
	Id        string     `json:"id"`
	UserId    string     `db:"user_id" json:"userId"`
	FirstName string     `db:"first_name" json:"firstName"`
	LastName  string     `db:"last_name" json:"lastName"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type InsuranceInput struct {
	Id              string     `json:"id"`
	DriverId        string     `db:"driver_id" json:"driverId"`
	PolicyNumber    string     `db:"policy_number" json:"policyNumber"`
	InsProvider     string     `db:"ins_provider" json:"insProvider"`
	PolicyStartDate string     `db:"policy_start_date" json:"policyStartDate"`
	PolicyEndDate   string     `db:"policy_end_date" json:"policyEndDate"`
	CreatedAt       time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updatedAt"`
}

type LicenseInput struct {
	Id                string     `json:"id"`
	DriverId          string     `db:"driver_id" json:"driverId"`
	LicenseNumber     string     `db:"license_number" json:"licenseNumber"`
	LicenseExpireDate time.Time  `db:"license_expire_date" json:"licenseExpireDate"`
	CreatedAt         time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt         *time.Time `db:"updated_at" json:"updatedAt"`
}

type HaulerInput struct {
	Id             string     `json:"id"`
	DriverId       string     `db:"driver_id" json:"driverId"`
	Make           string     `json:"make"`
	Model          string     `json:"model"`
	Year           int        `json:"year"`
	Mileage        float64    `json:"mileage"`
	TowingCapacity float64    `json:"towingCapacity"`
	CreatedAt      time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updatedAt"`
}

type TrailerInput struct {
	Id        string     `json:"id"`
	HaulerIds []string   `json:"haulerIds"`
	Type      string     `json:"type"`
	Length    float64    `json:"length"`
	Width     float64    `json:"width"`
	Capacity  float64    `json:"capacity"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type OfferInput struct {
	ID               string     `json:"id"`
	DriverId         string     `db:"driver_id" json:"driverId"`
	TransportationId string     `db:"transportation_id" json:"transportationId"`
	Amount           float64    `json:"amount"`
	DeadlineDate     time.Time  `db:"deadline_date" json:"deadlineDate"`
	HaulerId         string     `db:"hauler_id" json:"haulerId"`
	TrailerId        string     `db:"trailer_id" json:"trailerId"`
	CreatedAt        time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updated_at" json:"updatedAt"`
}

// Owner & Owner Associated structs
type OwnerInput struct {
	Id            string     `json:"id"`
	UserId        string     `db:"user_id" json:"userId"`
	FirstName     string     `db:"first_name" json:"firstName"`
	LastName      string     `db:"last_name" json:"lastName"`
	ContactInfoId string     `db:"contact_info_id" json:"contactInfoId"`
	CreatedAt     time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updated_at" json:"updatedAt"`
}

// Transportation & Transportation structs
type TransportationInput struct {
	Description         string  `json:"description" db:"description"`
	TransportDate       string  `json:"transportDate" db:"transport_date"`
	PickupAddress       string  `json:"pickupAddress" db:"pickup_address"`
	DeliveryAddress     string  `json:"deliveryAddress" db:"delivery_address"`
	DeliverByDate       string  `json:"deliverByDate" db:"deliver_by_date"`
	PickupByDate        string  `json:"pickupByDate" db:"pickup_by_date"`
	PickupAvailableDate string  `json:"pickupAvailableDate" db:"pickup_available_date"`
	RequestPrice        float64 `json:"requestPrice" db:"request_price"`
	VehicleId           string  `json:"vehicleId" db:"vehicle_id"`
}

type RatingInput struct {
	Id               string     `json:"id"`
	DriverId         string     `db:"driver_id" json:"driverId"`
	TransportationId string     `db:"transportation_id" json:"transportationId"`
	OwnerId          string     `db:"owner_id" json:"ownerId"`
	PastDeliveries   string     `db:"past_deliveries" json:"PastDeliveries"`
	AverageRating    string     `db:"average_rating" json:"averageRating"`
	CreatedAt        time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updated_at" json:"updatedAt"`
}

type TransactionInput struct {
	Id               string     `json:"id"`
	TransportationId string     `db:"transportation_id" json:"transportationId"`
	DriverId         string     `db:"driver_id" json:"driverId"`
	BuyerId          string     `db:"buyer_id" json:"buyerId"`
	PaymentMethod    string     `db:"payment_id" json:"paymentMethod"`
	Amount           float64    `db:"amount" json:"amount"`
	TransactionDate  time.Time  `db:"transaction_date" json:"transactionDate"`
	CreatedAt        time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updated_at" json:"updatedAt"`
}

type VehicleInput struct {
	Id               string     `json:"id"`
	TransportationId string     `db:"transportation_id" json:"transportationID"`
	AutoId           string     `db:"auto_id" json:"autoID"`
	BoatId           string     `db:"boat_id" json:"boatID"`
	Length           int        `json:"length"`
	Width            int        `json:"width"`
	Height           int        `json:"height"`
	CreatedAt        time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updated_at" json:"updatedAt"`
}

type AutoInput struct {
	Id        string     `json:"id"`
	VehicleId string     `db:"vehicle_id" json:"vehicleID"`
	Make      string     `json:"make"`
	Model     string     `json:"model"`
	Year      int        `json:"year"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type BoatInput struct {
	Id          string     `json:"id"`
	VehicleId   string     `db:"vehicle_id" json:"vehicleID"`
	Make        string     `json:"make"`
	Model       string     `json:"model"`
	Year        int        `json:"year"`
	WithTrailer bool       `json:"withTrailer"`
	CreatedAt   time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updatedAt"`
}

// Conversation represents a conversation between two parties
type ConversationInput struct {
	Id          string    `json:"id"`
	SenderId    string    `json:"sender_id"`
	RecipientId string    `json:"recipient_id"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
	Messages    []Message `json:"messages"` // Points to List of Messages
}

// Message represents individual messages within a conversation
type MessageInput struct {
	Id           string   `json:"id"`
	Participants []string `json:"participants"`
	Subject      string   `json:"subject"`
}

// Add this new struct to inputs.go
type DriverRegistrationInput struct {
	Driver    DriverInput     `json:"driver"`
	Insurance *InsuranceInput `json:"insurance,omitempty"`
	License   *LicenseInput   `json:"license,omitempty"`
}
