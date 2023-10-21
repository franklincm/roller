package role

import (
	"slices"
	"sort"

	"google.golang.org/api/iam/v1"
)

func GetUniquePermissions(roles []iam.Role) []string {
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

func RolesFromPermission(roles []iam.Role, permission string) []string {
	matchedRoles := make([]string, 0)
	for _, role := range roles {
		if slices.Contains(role.IncludedPermissions, permission) {
			matchedRoles = append(matchedRoles, role.Name)
		}
	}
	return matchedRoles
}