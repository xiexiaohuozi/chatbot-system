// database/message.go
package database

import (
 "fmt"
 "database/sql"//导入sql包
 "github.com/xiexiaohuozi/chatbot-system/model" // Replace with your actual module path

)


// InsertMessage inserts a new message into the database


func InsertMessage(db *sql.DB, message *model.Message) error {
 stmt, err := db.Prepare("INSERT INTO messages(user_id, content, timestamp) VALUES(?, ?, ?)")
 if err != nil {
  return fmt.Errorf("failed to prepare statement: %w", err)
 }
 defer stmt.Close()

 result, err := stmt.Exec(message.UserID, message.Content, message.Timestamp)
 if err != nil {
  return fmt.Errorf("failed to insert message: %w", err)
 }

 lastInsertID, err := result.LastInsertId()
 if err != nil {
  return fmt.Errorf("failed to get last inserted ID: %w", err)
 }

 message.ID = int(lastInsertID)
 return nil
}

// GetMessages retrieves all messages from the database
func GetMessages(db *sql.DB) ([]*model.Message, error) {
 rows, err := db.Query("SELECT id, user_id, content, timestamp FROM messages ORDER BY timestamp ASC")
 if err != nil {
  return nil, fmt.Errorf("failed to get messages: %w", err)
 }
 defer rows.Close()

 var messages []*model.Message
 for rows.Next() {
  message := &model.Message{}
  err := rows.Scan(&message.ID, &message.UserID, &message.Content, &message.Timestamp)
  if err != nil {
   return nil, fmt.Errorf("failed to scan message: %w", err)
  }
  messages = append(messages, message)
 }

 return messages, nil
}