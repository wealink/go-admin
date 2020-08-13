package models

type Base struct {
	Id          int  `json:"id" gorm:"primary_key"`
	Created_On  int  `json:"created_on"`
	Modified_On int  `json:"modified_on"`
	Deleted_On  *int `json:"deleted_on" sql:"index"`
}
