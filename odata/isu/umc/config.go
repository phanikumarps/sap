package umc

import (
	"os"

	"github.com/phanikumarps/sap/odata/config"
)

func DefaultUmcRootPath() *config.RootPath {
	r := os.Getenv("UMC_ROOT_PATH")
	if r != "" {
		p := config.DefaultRootPath(r)
		return p
	}
	return nil
}
