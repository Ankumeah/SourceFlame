package git

import (
	"context"
	"io"
	"os/exec"
)

func UploadPack(
	ctx context.Context,
	path string,
	reader io.Reader,
	writer io.Writer,
) error {
	cmd := exec.CommandContext(
		ctx,
		"git-upload-pack",
		"--stateless-rpc",
		path,
	)
	cmd.Stdin = reader
	cmd.Stdout = writer

	return cmd.Run()
}
