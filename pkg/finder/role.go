package finder

import (
	"fmt"
	"log"

	"github.com/franklincm/roller/pkg/config"
	"github.com/franklincm/roller/pkg/role"
	"github.com/ktr0731/go-fuzzyfinder"
	"google.golang.org/api/iam/v1"
)

func FindRoles() {
	var roles []iam.Role

	roles, err := config.FetchData()
	if err != nil {
		log.Fatal(err)
	}
	uniquePermissions := role.GetUniquePermissions(roles)

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
			matchedRoles := role.RolesFromPermission(roles, uniquePermissions[i])
			for _, role := range matchedRoles {
				s = fmt.Sprintf("%s\n%s", s, role)
			}
			return s
		}),
	)
	matchedRoles := role.RolesFromPermission(roles, uniquePermissions[idx])
	for _, role := range matchedRoles {
		fmt.Printf("%s\n", role)
	}
}
