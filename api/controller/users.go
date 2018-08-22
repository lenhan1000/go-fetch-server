package controller

import (
	. "go-fetch-server/model"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION = "users"
)

//Find list of users
func FindAll(db *mgo.Database) ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

//Find user by ID
func FindByID(db *mgo.Database, id string) (User, error) {
	var user User
	err := db.C(COLLECTION).Find(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

//Insert new user
func Insert(db *mgo.Database, user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

//Delete a user
func Delete(db *mgo.Database, user User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

//Update a user
func Update(db *mgo.Database, user User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}

//
func FindAllUsers(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	users, err := FindAll(db)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

func CreateUser(db *mgo.Database, w http.ResponseWriter, r *http.Request) {
	user := User{}
	if err := user.CreateFromReader(r.Body); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	r.Body.Close()
	err := Insert(db, user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithMessage(w, http.StatusOK, "Successful")
}
