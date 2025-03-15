package httplib

type HTTPLib interface {
	DoSomething()
}

func NewMockHTTPLib() HTTPLib {
	return &MockHTTPLib{}
}

type MockHTTPLib struct{}

func (m *MockHTTPLib) DoSomething() {

}
