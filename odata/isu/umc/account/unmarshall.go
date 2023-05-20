package account

func ConvRespToOutput(resp *AcctResponse) *AcctOutput {
	acctop := AcctOutput{
		AccountTypeID:             resp.D.AccountTypeID,
		AccountID:                 resp.D.AccountID,
		AccountTitleID:            resp.D.AccountTitleID,
		FirstName:                 resp.D.FirstName,
		LastName:                  resp.D.LastName,
		MiddleName:                resp.D.MiddleName,
		SecondName:                resp.D.SecondName,
		Sex:                       resp.D.Sex,
		Name1:                     resp.D.Name1,
		Name2:                     resp.D.Name2,
		Name3:                     resp.D.Name3,
		Name4:                     resp.D.Name4,
		GroupName1:                resp.D.GroupName1,
		GroupName2:                resp.D.GroupName2,
		FullName:                  resp.D.FullName,
		CorrespondenceLanguage:    resp.D.CorrespondenceLanguage,
		CorrespondenceLanguageISO: resp.D.CorrespondenceLanguageISO,
		Language:                  resp.D.Language,
		LanguageISO:               resp.D.LanguageISO,
	}
	return &acctop
}
