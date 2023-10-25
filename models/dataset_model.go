package models

type DataSetModel []struct {
	ID             string   `json:"_id"`
	ActualPrice    string   `json:"actual_price"`
	AverageRating  string   `json:"average_rating"`
	Brand          string   `json:"brand"`
	Category       string   `json:"category"`
	CrawledAt      string   `json:"crawled_at"`
	Description    string   `json:"description"`
	Discount       string   `json:"discount"`
	Images         []string `json:"images"`
	OutOfStock     bool     `json:"out_of_stock"`
	Pid            string   `json:"pid"`
	ProductDetails []struct {
		StyleCode string `json:"Style Code,omitempty"`
		Closure   string `json:"Closure,omitempty"`
		Pockets   string `json:"Pockets,omitempty"`
		Fabric    string `json:"Fabric,omitempty"`
		Pattern   string `json:"Pattern,omitempty"`
		Color     string `json:"Color,omitempty"`
	} `json:"product_details"`
	Seller       string `json:"seller"`
	SellingPrice string `json:"selling_price"`
	SubCategory  string `json:"sub_category"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}
