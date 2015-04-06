package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var (
	u_names = []string{
		"Pepe",
		"Gozalo",
		"Juan",
		"Carolina",
	}
	u_last_names = []string{
		"Escobar",
		"Sierra",
		"Velez",
		"Mejia",
	}
	u_usernames = []string{
		"sescob",
		"jsi3rra",
		"jvlez8",
		"caro27",
	}
	u_emails = []string{
		"pepe27@gmail.com",
		"gozalosierra@gmail.com",
		"juanv@gmail.com",
		"carolina@gmail.com",
	}
	u_passwords = []string{
		"qwerty",
		"123456",
		"AeIoU!@",
		"S3CUR3P455W0RD!\"#$%&/()=",
	}
)

type User struct {
	Username     string
	Email        string
	LastName     string
	Name         string
	PasswordHash string
}

func makeUsers() []User {
	users := []User{}
	for i := 0; i < 50; i++ {
		u := User{
			Username:     u_usernames[i%4] + string(i),
			Email:        u_emails[i%4] + string(i),
			LastName:     u_last_names[i%4] + string(i),
			Name:         u_names[i%4] + string(i),
			PasswordHash: u_passwords[i%4] + string(i),
		}
		users = append(users, u)
	}
	return users
}

func TestMemorySession_Set_and_Get(t *testing.T) {
	t.Parallel()
	users := makeUsers()
	mStore := NewMemorySessionStore("memSession")
	var wg sync.WaitGroup
	for _, u := range users {
		wg.Add(1)
		go func(u User) {
			mStore.Set(u.Username, u)
			userSession := mStore.Get(u.Username)
			assert.Equal(t, u, userSession)
			wg.Done()
		}(u)
	}
	wg.Wait()
}

func TestMemorySession_Set_and_Delete(t *testing.T) {
	t.Parallel()
	users := makeUsers()
	mStore := NewMemorySessionStore("memSession")
	var wg sync.WaitGroup
	for _, u := range users {
		wg.Add(1)
		go func(u User) {
			mStore.Set(u.Username, u)
			mStore.Delete(u.Username)
			assert.Nil(t, mStore.Get(u.Username))
			wg.Done()
		}(u)
	}
	wg.Wait()
}
