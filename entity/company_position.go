package entity

// CompanyPosition struct to refer to different positions inside a company
type CompanyPosition struct {
	PositionID uint64 `json:"position_id" db:"id"`
	Position   string `json:"position" db:"position"`
	*Company   `json:"company" db:"company_id"`
	*User      `json:"user" db:"user_id"`
}
