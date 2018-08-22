package model

import (
	"encoding/json"
	"io"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"_id"`
	Local     Local         `bson:"local" json:"local"`
	Info      Info          `bson:"info" json:"info"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt" json:"updatedAt"`
}

type Local struct {
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type Info struct {
	DisplayName string `bson:"displayName" json:"displayName"`
	Address     string `bson:"address" json:"address"`
	Country     string `bson:"country" json:"country"`
	State       string `bson:"state" json:"state"`
	City        string `bson:"city" json:"city"`
	ZipCode     string `bson:"zipCode" json:"zipCode"`
	CountryCode string `bson:"countryCode" json:"countryCode"`
	MobilePhone string `bson:"mobilePhone" json:"mobilePhone"`
	Carrier     string `bson:"carrier" json:"carrier"`
}

func (u *User) CreateFromReader(r io.Reader) error {
	decode := json.NewDecoder(r)
	if err := decode.Decode(u); err != nil {
		return err
	}
	u.ID = bson.NewObjectId()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
