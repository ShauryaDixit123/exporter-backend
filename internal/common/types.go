package common

type PermissionCheck struct {
	RoleId int `json:"role_id"`
	Action string
}
type RolePermissions struct {
	Action []string `json:"action"`
}

// Permissions represents the root JSON structure
type Permissions struct {
	Buyer    RolePermissions `json:"buyer"`
	Supplier RolePermissions `json:"supplier"`
	Admin    RolePermissions `json:"admin"`
}
