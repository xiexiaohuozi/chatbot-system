package model

import "time"

type Message struct {
  ID        int       `json:"id"`
  UserID    int       `json:"user_id"`
  Content   string    `json:"content"`
  Timestamp time.Time `json:"timestamp"`
}