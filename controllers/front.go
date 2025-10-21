package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	// ✅ route رئيسي لصفحة الترحيب
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "Welcome to Go User API 🚀")
		fmt.Fprintln(w, "Available routes:")
		fmt.Fprintln(w, "- /users")
	})

	// ✅ route خاص بالمستخدمين
	http.Handle("/users", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
