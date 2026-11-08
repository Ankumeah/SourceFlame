package git

import (
	"errors"
	"path/filepath"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/object"

	"github.com/Ankumeah/SourceFlame/internal/errs"
)

func (g GitController) GetCommits(
	path string,
	ref string,
) ([]Commit, error) {
	repo, err := git.PlainOpen(
		filepath.Join(g.basePath, path),
	)

	if errors.Is(err, git.ErrRepositoryNotExists) {
		return nil, errs.ErrInvalidRepo
	}

	hash, ok := plumbing.FromHex(ref)
	if !ok {
		return nil, errs.ErrInvalidRef
	}

	commitsIter, err := repo.Log(
		&git.LogOptions{
			From: hash,
		},
	)
  if errors.Is(err, plumbing.ErrObjectNotFound) {
    return nil, errs.ErrInvalidRef
  } else if err != nil {
		return nil, err
	}
	defer commitsIter.Close()

	commits := []Commit{}
	commitsIter.ForEach(
		func(c *object.Commit) error {
			commits = append(commits, Commit{
				Hash: c.Hash.String(),
				Author: Author{
					Name:  c.Author.Name,
					Email: c.Author.Email,
				},
				Message: c.Message,
			})

      return nil
		},
	)

  return commits, nil
}
