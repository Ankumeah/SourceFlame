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

func (g GitController) GetRefs(
	ctx context.Context,
	path string,
) ([]Ref, error) {
	if _, err := os.Stat(
		filepath.Join(g.basePath, path),
	); errors.Is(err, os.ErrNotExist) {
		return nil, errs.ErrInvalidRepo
	}

	buf := bytes.NewBuffer(nil)

	cmd := exec.CommandContext(ctx, "git", "-C", filepath.Join(g.basePath, path), "for-each-ref")
	cmd.Stdout = buf

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	refs := []Ref{}

	for ref := range strings.SplitSeq(buf.String(), "\n") {
		if ref == "" {
			continue
		}

		fields := strings.Fields(ref)
		if len(fields) != 3 {
			return nil, errors.New("len(fields) in git for-each-ref != 3")
		}

		refs = append(refs, Ref{
			Hash: fields[0],
			Name: fields[2],
		})
	}

	return refs, nil
}
