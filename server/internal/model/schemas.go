package model

import "time"

// User and Login structs
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
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	DriverId  *string    `db:"driver_id" json:"driverId"`
	OwnerId   *string    `db:"owner_id" json:"ownerId"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}
type UserDriver struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	DriverId  string     `db:"driver_id" json:"driverId"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}
type UserOwner struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	OwnerId   string     `db:"owner_id" json:"ownerId"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}
type User struct {
	Id        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	DriverId  string     `db:"driver_id" json:"driverId"`
	OwnerId   string     `db:"owner_id" json:"ownerId"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

// Account associated Structs
type ContactInfo struct {
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

// Driver & Driver Associated Structs
type Driver struct {
	Id              string           `json:"id" db:"id"`
	UserId          string           `db:"user_id" json:"userId"`
	FirstName       string           `db:"first_name" json:"firstName"`
	LastName        string           `db:"last_name" json:"lastName"`
	ContactInfo     *ContactInfo     `db:"-" json:"contactInfo,omitempty"`
	Insurance       *Insurance       `db:"-" json:"insurance,omitempty"`
	License         *License         `db:"-" json:"license,omitempty"`
	Rating          *Rating          `db:"-" json:"rating,omitempty"`
	Haulers         []Hauler         `db:"-" json:"haulers,omitempty"`
	Offers          []Offer          `db:"-" json:"offers,omitempty"`
	Transportations []Transportation `db:"-" json:"transportations,omitempty"`
	CreatedAt       time.Time        `db:"created_at" json:"createdAt"`
	UpdatedAt       *time.Time       `db:"updated_at" json:"updatedAt"`
}

type Insurance struct {
	Id              string     `json:"id"`
	DriverId        string     `db:"driver_id" json:"driverId"`
	PolicyNumber    string     `db:"policy_number" json:"policyNumber"`
	InsProvider     string     `db:"ins_provider" json:"insProvider"`
	PolicyStartDate time.Time  `db:"policy_start_date" json:"policyStartDate"`
	PolicyEndDate   time.Time  `db:"policy_end_date" json:"policyEndDate"`
	CreatedAt       time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updatedAt"`
}

type License struct {
	Id                string     `json:"id"`
	DriverId          string     `db:"driver_id" json:"driverId"`
	LicenseNumber     string     `db:"license_number" json:"licenseNumber"`
	LicenseExpireDate time.Time  `db:"license_expire_date" json:"licenseExpireDate"`
	CreatedAt         time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt         *time.Time `db:"updated_at" json:"updatedAt"`
}
type Hauler struct {
	Id             string     `json:"id" db:"id"`
	DriverId       string     `db:"driver_id" json:"driverId"`
	Trailers       []Trailer  `json:"trailers" db:"-"`
	Make           string     `json:"make" db:"make"`
	Model          string     `json:"model" db:"model"`
	Year           int        `json:"year" db:"year"`
	Mileage        float64    `json:"mileage" db:"mileage"`
	TowingCapacity float64    `db:"towing_capacity" json:"towingCapacity"`
	CreatedAt      time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updatedAt"`
}

type Trailer struct {
	Id        string     `json:"id" db:"id"`
	HaulerIds []string   `json:"haulerIds" db:"-"`
	Type      string     `json:"type" db:"type"`
	Length    float64    `json:"length" db:"length"`
	Width     float64    `json:"width" db:"width"`
	Capacity  float64    `json:"capacity" db:"capacity"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type Offer struct {
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

// Owner & Owner Associated Structs
type Owner struct {
	Id            string     `json:"id"`
	UserId        string     `db:"user_id" json:"userId"`
	FirstName     string     `db:"first_name" json:"firstName"`
	LastName      string     `db:"last_name" json:"lastName"`
	ContactInfoId string     `db:"contact_info_id" json:"contactInfoId"`
	CreatedAt     time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt     *time.Time `db:"updated_at" json:"updatedAt"`
}

// Transportation & Transportation Structs
type Transportation struct {
	Id                  string     `json:"id" db:"id"`
	DriverId            string     `json:"driverId" db:"driver_id"`
	Description         string     `json:"description" db:"description"`
	TransportDate       time.Time  `json:"transportDate" db:"transport_date"`
	PickupAddress       string     `json:"pickupAddress" db:"pickup_address"`
	DeliveryAddress     string     `json:"deliveryAddress" db:"delivery_address"`
	DeliverByDate       time.Time  `json:"deliverByDate" db:"deliver_by_date"`
	PickupByDate        time.Time  `json:"pickupByDate" db:"pickup_by_date"`
	PickupAvailableDate time.Time  `json:"pickupAvailableDate" db:"pickup_available_date"`
	AcceptedOfferId     string     `json:"acceptedOfferId" db:"accepted_offer_id"`
	VehicleId           string     `json:"vehicleId" db:"vehicle_id"`
	RequestPrice        float64    `json:"requestPrice" db:"request_price"`
	Offers              []Offer    `json:"participants" db:"-"`
	CreatedAt           time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt           *time.Time `json:"updatedAt" db:"updated_at"`
}

type Rating struct {
	Id               string     `json:"id"`
	DriverId         string     `db:"driver_id" json:"driverId"`
	TransportationId string     `db:"transportation_id" json:"transportationId"`
	OwnerId          string     `db:"owner_id" json:"ownerId"`
	PastDeliveries   string     `db:"past_deliveries" json:"PastDeliveries"`
	AverageRating    string     `db:"average_rating" json:"averageRating"`
	CreatedAt        time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updated_at" json:"updatedAt"`
}

type Transaction struct {
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

type Vehicle struct {
	Id               string     `json:"id" db:"id"`
	TransportationId string     `json:"transportationId" db:"transportation_id"`
	Length           int        `json:"length" db:"length"`
	Width            int        `json:"width" db:"width"`
	Height           int        `json:"height" db:"height"`
	AutoId           string     `json:"autoId" db:"auto_id"`
	BoatId           string     `json:"boatId" db:"boat_id"`
	CreatedAt        time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt        *time.Time `db:"updated_at" json:"updatedAt"`
}

type Auto struct {
	Id        string     `json:"id"`
	VehicleId string     `db:"vehicle_id" json:"vehicleID"`
	Make      string     `json:"make"`
	Model     string     `json:"model"`
	Year      int        `json:"year"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
}

type Boat struct {
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
type Conversation struct {
	Id          string    `json:"id"`
	SenderId    string    `json:"sender_id"`
	RecipientId string    `json:"recipient_id"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
	Messages    []Message `json:"messages"` // Points to List of Messages
}

// Message represents individual messages within a conversation
type Message struct {
	Id           string   `json:"id"`
	Participants []string `json:"participants"`
	Subject      string   `json:"subject"`
}
