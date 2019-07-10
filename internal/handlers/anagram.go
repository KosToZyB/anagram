package handlers

import (
	"encoding/json"
	"net/http"
)

func (a *API) getAnagram(w http.ResponseWriter, r *http.Request) {
	word, ok := r.URL.Query()["word"]

	if !ok || len(word[0]) < 1 {
		a.log.Error("bad request. Param 'word' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	array := a.anagram.GetAnagram(word[0])
	if len(array) == 0 {
		array = nil
	}

	b, err := json.Marshal(array)
	if err != nil {
		a.log.Errorf("error marshal array of string to array byte: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		a.log.Errorf("error write byte to response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
