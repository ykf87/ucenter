package models

type LanguageModel struct {
	Id   int64  `json:"id"`
	Iso  string `json:"iso"`
	Name string `json:"name"`
}

var LangLists map[string]*LanguageModel

func GetAllLanguages(reget bool) (list map[string]*LanguageModel, err error) {
	if LangLists != nil && reget == false {
		list = LangLists
		return
	}
	var langs []*LanguageModel
	rs := DB.Table("languages").Where("status = 1").Order("sort DESC").Find(&langs)
	if rs.Error != nil {
		err = rs.Error
		return
	}
	list = make(map[string]*LanguageModel)
	for _, v := range langs {
		list[v.Iso] = v
	}
	LangLists = list
	return
}
