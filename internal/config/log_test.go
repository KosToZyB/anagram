package config

import "testing"

func TestLog_GetLevel(t *testing.T) {
	t.Parallel()
	value := "DEBUG"
	err := SetEnv(logLevel, value)
	if err != nil {
		t.Fatalf("Error set env %s=%s\n", logLevel, value)
	}

	l := Log{}
	level := l.GetLevel()
	if level != value {
		t.Fatalf("Wrong read %s. Expected %s but found %s\n", logLevel, value, level)
	}
}
