package git

import (
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Ankumeah/SourceFlame/internal/errs"
)

func (g GitController) ReceivePack(
	ctx context.Context,
	path string,
	reader io.Reader,
	writer io.Writer,
) error {
	if _, err := os.Stat(
		filepath.Join(g.basePath, path),
	); errors.Is(err, os.ErrNotExist) {
		return errs.ErrInvalidRepo
	}

	cmd := exec.CommandContext(
		ctx,
		"git-receive-pack",
		"--stateless-rpc",
		filepath.Join(g.basePath, path),
	)
	cmd.Stdin = reader
	cmd.Stdout = writer

	return cmd.Run()
}
