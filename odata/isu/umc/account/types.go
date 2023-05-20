package account

type AcctOutput struct {
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
}
