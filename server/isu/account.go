package isu

// type resource struct {
// 	server.Server
// }

// func (r *resource) handleAccount() http.HandlerFunc {
// 	// any closure logic here
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		h := "https://saptq1.fpl.com" + "443"
// 		u := umc.NewService(h, "token")
// 		ctx := context.TODO()
// 		var accountID string

// 		resp, err := u.GetAccount(ctx, accountID)
// 		if err != nil {
// 			log.Err(err)
// 		}

// 		err = json.NewEncoder(w).Encode(resp)
// 		if err != nil {
// 			log.Err(err)
// 		}

// 	}
// }

// func (r *resource) handleContractAccount() http.HandlerFunc {
// 	// any closure logic here
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }

// func (r *resource) handlePremise() http.HandlerFunc {
// 	// any closure logic here
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }
