package dto

type ProductRequestDto struct {
	Name    string `json:"name"`
	Description string `json:"description"`
	Image_url     string `json:"image_url"`
}

type ProductResponseDto struct {
	Uuid  string `json:"uuid"`
	Name string `json:"name"`
	Price  int `json:"price"`
}
