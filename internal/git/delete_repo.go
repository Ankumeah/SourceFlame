package git

import (
	"os"
)

func DeleteRepo(
	path string,
) error {
	return os.RemoveAll(path)
}
