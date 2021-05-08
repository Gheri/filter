package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUsers(t *testing.T) {
	expectedValidData := Users{
		{
			"Viennia Sturm",
			"3434 Street 27",
			"17565",
			"ea0c4",
		},
		{
			"Amerah Lang",
			"5037 Providence Bouled",
			"44109",
			"8d322",
		},
		{
			"Wendolyn Sweat",
			"1521 Gem Avenue",
			"77701",
			"a6b3b",
		},
	}
	expectedValidDataWithEmptyValues := Users{
		{
			"Viennia Sturm",
			"3434 Street 27",
			"",
			"ea0c4",
		},
		{
			"Amerah Lang",
			"",
			"44109",
			"8d322",
		},
		{
			"",
			"1521 Gem Avenue",
			"77701",
			"a6b3b",
		},
	}
	t.Run("users created with valid data", func(t *testing.T) {
		users, err := NewUsers("./test_data/valid_data.json")
		assert.NoError(t, err)
		assert.Equal(t, expectedValidData, users)
	})
	t.Run("users created with empty data values", func(t *testing.T) {
		users, err := NewUsers("./test_data/valid_data_empty_values.json")
		assert.NoError(t, err)
		assert.Equal(t, expectedValidDataWithEmptyValues, users)
	})
	t.Run("users created with null data values", func(t *testing.T) {
		users, err := NewUsers("./test_data/valid_data_null_values.json")
		assert.NoError(t, err)
		assert.Equal(t, expectedValidDataWithEmptyValues, users)
	})
	t.Run("users created with missing data values", func(t *testing.T) {
		users, err := NewUsers("./test_data/valid_data_missing_values.json")
		assert.NoError(t, err)
		assert.Equal(t, expectedValidDataWithEmptyValues, users)
	})
	t.Run("users created with valid data ignores extra fields in json", func(t *testing.T) {
		users, err := NewUsers("./test_data/valid_data.json")
		assert.NoError(t, err)
		assert.Equal(t, expectedValidData, users)
	})
	t.Run("users creation failed for text file", func(t *testing.T) {
		users, err := NewUsers("./test_data/file.txt")
		assert.Error(t, err)
		assert.Nil(t, users)
	})
	t.Run("users creation failed for invalid json", func(t *testing.T) {
		users, err := NewUsers("./test_data/invalid.json")
		assert.Error(t, err)
		assert.Nil(t, users)
	})
	t.Run("users creation failed for nonexisting file", func(t *testing.T) {
		users, err := NewUsers("./test_data/non_existing.json")
		assert.Error(t, err)
		assert.Nil(t, users)
	})
}

func TestInvalidUsers(t *testing.T) {
	t.Run("test with identical values", func(t *testing.T) {
		usersWithIdenticalValues := Users{
			{
				"Viennia Sturm",
				"3434 Street 27",
				"17565",
				"ea0c4",
			},
			{
				"Viennia Sturm",
				"3434 Street 27",
				"17565",
				"ea0c5",
			},
		}
		got := usersWithIdenticalValues.GetInvalidIds()
		want := []string{"ea0c4", "ea0c5"}
		assert.Equal(t, got, want)
	})
	t.Run("test with missing/empty name values", func(t *testing.T) {
		usersWithNameMissingValues := Users{
			{
				"",
				"3434 Street 27",
				"17565",
				"ea0c5",
			},
		}
		got := usersWithNameMissingValues.GetInvalidIds()
		want := []string{"ea0c5"}
		assert.Equal(t, got, want)
	})
	t.Run("test with missing/empty zip values", func(t *testing.T) {
		usersWithZipMissingValues := Users{
			{
				"Name",
				"3434 Street 27",
				"",
				"ea0c5",
			},
		}
		got := usersWithZipMissingValues.GetInvalidIds()
		want := []string{"ea0c5"}
		assert.Equal(t, got, want)
	})
	t.Run("test with missing/empty address values", func(t *testing.T) {
		usersWithAddressMissingValues := Users{
			{
				"Name",
				"",
				"12345",
				"ea0c5",
			},
		}
		got := usersWithAddressMissingValues.GetInvalidIds()
		want := []string{"ea0c5"}
		assert.Equal(t, got, want)
	})
	t.Run("test with some identical records and some missing values", func(t *testing.T) {
		usersWithSomeIdenticalAndSomeEmptyValues := Users{
			{
				"Viennia Sturm",
				"3434 Street 27",
				"17565",
				"ea0c4",
			},
			{
				"Viennia Sturm",
				"3434 Street 27",
				"17565",
				"ea0c5",
			},
			{
				"Sturm",
				"Street 27",
				"27565",
				"ea0c6",
			},
			{
				"",
				"",
				"17565",
				"ea0c7",
			},
			{
				"Viennia Sturm",
				"3434 Street 27",
				"",
				"ea0c8",
			},
		}
		got := usersWithSomeIdenticalAndSomeEmptyValues.GetInvalidIds()
		want := []string{"ea0c4", "ea0c5", "ea0c7", "ea0c8"}
		assert.Equal(t, got, want)
	})
}
