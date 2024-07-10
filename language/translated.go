package language

type TranslatedLanguages struct {
	Vietnamese string `json:"vi"` // Vietnamese
}

func (l TranslatedLanguages) GetLanguageValue(lang string) string {
	dLang := ToLanguage(lang)
	if dLang.IsVietnamese() {
		return l.Vietnamese
	}

	return l.Vietnamese
}
