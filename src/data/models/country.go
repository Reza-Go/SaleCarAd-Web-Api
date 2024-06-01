package models

type Country struct {
	BaseModel
	Name   string `gorm:"size:15;type:string;not null"`
	Cities []City
}

type City struct {
	BaseModel
	Name      string `gorm:"size:10;type:string;not null"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:No ACTION;OnDelete:No ACTION"`
}
