package casbin

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestNewEnforcerWithGorm(t *testing.T) {
	db, err := gorm.Open(mysql.Open(os.Getenv("dsn")), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open mysql: %s", err)
	}
	e, err := NewEnforcerWithGorm(db)
	if err != nil {
		t.Fatalf("failed to create enforcer: %s", err)
	}
	if _, err := e.AddPolicy("admin", "/user/list", "read"); err != nil {
		t.Fatalf("failed to add policy: %s", err)
	}
	if _, err := e.AddGroupingPolicy("alice", "admin"); err != nil {
		t.Fatalf("failed to add group policy: %s", err)
	}
	if ok, err := e.Enforce("alice", "/user/list", "read"); err != nil {
		t.Fatalf("failed to check permission: %s", err)
	} else if !ok {
		t.Fatal("permission to `alice /user/list read` is not as aspect")
	}
	t.Log("all roles: ", e.GetAllRoles())
	t.Log("admin policy: ", e.GetPermissionsForUser("admin"))
	aliceRoles, err := e.GetRolesForUser("alice")
	if err != nil {
		t.Fatalf("failed to get alice roles: %s", err)
	}
	adminUsers, err := e.GetUsersForRole("admin")
	if err != nil {
		t.Fatalf("failed to get admin users: %s", err)
	}
	t.Log("alice has roles: ", aliceRoles)
	t.Log("admin has users: ", adminUsers)
	t.Log("pass")
}
