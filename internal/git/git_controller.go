package git

type GitController struct {
	basePath string
}

func GetGitController(basePath string) GitController {
	return GitController{
		basePath: basePath,
	}
}
