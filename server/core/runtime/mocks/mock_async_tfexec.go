// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime (interfaces: AsyncTFExec)

package mocks

import (
	go_version "github.com/hashicorp/go-version"
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/core/runtime/models"
	command "github.com/runatlantis/atlantis/server/events/command"
	"reflect"
	"time"
)

type MockAsyncTFExec struct {
	fail func(message string, callerSkip ...int)
}

func NewMockAsyncTFExec(options ...pegomock.Option) *MockAsyncTFExec {
	mock := &MockAsyncTFExec{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockAsyncTFExec) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockAsyncTFExec) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockAsyncTFExec) RunCommandAsync(_param0 command.ProjectContext, _param1 string, _param2 []string, _param3 map[string]string, _param4 *go_version.Version, _param5 string) (chan<- string, <-chan models.Line) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockAsyncTFExec().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4, _param5}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RunCommandAsync", params, []reflect.Type{reflect.TypeOf((*chan<- string)(nil)).Elem(), reflect.TypeOf((*<-chan models.Line)(nil)).Elem()})
	var ret0 chan<- string
	var ret1 <-chan models.Line
	if len(result) != 0 {
		if result[0] != nil {
			var ok bool
			ret0, ok = result[0].(chan string)
			if !ok {
				ret0 = result[0].(chan<- string)
			}
		}
		if result[1] != nil {
			var ok bool
			ret1, ok = result[1].(chan models.Line)
			if !ok {
				ret1 = result[1].(<-chan models.Line)
			}
		}
	}
	return ret0, ret1
}

func (mock *MockAsyncTFExec) VerifyWasCalledOnce() *VerifierMockAsyncTFExec {
	return &VerifierMockAsyncTFExec{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockAsyncTFExec) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockAsyncTFExec {
	return &VerifierMockAsyncTFExec{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockAsyncTFExec) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockAsyncTFExec {
	return &VerifierMockAsyncTFExec{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockAsyncTFExec) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockAsyncTFExec {
	return &VerifierMockAsyncTFExec{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockAsyncTFExec struct {
	mock                   *MockAsyncTFExec
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockAsyncTFExec) RunCommandAsync(_param0 command.ProjectContext, _param1 string, _param2 []string, _param3 map[string]string, _param4 *go_version.Version, _param5 string) *MockAsyncTFExec_RunCommandAsync_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4, _param5}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RunCommandAsync", params, verifier.timeout)
	return &MockAsyncTFExec_RunCommandAsync_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockAsyncTFExec_RunCommandAsync_OngoingVerification struct {
	mock              *MockAsyncTFExec
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockAsyncTFExec_RunCommandAsync_OngoingVerification) GetCapturedArguments() (command.ProjectContext, string, []string, map[string]string, *go_version.Version, string) {
	_param0, _param1, _param2, _param3, _param4, _param5 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1], _param4[len(_param4)-1], _param5[len(_param5)-1]
}

func (c *MockAsyncTFExec_RunCommandAsync_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext, _param1 []string, _param2 [][]string, _param3 []map[string]string, _param4 []*go_version.Version, _param5 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]command.ProjectContext, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(command.ProjectContext)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([][]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.([]string)
		}
		_param3 = make([]map[string]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(map[string]string)
		}
		_param4 = make([]*go_version.Version, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(*go_version.Version)
		}
		_param5 = make([]string, len(c.methodInvocations))
		for u, param := range params[5] {
			_param5[u] = param.(string)
		}
	}
	return
}
