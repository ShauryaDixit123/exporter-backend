package helper

import (
	"encoding/json"
	"exporterbackend/internal/common"
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"fmt"
	"os"
	"strings"
)

func (h *HelperRepository) CheckForPermissions(p common.PermissionCheck) bool {
	perms, er := readPermissionsJSON("/Users/shauryadixit007/Documents/exp-full/exp/internal/permissions.json")
	if er != nil {
		h.logger.Error(
			"permissions.json",
			"permissions parsing failed",
			er,
			map[string]any{},
			map[string]any{
				"permissions": perms,
			},
		)
		return false
	}
	role, er := h.rolesRepo.GetById(rdbms.Id{
		Id: fmt.Sprint(p.RoleId),
	})
	if er != nil {
		h.logger.Error(
			"role get failed",
			"permissions parsing failed",
			er,
			map[string]any{},
			map[string]any{
				"permissions": perms,
			},
		)
		return false
	}
	fmt.Println(perms, "perms", p.Action)
	// admin parsing to be added.
	if role.Role == "buyer" {
		for _, v := range perms.Buyer.Action {
			if v == p.Action {
				return true
			}
		}
	}
	if role.Role == "supplier" {
		for _, v := range perms.Supplier.Action {
			if v == p.Action {
				return true
			}
		}
	}
	return false
}

func readPermissionsJSON(filePath string) (*common.Permissions, error) {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Decode the JSON into the struct
	var permissions common.Permissions
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&permissions); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &permissions, nil
}

func (r *HelperRepository) ParseURLAndAction(url, method string) string {
	// Remove the prefix "/v1/"
	parts := strings.SplitN(url, "/v1/", 2)
	if len(parts) < 2 {
		return ""
	}

	// Split the remaining part into segments
	path := strings.Trim(parts[1], "/") // Remove leading/trailing slashes
	segments := strings.Split(path, "/")

	// Determine the relevant value based on the number of segments
	var resource string
	if len(segments) == 1 {
		resource = segments[0] // Only one segment, take it as resource
	} else {
		resource = segments[len(segments)-1] // Take the last segment
	}

	// For deeper nested paths (more than 2 segments), join with underscores
	if len(segments) > 2 {
		resource = strings.Join(segments[1:], "_")
	}

	// Map the HTTP method to the action
	methodToAction := map[string]string{
		"GET":    "read",
		"POST":   "create",
		"PUT":    "update",
		"PATCH":  "update",
		"DELETE": "delete",
		"READ":   "read",
		"UPDATE": "update",
		"CREATE": "create",
	}

	action, exists := methodToAction[method]
	if !exists {
		action = "unknown" // Default to "unknown" for unsupported methods
	}

	// Combine resource and action
	return fmt.Sprintf("%s.%s", resource, action)
}
