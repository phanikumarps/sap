package account

type acctResponse struct {
	D struct {
		Metadata struct {
			ID   string `json:"id"`
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		AccountTypeID             string `json:"AccountTypeID"`
		AccountID                 string `json:"AccountID"`
		AccountTitleID            string `json:"AccountTitleID"`
		FirstName                 string `json:"FirstName"`
		LastName                  string `json:"LastName"`
		MiddleName                string `json:"MiddleName"`
		SecondName                string `json:"SecondName"`
		Sex                       string `json:"Sex"`
		Name1                     string `json:"Name1"`
		Name2                     string `json:"Name2"`
		Name3                     string `json:"Name3"`
		Name4                     string `json:"Name4"`
		GroupName1                string `json:"GroupName1"`
		GroupName2                string `json:"GroupName2"`
		FullName                  string `json:"FullName"`
		CorrespondenceLanguage    string `json:"CorrespondenceLanguage"`
		CorrespondenceLanguageISO string `json:"CorrespondenceLanguageISO"`
		Language                  string `json:"Language"`
		LanguageISO               string `json:"LanguageISO"`
		AccountRelationships      struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountRelationships"`
		AccountBalance struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountBalance"`
		AccountAddresses struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountAddresses"`
		ContractAccounts struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"ContractAccounts"`
		PaymentCards struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"PaymentCards"`
		Correspondences struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"Correspondences"`
		CommunicationPreferences struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"CommunicationPreferences"`
		AccountContacts struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountContacts"`
		AccountType struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountType"`
		AccountAlerts struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountAlerts"`
		AccountAddressIndependentEmails struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountAddressIndependentEmails"`
		StandardAccountAddress struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"StandardAccountAddress"`
		AccountAddressIndependentMobilePhones struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountAddressIndependentMobilePhones"`
		AccountAddressIndependentPhones struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountAddressIndependentPhones"`
		AccountAddressIndependentFaxes struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountAddressIndependentFaxes"`
		PaymentDocuments struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"PaymentDocuments"`
		BankAccounts struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"BankAccounts"`
		AccountSex struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountSex"`
		AccountTitle struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"AccountTitle"`
		ServiceNotifications struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"ServiceNotifications"`
		ServiceOrders struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"ServiceOrders"`
		Outages struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"Outages"`
		Invoices struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"Invoices"`
		InteractionRecords struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"InteractionRecords"`
	} `json:"d"`
}
