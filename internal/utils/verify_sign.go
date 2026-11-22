package utils

import (
	"context"
)

func VerifySign(ctx context.Context, key string, data []byte) (bool, error) {
	// Treating all patches as unverified.
	// As Im working on PRs instead I do not want
	// to spend effort here but I though I should get
	// it to alreday a semi working state
	//
	// TODO: Check signs
	return false, nil
}
