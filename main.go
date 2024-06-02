package main

import (
 //"fmt"
 "log"
 "net/http"
 //"database/sql"

 "github.com/gorilla/mux"
 _ "github.com/go-sql-driver/mysql"
 "github.com/xiexiaohuozi/chatbot-system/api"      // 替换为你的模块路径
 "github.com/xiexiaohuozi/chatbot-system/database" // 替换为你的模块路径
)

func main() {
 // 数据库配置
 dbConfig := &database.DBConfig{
  Username: "root",
  Password: "1234",
  Host:     "localhost",
  Port:     "3306",
  DBName:   "chatbot",
 }

 // 初始化数据库连接
 db, err := database.InitDB(dbConfig)
 if err != nil {
  log.Fatal(err)
 }
 defer db.Close()

 // 路由配置
 router := mux.NewRouter()
 router.HandleFunc("/", serveHome)



 // 用户 API 路由
 userAPI := router.PathPrefix("/api/user").Subrouter()
 userAPI.HandleFunc("/register", api.RegisterUser(db)).Methods("POST")
 userAPI.HandleFunc("/login", api.LoginUser(db)).Methods("POST")

 // 消息 API 路由
messageAPI := router.PathPrefix("/api/message").Subrouter()
messageAPI.HandleFunc("/send", api.SendMessage(db)).Methods("POST")
messageAPI.HandleFunc("/get", api.GetMessages(db)).Methods("GET")



 // 静态文件服务
 router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

 // 启动服务器
 log.Fatal(http.ListenAndServe(":8080", router))
}

// serveHome 处理主页请求
func serveHome(w http.ResponseWriter, r *http.Request) {
 http.ServeFile(w, r, "static/index.html")
}