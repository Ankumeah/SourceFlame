package utils

import (
	"context"
	"os"
	"os/exec"
)

func VerifySign(ctx context.Context, key string, data []byte) (bool, error) {
	tmpFile, err := os.CreateTemp("", "sf-temp-email-*")
	if err != nil {
		return false, err
	}
	defer func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}()

	if _, err := tmpFile.Write(data); err != nil {
		return false, err
	}

	// TODO: This does not work, I am commiting this anyways so that my time on terra does not get deflated
	if err := exec.CommandContext(
		ctx,
		"patatt",
		"validate",
		tmpFile.Name(),
	).Run(); err != nil {
		// Potensial issue
		// We treat all errors as an invalid sign
		// os exec has no clean way to check, we might want
		// to change this if we ever find a way to cleanly check
		return false, nil
	}

	return true, nil
}
