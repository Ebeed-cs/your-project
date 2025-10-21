package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type User struct {
	ID      int
	Name    string
	Age     int
	Address Address
}

var (
	nextID   = 1
	usersDir = "users_saved"
)

func init() {
	if err := os.MkdirAll(usersDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create users directory: %v", err))
	}

	nextID = countExistingUsers() + 1
}

func countExistingUsers() int {
	files, err := ioutil.ReadDir(usersDir)
	if err != nil {
		return 0
	}

	maxID := 0
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".txt" {
			idStr := file.Name()[:len(file.Name())-4]
			if id, err := strconv.Atoi(idStr); err == nil {
				if id > maxID {
					maxID = id
				}
			}
		}
	}
	return maxID
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}

	u.ID = nextID
	nextID++

	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return User{}, fmt.Errorf("Failed to marshal user: %v", err)
	}

	filename := filepath.Join(usersDir, fmt.Sprintf("%d.txt", u.ID))

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return User{}, fmt.Errorf("Failed to write user file: %v", err)
	}

	return u, nil
}

func GetUserByID(id int) (User, error) {
	filename := filepath.Join(usersDir, fmt.Sprintf("%d.txt", id))

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return User{}, fmt.Errorf("User with ID '%v' not found", id)
	}

	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return User{}, fmt.Errorf("Failed to read user file: %v", err)
	}

	var u User
	err = json.Unmarshal(jsonData, &u)
	if err != nil {
		return User{}, fmt.Errorf("Failed to unmarshal user data: %v", err)
	}

	return u, nil
}