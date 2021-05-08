package main

import "github.com/mitchellh/hashstructure/v2"

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Zip     string `json:"zip"`
	Id      string `json:"id"`
}

// Includable interface for calculating hash ignoring id
func (u User) HashInclude(field string, v interface{}) (bool, error) {
	if field == "Id" {
		return false, nil
	}
	return true, nil
}

func (u User) Hash() uint64 {
	hash, err := hashstructure.Hash(u, hashstructure.FormatV2, nil)
	if err != nil {
		panic(err)
	}
	return hash
}

func (u User) DoesContainEmptyValues() bool {
	return u.Address == "" || u.Zip == "" || u.Name == ""
}
