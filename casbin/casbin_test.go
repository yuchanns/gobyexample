package casbin_test

import (
	casbin2 "github.com/casbin/casbin/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	. "github.com/yuchanns/gobyexample/casbin"
)

var _ = Describe("Casbin", func() {
	var (
		db  *gorm.DB
		e   *casbin2.Enforcer
		err error
	)

	Describe("Connect Database and Create Enforcer", func() {
		It("should open sqlite", func() {
			db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
			Expect(err).NotTo(HaveOccurred())
		})
		It("should create enforcer", func() {
			e, err = NewEnforcerWithGorm(db)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Add Policy", func() {
		Context("With Add Policy", func() {
			It("should add a role admin", func() {
				_, err = e.AddPolicy("admin", "/user/list", "read")
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("With Add Group Policy", func() {
			It("should add role admin for user alice", func() {
				_, err = e.AddGroupingPolicy("alice", "admin")
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Describe("Access Object", func() {
		Context("With access object", func() {
			It("should access /user/list by read with user alice", func() {
				ok, err := e.Enforce("alice", "/user/list", "read")
				Expect(err).NotTo(HaveOccurred())
				Expect(ok).To(Equal(true))
			})
		})
		Context("With not access object", func() {
			It("should not access /user/list by write with user alice", func() {
				ok, err := e.Enforce("alice", "/user/list", "write")
				Expect(err).NotTo(HaveOccurred())
				Expect(ok).To(Equal(false))
			})
		})
	})

	Describe("Get Roles and Users", func() {
		Context("With Get All Data", func() {
			It("should get all roles", func() {
				roles := e.GetAllRoles()
				Expect(roles).To(Equal([]string{"admin"}))
			})
			It("should get admin permissions", func() {
				permissions := e.GetPermissionsForUser("admin")
				Expect(permissions).To(Equal([][]string{{"admin", "/user/list", "read"}}))
			})
		})
		Context("With Role and User", func() {
			It("should get roles for user alice", func() {
				roles, err := e.GetRolesForUser("alice")
				Expect(err).NotTo(HaveOccurred())
				Expect(roles).To(Equal([]string{"admin"}))
			})
			It("should get users for role admin", func() {
				roles, err := e.GetUsersForRole("admin")
				Expect(err).NotTo(HaveOccurred())
				Expect(roles).To(Equal([]string{"alice"}))
			})
		})
	})
})
