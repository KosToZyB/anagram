package anagram

import (
	"sort"
	"strings"
)

type Anagram struct {
	data map[string][]string
}

func NewAnagram() *Anagram {
	a := &Anagram{
		data: make(map[string][]string, 0),
	}

	return a
}

func (a *Anagram) Load(array []string) {
	for i := range array {
		key := a.getKey(array[i])
		a.addIfUnique(key, array[i])
	}
}

func (a *Anagram) getKey(s string) string {
	l := strings.ToLower(s)
	array := strings.Split(l, "")
	sort.Strings(array)

	return strings.Join(array, "")
}

func (a *Anagram) addIfUnique(key, value string) {
	array, ok := a.data[key]
	if !ok {
		a.data[key] = append([]string{}, value)
		return
	}

	for _, v := range array {
		if v == value {
			return
		}
	}

	array = append(array, value)
	sort.Strings(array)
	a.data[key] = array
}

func (a *Anagram) GetAnagram(word string) []string {
	key := a.getKey(word)
	array, ok := a.data[key]
	if ok {
		return array
	}

	return []string{}
}