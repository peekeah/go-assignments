package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Message struct {
	Id         int
	SenderId   int
	ReceiverId int
	Content    string
	Timestamp  time.Time
}

type User struct {
	Id   int
	Name string
}

type UserStore struct {
	Users []User
}

type MessageStore struct {
	Messages []Message
}

func readFileContent(path string) (users []User, messages []Message, errMsg error) {
	content, err := os.ReadFile(path)
	if err != nil {
		errMsg = errors.New("error while reading file")
	}
	if path == "chat_simulator_messages.json" {
		json.Unmarshal(content, &messages)
	}
	if path == "chat_simulator_users.json" {
		json.Unmarshal(content, &users)
	}
	return
}

func saveFileContent(path string, content []byte) error {
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		return errors.New("error while storing data")
	}
	return nil
}

func InitializeMsgStore() *MessageStore {
	_, msgStore, err := readFileContent("chat_simulator_messages.json")
	if err != nil {
		fmt.Println(err)
	}
	return &MessageStore{msgStore}
}

func InitializeUserStore() *UserStore {
	users, _, err := readFileContent("chat_simulator_users.json")
	if err != nil {
		fmt.Println(err)
	}
	return &UserStore{users}
}

func (m *MessageStore) SendMessage(sender *User, receiver *User, content string) {
	newId := rand.Intn(900000) + 100000
	m.Messages = append(m.Messages, Message{newId, sender.Id, receiver.Id, content, time.Now()})
	data, err := json.Marshal(m.Messages)
	if err != nil {
		fmt.Println("error while parsing data")
	}
	saveFileContent("chat_simulator_messages.json", data)
}

func (m *MessageStore) DeleteMessage(messageId int) error {
	for id := range m.Messages {
		if m.Messages[id].Id == messageId {
			m.Messages = append(m.Messages[:id], m.Messages[id+1:]...)
			data, err := json.Marshal(m.Messages)
			if err != nil {
				fmt.Println("error while parsing data")
			}
			saveFileContent("chat_simulator_messages.json", data)
			return nil
		}
	}
	return errors.New("message not found")
}

func (m *MessageStore) GetMessages() []Message {
	return m.Messages
}

func (u *UserStore) GetUsers() []User {
	return u.Users
}

func (u *UserStore) AddUser(name string) {
	newId := rand.Intn(900000) + 100000
	u.Users = append(u.Users, User{newId, name})
	data, err := json.Marshal(u.Users)
	if err != nil {
		fmt.Println("error while parsing data")
	}
	saveFileContent("chat_simulator_users.json", data)
}

func (u *UserStore) GetUser(userId int) (*User, error) {
	for id := range u.Users {
		if u.Users[id].Id == userId {
			return &u.Users[id], nil
		}
	}
	return nil, errors.New("user not found")
}
