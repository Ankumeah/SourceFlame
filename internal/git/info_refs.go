package git

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Ankumeah/SourceFlame/internal/errs"
)

func (g GitController) InfoRefs(
	ctx context.Context,
	path string,
	service string,
	writer io.Writer,
) error {
	if _, err := os.Stat(
		filepath.Join(g.basePath, path),
	); errors.Is(err, os.ErrNotExist) {
		return errs.ErrInvalidRepo
	}

	pktLine := "# service=" + service + "\n"
	fmt.Fprintf(writer, "%04x%v0000", len(pktLine)+4, pktLine)

	cmd := exec.CommandContext(
		ctx,
		service,
		"--stateless-rpc",
		"--advertise-refs",
		filepath.Join(g.basePath, path),
	)
	cmd.Stdout = writer

	return cmd.Run()
}
