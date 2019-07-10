package config

import "testing"

func TestSetGetEnv(t *testing.T) {
	t.Parallel()
	key := "test_key"
	value := "test_value"
	err := SetEnv(key, value)
	if err != nil {
		t.Fatalf("Error set env %s=%s\n", key, value)
	}

	result := GetEnv(key)
	if result != value {
		t.Fatalf("Wrong read %s. Expected %s but found %s\n", key, value, result)
	}
}
