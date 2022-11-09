package mock

import (
	"fmt"
	"io"
)

type MockPackage struct {
	io.Closer
	SaveCalled      bool
	SaveAsInterface interface{}
	ForceError      bool
}

func NewMockPackage() *MockPackage {
	return &MockPackage{}
}

func (m *MockPackage) Save() error {
	if m.ForceError {
		return fmt.Errorf("error forced")
	}
	m.SaveCalled = true
	return nil
}

func (m *MockPackage) SaveAs(target interface{}) error {
	if m.ForceError {
		return fmt.Errorf("error forced")
	}
	m.SaveAsInterface = target
	return nil
}
