package git

import (
	"context"
	"os/exec"
	"path/filepath"
)

func (g GitController) CreateRepo(
	ctx context.Context,
	path string,
) error {
	// In case a earlier delete failed
	//
	// Repo is added to DB first and returns an
	// error in case of duplicate so assuming caller
	// checks the error, it is safe to delete the repo
	// This assumes a repo not in the db is a repo upon which delete failed
	//
	// DeleteRepo repos also does not error on a repo that does not exist
	if err := g.DeleteRepo(path); err != nil {
		return err
	}

	return exec.CommandContext(
		ctx, "git", "init", "--bare", filepath.Join(g.basePath, path),
	).Run()
}
