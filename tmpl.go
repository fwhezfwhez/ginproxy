package util

type errnoTmpl struct {
	ErrnoFiledName  string
	ErrmsgFieldName string
}

func newErrnoTmpl(errnoFieldName string, errmsgFieldName string) errnoTmpl {
	return errnoTmpl{
		ErrnoFiledName:  errnoFieldName,
		ErrmsgFieldName: errmsgFieldName,
	}
}

var defaultErrnoTmpl = newErrnoTmpl("errno", "errmsg")
