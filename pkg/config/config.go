package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/franklincm/roller/pkg/fetch"
	"google.golang.org/api/iam/v1"
)

var (
	CONFIG_DIR  = os.Getenv("HOME") + "/.config/roller"
	CONFIG_FILE = path.Join(CONFIG_DIR, "roles.json")
)

func FetchData() ([]iam.Role, error) {
	// if data exists, unmarshal and return
	var roles []iam.Role
	if _, err := os.Stat(CONFIG_FILE); err == nil {
		content, err := os.ReadFile(CONFIG_FILE)
		if err != nil {
			return roles, err
		}

		err = json.Unmarshal(content, &roles)
		return roles, err
	}

	// confirm data download and save
	input := confirmation.New(
		fmt.Sprintf("download data to: %s (~3.5 MB) ?\n", CONFIG_FILE),
		confirmation.Undecided,
	)
	ready, err := input.RunPrompt()
	if err != nil {
		return roles, err
	}

	if ready {
		err = os.MkdirAll(CONFIG_DIR, 0o755)
		if err != nil {
			return roles, err
		}
		roles, _ = fetch.GetRolesAndPermissions()
		b, _ := json.MarshalIndent(roles, "", "  ")
		err = os.WriteFile(CONFIG_FILE, b, 0o755)
	}

	return roles, err
}
