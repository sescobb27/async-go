// +build ignore

package main

import (
	"fmt"
	"sync"
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
		"pep66",
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
	for i := 0; i < 10; i++ {
		u := User{
			Username:     u_usernames[i%4],
			Email:        u_emails[i%4],
			LastName:     u_last_names[i%4],
			Name:         u_names[i%4],
			PasswordHash: u_passwords[i%4],
		}
		users = append(users, u)
	}
	return users
}

// START PARALLELFORLOOP OMIT
func ParallelForLoop() {
	users := makeUsers()
	var wg sync.WaitGroup
	for _, u := range users {
		wg.Add(1)
		go func(u User) {
			fmt.Printf("%s: (%s)\n", u.Email, u.Username)
			wg.Done()
		}(u)
	}
	wg.Wait()
}

// END PARALLELFORLOOP OMIT
func main() {
	ParallelForLoop()
}
