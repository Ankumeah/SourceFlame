package git

import (
	"context"
	"os/exec"
)

func CreateRepo(
	ctx context.Context,
	path string,
) error {
	return exec.CommandContext(
		ctx, "git", "init", "--bare", path,
	).Run()
}
