// main
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type RedactedDir struct {
	http.Dir
	roleMap RoleMap
}

func (d RedactedDir) Open(name string) (http.File, error) {
	log.Printf("name: %v", name)
	role := getRole()

	matched := false
	for _, v := range d.roleMap.get(role) {
		if v.MatchString(name) {
			matched = true
			break
		}
	}
	if !matched {
		return nil, os.ErrNotExist
	}
	return d.Dir.Open(name)
}

func main() {
	fmt.Println("Hello World!")

	configFile := "./include.conf"
	file, err := getFile(configFile)
	if err != nil {
		panic("Config error")
	}

	roleMap, _ := getRoleAndPaths(file)
	redactedDir := RedactedDir{Dir: http.Dir("./"), roleMap: RoleMap(roleMap)}
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(redactedDir)))
}

func getRole() Role {
	return Role("MINIMAL")
}
