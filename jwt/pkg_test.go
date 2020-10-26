package jwt

import (
	"testing"
)

func TestNewJwt(t *testing.T) {
	var (
		id   = 1
		name = "yuchanns"
	)

	jwt := NewJwt()
	tokenString, err := jwt.Generate(id, name)

	if err != nil {
		t.Errorf("failed to generate token: %+v", err)
		return
	}

	if err := jwt.Validate(tokenString); err != nil {
		t.Errorf("token validate failed: %+v", err)
		return
	}

	if err := jwt.Validate(tokenString); err != nil {
		t.Errorf("token validate failed: %+v", err)
		return
	}

	if err := jwt.Invalidate(tokenString); err != nil {
		t.Errorf("failed to delete token: %+v", err)
		return
	}

	if err := jwt.Validate(tokenString); err == nil {
		t.Errorf("err should not be equal to nil: %+v", err)
		return
	}
}
