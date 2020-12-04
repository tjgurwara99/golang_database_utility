package entity

// Category Struct to represent Category of products of a Company
type Category struct {
	CategoryID int64
	Name       string
	Visible    bool
	Active     bool
	Code       string
	*Company
}
