package handler_models

type CompanyParams struct {
	ID            *string `form:"id"`
	Name          *string `form:"name"`
	StreetAddress *string `form:"streetaddress"`
	City          *string `form:"city"`
	Zipcode       *string `form:"zipcode"`
}

type Authentication struct {
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required"`
}
