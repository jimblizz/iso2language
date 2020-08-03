package isotolanguage

import (
	"github.com/gocarina/gocsv"
	"strings"
)

type (
	Language struct {
		Iso2 string `csv:"iso2"`
		Iso3 string `csv:"iso3"`
		EnglishName string `csv:"english_name"`
		FrenchName string `csv:"french_name"`
		GermanName string `csv:"german_name"`
	}

	Service struct {
		indexIso2 map[string]Language
		indexIso3 map[string]Language
	}
)

func New() (service Service, err error) {
	service.indexIso2 = make(map[string]Language, 0)
	service.indexIso3 = make(map[string]Language, 0)

	languages := make([]Language, 0)
	if err = gocsv.UnmarshalString(Data, &languages); err != nil {
		return
	}

	for _, l := range languages {
		if l.Iso2 != "" {
			for _, iso2 := range strings.Split(l.Iso2, "/") {
				service.indexIso2[iso2] = l
			}
		}

		if l.Iso3 != "" {
			for _, iso3 := range strings.Split(l.Iso3, "/") {
				service.indexIso3[iso3] = l
			}
		}
	}

	return
}

func (s Service) FindFromIso(iso string) (language Language) {

	if val, ok := s.indexIso2[iso]; ok {
		return val
	}

	if val, ok := s.indexIso3[iso]; ok {
		return val
	}

	return
}
