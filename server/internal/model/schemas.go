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
	Password  string     `json:"-"`
	DriverId  *string    `db:"driver_id" json:"driverId"`
	OwnerId   *string    `db:"owner_id" json:"ownerId"`
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
type UserOwner struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	OwnerId   string     `db:"owner_id" json:"ownerId"`
}
type User struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	DriverId  string     `db:"driver_id" json:"driverId"`
	OwnerId   string     `db:"owner_id" json:"ownerId"`
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
	Id              string           `json:"id"`
	UserId          string           `db:"user_id" json:"userId"`
	FirstName       string           `db:"first_name" json:"firstName"`
	LastName        string           `db:"last_name" json:"lastName"`
	ContactInfoId   string           `db:"contact_info_id" json:"contactInfoId"`
	InsuranceId     string           `db:"insurance_id" json:"insuranceId"`
	LicenseId       string           `db:"license_id" json:"licenseId"`
	RatingId        string           `db:"rating_id" json:"ratingId"`
	Haulers         []Hauler         `json:"Haulers"`
	Offers          []Offer          `json:"Offers"`
	Transportations []Transportation `json:"Transportations"`
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
	Id             string     `json:"id"`
	DriverId       string     `db:"driver_id" json:"driverId"`
	Trailers       []Trailer  `json:"trailers"`
	Make           string     `json:"make"`
	Model          string     `json:"model"`
	Year           int        `json:"year"`
	Mileage        float64    `json:"mileage"`
	TowingCapacity float64    `json:"towingCapacity"`
	CreatedAt      time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updatedAt"`
}

type Trailer struct {
	Id        string     `json:"id"`
	HaulerId  string     `json:"haulerId"`
	Type      string     `json:"type"`
	Length    float64    `json:"length"`
	Width     float64    `json:"width"`
	Capacity  float64    `json:"capacity"`
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
	Id                  string     `json:"id"`
	Description         string     `json:"description"`
	TransportDate       time.Time  `db:"transportation_date" json:"transportDate"`
	PickupAddress       string     `db:"pickup_address" json:"pickupAddress"`
	DeliveryAddress     string     `db:"delivery_address" json:"deliveryAddress"`
	DeliverByDate       time.Time  `db:"delivery_by_date" json:"deliverByDate"`
	PickupByDate        time.Time  `db:"pickup_by_date" json:"pickupByDate"`
	PickupAvailableDate time.Time  `db:"pickup_by_Avialable_date" json:"pickupAvailableDate"`
	AcceptedOfferId     string     `db:"accecpted_offer_id" json:"acceptedOfferId"`
	VehicleId           string     `db:"vechicle_id" json:"vehicleId"`
	RequestPrice        float64    `db:"request_price" json:"requestPrice"`
	Offers              []Offer    `json:"participants"`
	CreatedAt           time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt           *time.Time `db:"updated_at" json:"updatedAt"`
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
