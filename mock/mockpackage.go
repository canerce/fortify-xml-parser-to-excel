package mock

import (
	"fmt"
	"io"
)

//Package mock of io.Closer and other information for testing
type Package struct {
	io.Closer
	SaveCalled      bool
	SaveAsInterface interface{}
	ForceError      bool
}

//NewPackage creates new package
func NewPackage() *Package {
	return &Package{}
}

//Save - mocked
func (m *Package) Save() error {
	if m.ForceError {
		return fmt.Errorf("error forced")
	}
	m.SaveCalled = true
	return nil
}

//SaveAs - mocked
func (m *Package) SaveAs(target interface{}) error {
	if m.ForceError {
		return fmt.Errorf("error forced")
	}
	m.SaveAsInterface = target
	return nil
}
