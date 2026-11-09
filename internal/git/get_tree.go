package git

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Ankumeah/SourceFlame/internal/errs"
)

func (g GitController) GetTree(
	ctx context.Context,
	path string,
	ref string,
	dir string,
) ([]File, error) {
	if _, err := os.Stat(
		filepath.Join(g.basePath, path),
	); errors.Is(err, os.ErrNotExist) {
		return nil, errs.ErrInvalidRepo
	}

	if dir == "" {
		dir = "."
	}

	buf := bytes.NewBuffer(nil)

	cmd := exec.CommandContext(
		ctx,
		"git", "-C",
		filepath.Join(g.basePath, path),
		"ls-tree",
		"--format",
		"%(objecttype) %(path)\n",
		ref, dir,
	)
	cmd.Stdout = buf

	// Potensial issue
	// We treat all errors as no commits
	// os exec has no clean way to check, we might want
	// to change this if we ever find a way to cleanly check
	if err := cmd.Run(); err != nil {
		return nil, errs.ErrInvalidRef
	}

	files := []File{}

	for entry := range strings.SplitSeq(buf.String(), "\n") {
		if strings.TrimSpace(entry) == "" {
			continue
		}

		fields := strings.Fields(entry)
		if len(fields) != 2 {
			return nil, errors.New("len(fields) != 2 in git ls-tree")
		}

		files = append(files, File{
			Name:   fields[1],
			Dir:    fields[0] == "tree" || fields[0] == "commit",
			Submod: fields[0] == "commit",
		})
	}

	return files, nil
}
