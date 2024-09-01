package speller

import (
	"encoding/json"
	"fmt"
	"io"
	"kode-notes/internal/models"
	"log"
	"net/http"
)

const API = "https://speller.yandex.net/services/spellservice.json/checkText?text="

type Spell struct {
	cli *http.Client
}

// ...
func NewSpell(cli *http.Client) *Spell {
	return &Spell{cli: cli}
}

// ...
func (r *Spell) CheckText(text string) (errSpeller []*models.Spell, err error) {
	log.Println("CheckText speller was invoked")

	response, err := http.Get(API + text)
	if err != nil {
		return nil, fmt.Errorf("can't make a reguest to Yandex.Speller API: %v", err)
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	errSpeller = make([]*models.Spell, 10)

	err = json.Unmarshal(content, &errSpeller)
	if err != nil {
		return nil, fmt.Errorf("can't spell marshal to json: %v", err)
	}

	return errSpeller, err
}
