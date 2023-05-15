package umc

const RootResource = "/sap/opu/odata/sap/ERP_ISU_UMC/"

/*
type Service struct {
	httpclient.Client
}

type Handler[T any] struct{}

func (handler *Handler[T]) Get(instance T, clnt *httpclient.Client, resource string) {
	ctx := context.TODO()
	resp, err := httpclient.Call(ctx, clnt, http.MethodGet, odata.RootResource, resource, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
}

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
*/
