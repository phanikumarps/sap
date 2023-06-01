package isu

func (r *resource) routes() {
	r.Router.HandleFunc("/account/%s", r.handleAccount())
	r.Router.HandleFunc("/contractaccount/%s", r.handleContractAccount())
	r.Router.HandleFunc("/premise/%s", r.handleContractAccount())
}
