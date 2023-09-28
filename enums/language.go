package enums

type LanguageEnum int

const (
	LANGUAGE_ENGLISH LanguageEnum = iota
	LANGUAGE_HINDI   LanguageEnum = iota
)

func (val LanguageEnum) ToString() string {
	if val == LANGUAGE_ENGLISH {
		return "English"
	} else if val == LANGUAGE_HINDI {
		return "Hindi"
	}
	return "NA"
}
