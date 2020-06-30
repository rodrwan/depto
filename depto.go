package depto

import "time"

// Item represents a depto item.
type Item struct {
	// Primary key of an item.
	ID string `json:"id,omitempty" db:"id"`
	// Item name.
	Name string `json:"name,omitempty" db:"name"`
	// Item description.
	Description string `json:"description,omitempty" db:"description"`

	CreatedAt time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`

	Info []*ItemInfo `json:"infos,omitempty"`
}

// ItemQuery is a helper struct to query database.
type ItemQuery struct {
	ID      string
	Name    string
	Deleted bool
}

// ItemInfo represents a more detailed item information.
type ItemInfo struct {
	ID     int64  `json:"id,omitempty" db:"id"`
	ItemID string `json:"item_id,omitempty" db:"item_id"`

	Brand string  `json:"brand,omitempty" db:"brand"`
	Price float64 `json:"price,omitempty" db:"price"`
	Unit  int64   `json:"unit,omitempty" db:"unit"`

	PurchaseDate   time.Time `json:"purchase_date,omitempty" db:"purchase_date"`
	PurchasePlace  string    `json:"purchase_place,omitempty" db:"purchase_place"`
	ExpirationDate time.Time `json:"expiration_date,omitempty" db:"expiration_date"`

	CreatedAt time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`

	Images []*ItemImage `json:"images,omitempty"`
}

// ItemInfoQuery is a helper struct to query database.
type ItemInfoQuery struct {
	ID             int64
	ItemID         string
	Brand          string
	Price          float64
	PurchaseDate   time.Time
	PurchasePlace  string
	ExpirationDate time.Time
	Deleted        bool
}

// ItemImage represent image from an item.
type ItemImage struct {
	ID         int64  `json:"id,omitempty" db:"id"`
	ItemInfoID string `json:"item_info_id,omitempty" db:"item_info_id"`

	ImageURL string `json:"image_url,omitempty" db:"image_url"`

	CreatedAt time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// ItemImageQuery is a helper struct to query database.
type ItemImageQuery struct {
	ID      int64
	ItemID  int64
	Deleted bool
}

// Inventory represent the stock of products in inventory.
// This make reference to inventory table.
type Inventory struct {
	ID   string `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`

	CreatedAt time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// InventoryQuery is a helper struct to query database.
type InventoryQuery struct {
	ID      string
	Deleted bool
}

// InventoryItem represent the stock of products in inventory.
// This make reference to inventory table.
type InventoryItem struct {
	ID          string `json:"id,omitempty" db:"id"`
	InventoryID string `json:"inventory_id,omitempty" db:"inventory_id"`
	ItemID      string `json:"item_id,omitempty" db:"item_id"`

	Count   int64      `json:"count,omitempty" db:"count"`
	LastUse *time.Time `json:"last_use,omitempty" db:"last_use"`

	CreatedAt time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

// InventoryItemQuery is a helper struct to query database.
type InventoryItemQuery struct {
	ID          string
	InventoryID string
	Deleted     bool
}
