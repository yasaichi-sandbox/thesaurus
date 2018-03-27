package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BigHuge struct {
	APIKey string
}

type words struct {
	Syn []string `json:"syn"`
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

const apiEndPoint = "http://words.bighugelabs.com/api/2/%s/%s/json"

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string

	url := fmt.Sprintf(apiEndPoint, b.APIKey, term)
	response, err := http.Get(url)
	if err != nil {
		return syns, fmt.Errorf("bighuge: %qの類語検索に失敗しました: %v", term, err)
	}
	defer response.Body.Close()

	var data synonyms
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}

	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}

	return syns, nil
}
