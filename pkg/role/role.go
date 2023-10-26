package role

import (
	"slices"
	"sort"

	"google.golang.org/api/iam/v1"
)

func GetRoleNames(roles []iam.Role) []string {
	roleNames := make([]string, 0, len(roles))
	for _, role := range roles {
		roleNames = append(roleNames, role.Name)
	}
	sort.Slice(roleNames, func(i, j int) bool { return roleNames[i] < roleNames[j] })
	return roleNames
}

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

func PermissionsFromRole(roles []iam.Role, roleName string) []string {
	for _, role := range roles {
		if role.Name == roleName {
			return role.IncludedPermissions
		}
	}
	return []string{}
}

func RolesFromPermission(roles []iam.Role, permission string) []string {
	matchedRoles := make([]string, 0)
	for _, role := range roles {
		if slices.Contains(role.IncludedPermissions, permission) {
			matchedRoles = append(matchedRoles, role.Name)
		}
	}
	sort.Slice(matchedRoles, func(i, j int) bool { return matchedRoles[i] < matchedRoles[j] })
	return matchedRoles
}
