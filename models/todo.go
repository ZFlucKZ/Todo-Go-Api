package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZFlucKZ/Todo-Go-Api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Content string `json:"content" bson:"content"`
	IsChecked bool `json:"isChecked" bson:"isChecked"`
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo

	cursor, err := db.DB.Collection("todos").Find(context.Background(), bson.D{{}})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting todos", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	var result Todo

	// defer db.DB.Client().Disconnect(context.Background())

	err := json.NewDecoder(r.Body).Decode(&todo)
	
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	objID, err := primitive.ObjectIDFromHex(todo.Id.Hex())

  if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting Id", http.StatusInternalServerError)
    return
  }

	filter := bson.M{"_id": bson.M{"$eq": objID}}


	err = db.DB.Collection("todos").FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting todo", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert to db
	_, err = db.DB.Collection("todos").InsertOne(context.Background(), todo)

	if err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Todo Created"))
}

func EditTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	// defer db.DB.Client().Disconnect(context.Background())

	err := json.NewDecoder(r.Body).Decode(&todo)
	
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	objID, err := primitive.ObjectIDFromHex(todo.Id.Hex())

  if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting Id", http.StatusInternalServerError)
    return
  }

	db.DB.Collection("todos").FindOneAndUpdate(context.Background(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"isChecked": todo.IsChecked}})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo Updated"))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	// defer db.DB.Client().Disconnect(context.Background())

	err := json.NewDecoder(r.Body).Decode(&todo)
	
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	objID, err := primitive.ObjectIDFromHex(todo.Id.Hex())

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting Id", http.StatusInternalServerError)
		return
	}

	result, err := db.DB.Collection("todos").DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error deleting todo", http.StatusInternalServerError)
		return
	}
	
	fmt.Println(result)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo Deleted"))
}
