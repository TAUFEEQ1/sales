package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

// Order is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Order struct {
	Region        string    `json:"region" db:"region"`
	Country       string    `json:"country" db:"country"`
	ItemType      string    `json:"itemType" db:"itemType"`
	SalesChannel  string    `json:"salesChannel" db:"salesChannel"`
	OrderPriority string    `json:"orderPriority" db:"orderPriority"`
	OrderDate     time.Time `json:"orderDate" db:"orderDate"`
	ID            int       `json:"id" db:"id"`
	OrderId       int       `json:"orderId" db:"orderId"`
	ShipDate      time.Time `json:"shipDate" db:"shipDate"`
	UnitsSold     int       `json:"unitsSold" db:"unitsSold"`
	UnitPrice     float32   `json:"unitPrice" db:"unitPrice"`
	UnitCost      float32   `json:"unitCost" db:"unitCost"`
	TotalRevenue  float32   `json:"totalRevenue" db:"totalRevenue"`
	TotalCost     float32   `json:"totalCost" db:"totalCost"`
	TotalProfit   float32   `json:"totalProfit" db:"totalProfit"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (o Order) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Orders is not required by pop and may be deleted
type Orders []Order

// String is not required by pop and may be deleted
func (o Orders) String() string {
	jo, _ := json.Marshal(o)
	return string(jo)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (o *Order) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (o *Order) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (o *Order) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
