package fetch

import (
	"context"
	"sync"

	"google.golang.org/api/iam/v1"
)

func getRoles(svc *iam.Service) ([]iam.Role, error) {
	var roles []iam.Role
	ctx := context.Background()
	err := svc.Roles.List().Pages(ctx, func(r *iam.ListRolesResponse) error {
		for _, role := range r.Roles {
			roles = append(roles, *role)
		}
		return nil
	})
	return roles, err
}

func GetRolesAndPermissions() ([]iam.Role, error) {
	svc, err := iam.NewService(context.Background())
	if err != nil {
		return []iam.Role{}, err
	}

	roleNames, err := getRoles(svc)
	if err != nil {
		return []iam.Role{}, err
	}

	var waitGroup sync.WaitGroup
	roles := make([]iam.Role, 0)
	roleChannel := make(chan *iam.Role)

	numThreads := 512
	start := 0
	for i := 0; i < numThreads; i++ {
		stop := start + (len(roleNames) / numThreads)

		if i < (len(roleNames) % numThreads) {
			stop += 1
		}

		waitGroup.Add(1)

		go func(
			role []iam.Role,
			service *iam.Service,
			rolesCh chan *iam.Role,
		) {
			defer waitGroup.Done()

			for _, role := range role {
				resp, _ := service.Roles.Get(role.Name).Do()
				rolesCh <- resp
			}
		}(roleNames[start:stop], svc, roleChannel)

		start = stop
	}

	go func() {
		waitGroup.Wait()
		close(roleChannel)
	}()

	for role := range roleChannel {
		roles = append(roles, *role)
	}
	return roles, nil
}
