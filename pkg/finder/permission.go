package finder

import (
	"fmt"
	"log"

	"github.com/franklincm/roller/pkg/config"
	"github.com/franklincm/roller/pkg/role"
	"github.com/ktr0731/go-fuzzyfinder"
	"google.golang.org/api/iam/v1"
)

func FindPermissions() {
	var roles []iam.Role

	roles, err := config.FetchData()
	if err != nil {
		log.Fatal(err)
	}

	roleNames := role.GetRoleNames(roles)

	idx, _ := fuzzyfinder.Find(
		roleNames,
		func(i int) string {
			return roleNames[i]
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			s := ""
			rolePermissions := role.PermissionsFromRole(roles, roleNames[i])
			for _, role := range rolePermissions {
				s = fmt.Sprintf("%s\n%s", s, role)
			}
			return s
		}),
	)
	rolePermissions := role.PermissionsFromRole(roles, roleNames[idx])
	for _, permission := range rolePermissions {
		fmt.Printf("%s\n", permission)
	}
}
