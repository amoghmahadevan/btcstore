package templates

type menuItem struct {
	label string
	href  string
}

var menuItems = []menuItem{
	{"Home", "/"},
	{"Books", "/books"},
	{"Devices", "/devices"},
	{"Clothes", "/clothes"},
}
