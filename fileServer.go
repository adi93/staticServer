package main

import (
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
