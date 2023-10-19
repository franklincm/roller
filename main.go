package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/ktr0731/go-fuzzyfinder"
	"golang.org/x/exp/slices"
)

type Role struct {
	Description         string   `json:"description"`
	Etag                string   `json:"etag"`
	IncludedPermissions []string `json:"includedPermissions"`
	Name                string   `json:"name"`
	Stage               string   `json:"stage"`
	Title               string   `json:"title"`
}

func getUniquePermissions(roles []Role) []string {
	permissions := make(map[string]bool)
	for _, r := range roles {
		for _, p := range r.IncludedPermissions {
			if _, ok := permissions[p]; !ok {
				permissions[p] = true
			}
		}
	}

	uniquePermissions := make([]string, 0, len(permissions))
	for v := range permissions {
		uniquePermissions = append(uniquePermissions, v)
	}
	sort.Slice(uniquePermissions, func(i, j int) bool { return uniquePermissions[i] < uniquePermissions[j] })
	return uniquePermissions
}

func rolesFromPermission(roles []Role, permission string) []string {
	matchedRoles := make([]string, 0)
	for _, role := range roles {
		if slices.Contains(role.IncludedPermissions, permission) {
			matchedRoles = append(matchedRoles, role.Name)
		}
	}
	return matchedRoles
}

func main() {
	content, err := os.ReadFile("permissions-formatted.json")
	if err != nil {
		log.Fatal(err)
	}

	var roles []Role
	err = json.Unmarshal(content, &roles)
	if err != nil {
		log.Fatal(err)
	}
	uniquePermissions := getUniquePermissions(roles)

	idx, _ := fuzzyfinder.Find(
		uniquePermissions,
		func(i int) string {
			return uniquePermissions[i]
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			s := ""
			matchedRoles := rolesFromPermission(roles, uniquePermissions[i])
			for _, role := range matchedRoles {
				s = fmt.Sprintf("%s\n%s", s, role)
			}
			return s
		}),
	)
	matchedRoles := rolesFromPermission(roles, uniquePermissions[idx])
	for _, role := range matchedRoles {
		fmt.Printf("%s\n", role)
	}
}
