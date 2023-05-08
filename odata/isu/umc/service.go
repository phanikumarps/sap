package umc

import (
	"context"
	"io"

	"github.com/phanikumarps/sap/odata"
)

const RootResource = "/sap/opu/odata/sap/ERP_ISU_UMC/"

type Getter interface {
	Get(context.Context, any) (*io.ReadCloser, error)
}

type Inserter interface {
	Insert(context.Context, any) (*io.WriteCloser, error)
}

type Updater interface {
	Update(context.Context, any) (*io.WriteCloser, error)
}

type Deleter interface {
	Delete(context.Context, any) (*io.WriteCloser, error)
}

type host odata.Host

type resource struct {
	Resource     odata.Resource
	RootResource string
}
