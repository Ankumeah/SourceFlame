package git

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Ankumeah/SourceFlame/internal/errs"
)

func (g GitController) GetLog(
	ctx context.Context,
	path string,
	ref string,
) ([]byte, error) {
	if _, err := os.Stat(
		filepath.Join(g.basePath, path),
	); errors.Is(err, os.ErrNotExist) {
		return nil, errs.ErrInvalidRepo
	}

	buf := bytes.NewBuffer(nil)

	cmd := exec.CommandContext(ctx, "git", "-C", filepath.Join(g.basePath, path), "log", "--oneline")
	cmd.Stdout = buf

	// Potensial issue
	// We treat all errors as no commits
	// os exec has no clean way to check, we might want
	// to change this if we ever find a way to cleanly check
	if err := cmd.Run(); err != nil {
		return nil, errs.ErrNotExistCommits
	}

	return buf.Bytes(), nil
}
