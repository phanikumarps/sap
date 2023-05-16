package config

type Host struct {
	address string
}

func NewHost(host string) *Host {

	return &Host{
		address: host,
	}
}

type RootPath string

type Resource struct {
	name      string
	host      Host
	rootPath  string
	sapClient string
	authToken string
}

func NewResource(host Host, name, rootPath, sapClient, authToken string) *Resource {
	return &Resource{
		name:      name,
		host:      host,
		rootPath:  rootPath,
		sapClient: sapClient,
		authToken: authToken,
	}
}

func DefaultRootPath(service string) *RootPath {
	if service == "" {
		r := RootPath("/sap/opu/odata/sap/")
		return &r
	}
	r := RootPath("/sap/opu/odata/sap/" + service + "/")
	return &r
}
