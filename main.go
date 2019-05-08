// main
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type NewFileSystem struct {
	http.Dir
	roleMap RoleMap
}

func (d NewFileSystem) Open(name string) (http.File, error) {
	log.Printf("name: %v", name)
	// do not display non accessible content
	role := getRole()

	matched := false
	for _, v := range d.roleMap.get(role) {
		if v.MatchString(name) {
			matched = true
			break
		}
	}
	if !matched {
		log.Printf("Permission denied for : %v", name)
		return nil, os.ErrNotExist
	}

	// fetch file
	file, err := d.Dir.Open(name)
	if err != nil {
		return nil, err
	}

	return file, err
	// TODO: Stop displaying directories. Or not?
}

func main() {
	fmt.Println("Hello World!")

	configFile := "./include.conf"
	file, err := os.Open(fileName)
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
