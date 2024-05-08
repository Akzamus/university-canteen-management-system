package entity

type Product struct {
	Uuid           string `db:"id"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	Price          int    `db:"price"`
	Image_url      string `db:"image_url"`
	Category_id    string `db:"category_id"`
	Total_quantity int    `db:"total_quantity"`
}
