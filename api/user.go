// api/user.go
package api

import (
 "encoding/json"
 "fmt"
 "net/http"
"database/sql"

 "github.com/xiexiaohuozi/chatbot-system/database" // Replace with your actual module path
 "github.com/xiexiaohuozi/chatbot-system/model"   // Replace with your actual module path
)

// RegisterUser handles user registration requests
func RegisterUser(db *sql.DB) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {
  var user model.User
  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
   http.Error(w, "Invalid request payload", http.StatusBadRequest)
   return
  }

  if err := database.CreateUser(db, &user); err != nil {
   http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
   return
  }

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(map[string]interface{}{"message": "User registered successfully"})
 }
}

// LoginUser handles user login requests
func LoginUser(db *sql.DB) http.HandlerFunc {
 return func(w http.ResponseWriter, r *http.Request) {
  var user model.User
  if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
   http.Error(w, "Invalid request payload", http.StatusBadRequest)
   return
  }

  storedUser, err := database.GetUserByUsername(db, user.Username)
  if err != nil {
   http.Error(w, "Invalid username or password", http.StatusUnauthorized)
   return
  }

  if storedUser.Password != user.Password {
   http.Error(w, "Invalid username or password", http.StatusUnauthorized)
   return
  }

  json.NewEncoder(w).Encode(map[string]interface{}{"message": "Login successful"})
 }
}