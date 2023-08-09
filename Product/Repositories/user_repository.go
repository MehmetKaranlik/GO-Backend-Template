package Repositories

import (
	"Backend/Core/Utilities/Base"
	"Backend/Core/Utilities/Hasher"
	"Backend/Product/Init/Databases/Mongo"
	"Backend/Product/Model/Auth"
	"Backend/Product/Model/User"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IUserRepository interface {
	Base.IRepository[User.User]
	SecureAccess(body Auth.LoginBody) *User.User
	ReadByEntry(email string) *User.User
}

type UserRepository struct {
	Ref *Mongo.MongoCollectionRef
}

func (u UserRepository) Create(model *User.User) error {
	_, err := u.Ref.Ref.InsertOne(context.Background(), model, options.InsertOne())
	return err
}

func (u UserRepository) Read(id int) (*User.User, error) {
	filter := bson.M{"id": id}
	var result User.User
	err := u.Ref.Ref.FindOne(context.Background(), filter).Decode(&result)
	return &result, err
}

func (u UserRepository) Update(model *User.User) error {
	filter := bson.M{"id": model.Id}
	update := bson.M{"$set": model}
	_, err := u.Ref.Ref.UpdateOne(context.Background(), filter, update, options.Update())
	return err
}

func (u UserRepository) Delete(id int) error {
	filter := bson.M{"id": id}
	_, err := u.Ref.Ref.DeleteOne(context.Background(), filter, options.Delete())
	return err
}

func (u UserRepository) Aggregate(pipeline []interface{}) (*[]User.User, error) {
	cursor, err := u.Ref.Ref.Aggregate(context.Background(), pipeline, options.Aggregate())
	if err != nil {
		return nil, err
	}
	var result []User.User
	err = cursor.All(context.Background(), &result)
	return &result, err
}

func (u UserRepository) SecureAccess(body Auth.LoginBody) *User.User {
	filter := bson.M{"email": body.Email}
	var result User.User
	err := u.Ref.Ref.FindOne(context.Background(), filter, options.FindOne()).Decode(&result)
	if err != nil {
		return nil
	}
	if Hasher.ComparePassword(body.Password, result.Password) == false {
		return nil
	}
	return &result
}

func (u UserRepository) ReadByEntry(email string) *User.User {
	filter := bson.M{"email": email}
	var result User.User
	err := u.Ref.Ref.FindOne(context.Background(), filter, options.FindOne()).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}
