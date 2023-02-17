package main

import (
	"log"
	"sync"
	"time"

	"github.com/b02330224/gosocks5/socks5"
)

func main() {
	users := map[string]string{
		"admin":    "123456",
		"zhangsan": "1234",
		"lisi":     "abde",
	}

	var mutex sync.Mutex

	server := socks5.SOCKS5Server{
		IP:   "localhost",
		Port: 1080,
		Config: &socks5.Config{
			AuthMethod: socks5.MethodPassword,
			PasswordChecker: func(username, password string) bool {
				mutex.Lock()
				defer mutex.Unlock()
				wantPassword, ok := users[username]
				if !ok {
					return false
				}
				return wantPassword == password
			},
			TCPTimeout: 5 * time.Second,
		},
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
