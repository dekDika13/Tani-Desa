package adminDto

type ProductDTO struct {
	ProductID   uint   `json:"product_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Type        string `json:"type"`
	Qty         uint   `json:"qty"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Owner       string `json:"owner"`
	AdminName   string `json:"admin_name"`
}

type ProductRequest struct {
	AdminID     uint   `json:"admin_id"`
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Qty         uint   `json:"qty"`
	Price       uint   `json:"price" validate:"required"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Owner       string `json:"owner" validate:"required"`
}
