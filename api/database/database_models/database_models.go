package database_models

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Name     string `yaml:"dbname"`
}

func (Auth) TableName() string {
	return "auth.users"
}

type Auth struct {
	ID       int `gorm:"primary_key"`
	Username string
	Password string
}

type Companies struct {
	ID            int `gorm:"primary_key"`
	Name          string
	StreetAddress string
	City          string
	State         string
	Zipcode       string
}

func (Companies) TableName() string {
	return "deals.companies"
}

type Discounts struct {
	ID          string `gorm:"primary_key"`
	CompanyName string
	Day         string
	Discount    string
}

func (Discounts) TableName() string {
	return "deals.discounts"
}
