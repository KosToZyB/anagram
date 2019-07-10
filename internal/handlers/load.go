package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (a *API) load(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		a.log.Errorf("wrong copy bytes from request: %s", err)
		return
	}

	var words []string
	err = json.Unmarshal(b ,&words)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		a.log.Errorf("error unmarshal request: %s", err)
		return
	}

	a.anagram.Load(words)
}
