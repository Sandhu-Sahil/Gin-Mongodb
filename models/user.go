package models

type Address struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	Pincode string `json:"pincode" bson:"pincode"`
}

type User struct {
	Name    string  `json:"name,omitempty" bson:"user_name,omitempty"`
	Age     int     `json:"age,string,omitempty" bson:"user_age,omitempty"`
	Address Address `json:"address" bson:"user_address"`
}
