// main
package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
)

func main() {
	log.Println("Starting server")

	configFile := "./include.conf"
	file, err := os.Open(configFile)
	if err != nil {
		panic("Could not read file")
	}

	roleMap, _ := getRoleAndPaths(file)
	redactedDir := NewFileSystem{Dir: http.Dir("./"), roleMap: RoleMap(roleMap)}
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(redactedDir)))
}

func getRole() Role {
	// TODO: Add authentication
	return Role("MINIMAL")
}
