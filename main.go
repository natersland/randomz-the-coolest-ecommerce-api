package main

import (
	"os"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}
func main() {
	// config := configs.ServerConfig(envPath())

	// TODO connect to database

	// TODO start server

}
