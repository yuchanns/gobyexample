package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

const rbacModelText = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`

func NewEnforcerWithGorm(db *gorm.DB) (*casbin.Enforcer, error) {
	var (
		a   persist.Adapter
		m   model.Model
		err error
		e   *casbin.Enforcer
	)
	if a, err = gormadapter.NewAdapterByDBUseTableName(db, "", "rbac"); err != nil {
		return nil, err
	}
	if m, err = model.NewModelFromString(rbacModelText); err != nil {
		return nil, err
	}
	if e, err = casbin.NewEnforcer(m, a); err != nil {
		return nil, err
	}
	if err = e.LoadPolicy(); err != nil {
		return nil, err
	}
	return e, nil
}
