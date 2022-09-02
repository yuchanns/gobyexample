package go_git

import (
	"context"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	giturls "github.com/whilp/git-urls"
)

// AnyTokenRepo support
type AnyTokenRepo struct {
	repo *git.Repository
}

// NewAnyTokenRepo ...
func NewAnyTokenRepo(ctx context.Context, rawURL, token string) (*AnyTokenRepo, error) {
	repoURL, err := giturls.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	url := repoURL.String()
	if repoURL.Scheme == "ssh" {
		url = "https://" + repoURL.Host + "/" + repoURL.RequestURI()
	}
	anyRepo := &AnyTokenRepo{}
	return anyRepo.clone(ctx, url, token)
}

// Branches ...
func (r *AnyTokenRepo) Branches() ([]string, error) {
	refIter, err := r.repo.Branches()
	if err != nil {
		return nil, err
	}
	defer refIter.Close()

	var branches []string

	refIter.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsBranch() {
			branches = append(branches, ref.Name().Short())
		}
		return nil
	})

	return branches, nil
}

// Tags ...
func (r *AnyTokenRepo) Tags() ([]string, error) {
	refIter, err := r.repo.Tags()
	if err != nil {
		return nil, err
	}
	defer refIter.Close()

	var tags []string

	refIter.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsTag() {
			tags = append(tags, ref.Name().Short())
		}
		return nil
	})

	return tags, nil
}

func (r *AnyTokenRepo) clone(ctx context.Context, repoURL, token string) (*AnyTokenRepo, error) {
	repo, err := git.CloneContext(ctx, memory.NewStorage(), nil, &git.CloneOptions{
		URL: repoURL,
		Auth: &http.BasicAuth{
			Username: "",
			Password: token, // for private access
		},
		NoCheckout: true,
		Depth:      1,
	})

	if err == nil {
		r.repo = repo
	}

	return r, err
}
