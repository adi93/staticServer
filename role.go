package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
)

type Role string
type matchers []*regexp.Regexp
type RoleMap map[Role][]*regexp.Regexp

func (rm RoleMap) get(r Role) []*regexp.Regexp {
	return rm[r]
}

var MINIMAL Role = "MINIMAL"

func GetRole() Role {
	return MINIMAL
}

// getRoleAndPaths returns mapping between various roles and the path
// regexs they support
func getRoleAndPaths(reader io.Reader) (map[Role][]*regexp.Regexp, error) {
	// loop while all eof file is ot reached
	// and keep track of current role
	rolePaths := make(map[Role][]*regexp.Regexp)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	currentRole := Role(scanner.Text())
	for scanner.Scan() {
		line := scanner.Text()
		if emptyLine(line) {
			scanner.Scan()
			currentRole = Role(scanner.Text())
			continue
		}

		regexPattern, err := regexp.Compile(line)
		if err != nil {
			log.Printf("Invalid regex pattern %v, skipping it", line)
			continue
		}

		rolePaths[currentRole] = append(rolePaths[currentRole], regexPattern)
	}
	return rolePaths, nil
}

func emptyLine(line string) bool {
	if line == "" {
		return true
	}
	return false
}
