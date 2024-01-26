package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Desc     string `json:"desc"`
}

func main() {

	// 注册路由处理函数
	http.HandleFunc("/users/", getUserHandler)

	// 启动 HTTP 服务器
	log.Println("Server started on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// 从 URL 中解析出用户 ID
	id := strings.TrimPrefix(r.URL.Path, "/users/")

	// 检查请求方法是否为 POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 将用户 ID 转换为整数
	userID := 0
	_, err := fmt.Sscanf(id, "%d", &userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// 查找用户
	user := User{
		ID:       userID,
		Username: fmt.Sprintf("user%v", userID),
		Desc:     "v1",
	}

	// 将用户转换为 JSON 格式
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 发送响应
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
