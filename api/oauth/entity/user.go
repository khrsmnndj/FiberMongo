package entity

// User struct
type User struct {
	ID     		string  `json:"id,omitempty" bson:"_id,omitempty"`
	Fullname   	string  `json:"fullname"`
	Email 		string 	`json:"email"`
	Phone 		string 	`json:"phone"`
}