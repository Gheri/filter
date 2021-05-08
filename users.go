package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pkg/errors"
)

type Users []User

// create new user from json file name
func NewUsers(filename string) (Users, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to open Json File.")
	}
	defer file.Close()

	// Todo use streams/pipe
	dataInBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "Error occurred in reading bytes from file.")
	}

	var users Users

	err = json.Unmarshal(dataInBytes, &users)
	if err != nil {
		return nil, errors.Wrap(err, "Error occurred in unmarshing json.")
	}

	return users, nil
}

func (list Users) getIdsWithIdenticalValues() []string {
	hashVsUserMap := make(map[uint64][]string, len(list))
	for _, user := range list {
		// not considering user ids with empty values for hash map
		if user.DoesContainEmptyValues() {
			continue
		}
		hash := user.Hash()
		ids, ok := hashVsUserMap[hash]
		// if already existing collate all ids
		if ok {
			ids = append(ids, user.Id)
		} else {
			ids = []string{user.Id}
		}
		hashVsUserMap[hash] = ids
	}
	idsWithIdenticalValues := []string{}
	for _, ids := range hashVsUserMap {
		if len(ids) > 1 {
			idsWithIdenticalValues = append(idsWithIdenticalValues, ids...)
		}
	}
	return idsWithIdenticalValues
}

func (list Users) getIdsWithEmptyValues() []string {
	idsWithEmptyValues := []string{}
	for _, user := range list {
		if user.DoesContainEmptyValues() {
			idsWithEmptyValues = append(idsWithEmptyValues, user.Id)
		}
	}
	return idsWithEmptyValues
}

// returns ids with identical records and users with missing/empty values
func (list Users) GetInvalidIds() []string {
	invalidIds := []string{}
	idsWithIdenticalValues := list.getIdsWithIdenticalValues()
	invalidIds = append(invalidIds, idsWithIdenticalValues...)
	idsWithEmptyValues := list.getIdsWithEmptyValues()
	invalidIds = append(invalidIds, idsWithEmptyValues...)
	return invalidIds
}
