package logger

import "fmt"

func GetMockLogger() *MockLog {
	return &MockLog{}
}

type MockLog struct {
	data []string
}

func (m *MockLog) Debug(args ...interface{}) {
	m.append(args...)
}

func (m *MockLog) Info(args ...interface{}) {
	m.append(args...)
}

func (m MockLog) Warn(args ...interface{}) {
	m.append(args...)
}

func (m *MockLog) Error(args ...interface{}) {
	m.append(args...)
}

func (m *MockLog) Panic(args ...interface{}) {
	m.append(args...)
}

func (m *MockLog) Fatal(args ...interface{}) {
	m.append(args...)
}

func (m *MockLog) Debugf(template string, args ...interface{}) {
	m.appendf(template, args...)
}

func (m *MockLog) Infof(template string, args ...interface{}) {
	m.appendf(template, args...)
}

func (m *MockLog) Warnf(template string, args ...interface{}) {
	m.appendf(template, args...)
}

func (m *MockLog) Errorf(template string, args ...interface{}) {
	m.appendf(template, args...)
}

func (m *MockLog) Panicf(template string, args ...interface{}) {
	m.appendf(template, args...)
}

func (m *MockLog) Fatalf(template string, args ...interface{}) {
	m.appendf(template, args...)
}

func (m *MockLog) append(args ...interface{}) {
	line := fmt.Sprint(args...)
	m.data = append(m.data, line)
}

func (m *MockLog) appendf(template string, args ...interface{}) {
	line := fmt.Sprintf(template, args...)
	m.data = append(m.data, line)
}

func (m *MockLog) GetLogs() []string {
	return m.data
}
