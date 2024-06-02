// api/message.go
package api

import (
 "encoding/json"
 "fmt"
 "net/http"
 "time"
"database/sql"

 "github.com/xiexiaohuozi/chatbot-system/database" // Replace with your actual module path
 "github.com/xiexiaohuozi/chatbot-system/model"   // Replace with your actual module path
)

// SendMessage handles sending a new message
func SendMessage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var message model.Message
		if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		message.Timestamp = time.Now()
		if err := database.InsertMessage(db, &message); err != nil {
			http.Error(w, fmt.Sprintf("Failed to send message: %v", err), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Message sent successfully"})
	}
}

// GetMessages retrieves all messages
func GetMessages(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messages, err := database.GetMessages(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get messages: %v", err), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(messages)
	}
}