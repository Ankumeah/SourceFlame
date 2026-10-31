package git

import (
	"os"
	"path/filepath"
)

func (g GitController) DeleteRepo(
	path string,
) error {
	return os.RemoveAll(filepath.Join(g.basePath, path))
}
