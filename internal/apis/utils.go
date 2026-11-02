package apis

import (
	"errors"
	"net/http"

	"github.com/Ankumeah/SourceFlame/internal/errs"
	"github.com/Ankumeah/SourceFlame/internal/utils"
)

func handle(err error, w http.ResponseWriter) bool {
	if errors.Is(err, errs.ErrExists) {
		utils.RespMsg(
			err.Error(),
			http.StatusConflict,
			w,
		)
		return false
	} else if errors.Is(err, errs.ErrInvalid) {
		utils.RespMsg(
			err.Error(),
			http.StatusNotFound,
			w,
		)
		return false
	} else if err != nil {
		utils.RespErr(err, w)
		return false
	}

	return true
}
