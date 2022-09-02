package go_git_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/yuchanns/gobyexample/go-git"
)

func Test_AnyTokenRepoBranches(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	anyRepo, err := NewAnyTokenRepo(ctx, "git@github.com:yuchanns/gobyexample.git", "")
	assert.NoError(t, err)

	branches, err := anyRepo.Branches()
	assert.NoError(t, err)

	expectedBranches := map[string]struct{}{"master": {}, "monorepo": {}}

	for _, branch := range branches {
		_, ok := expectedBranches[branch]
		assert.True(t, ok, "unexpected branch %s found", branch)
	}
}
