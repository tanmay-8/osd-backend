package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Email       string             `json:"email" bson:"email"`
	Phone       string             `json:"phone" bson:"phone"`
	CollegeName string             `json:"collegeName" bson:"collegeName"`
	YearOfStudy string             `json:"yearOfStudy" bson:"yearOfStudy"`
}

type UserInput struct {
	Name        string `json:"name" bson:"name"`
	Email       string `json:"email" bson:"email"`
	Phone       string `json:"phone" bson:"phone"`
	CollegeName string `json:"collegeName" bson:"collegeName"`
	YearOfStudy string `json:"yearOfStudy" bson:"yearOfStudy"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type ReferralScore struct {
	ReferralCode string `json:"referralCode" bson:"referralCode"`
	Score        int    `json:"score" bson:"score"`
}
