package template

var (
	// Frontend
	FrontendIndex = ParseBackEnd("themes/backend/index/index.html", "themes/backend/index/price.html")
	FrontendOther = ParseBackEnd("themes/backend/index/other.html")
)
