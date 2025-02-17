// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: PendingPlanFinder)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
	"reflect"
	"time"
)

type MockPendingPlanFinder struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPendingPlanFinder(options ...pegomock.Option) *MockPendingPlanFinder {
	mock := &MockPendingPlanFinder{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockPendingPlanFinder) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockPendingPlanFinder) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockPendingPlanFinder) DeletePlans(_param0 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPendingPlanFinder().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DeletePlans", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockPendingPlanFinder) Find(_param0 string) ([]events.PendingPlan, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPendingPlanFinder().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Find", params, []reflect.Type{reflect.TypeOf((*[]events.PendingPlan)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []events.PendingPlan
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]events.PendingPlan)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockPendingPlanFinder) VerifyWasCalledOnce() *VerifierMockPendingPlanFinder {
	return &VerifierMockPendingPlanFinder{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockPendingPlanFinder) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockPendingPlanFinder {
	return &VerifierMockPendingPlanFinder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockPendingPlanFinder) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockPendingPlanFinder {
	return &VerifierMockPendingPlanFinder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockPendingPlanFinder) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockPendingPlanFinder {
	return &VerifierMockPendingPlanFinder{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockPendingPlanFinder struct {
	mock                   *MockPendingPlanFinder
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockPendingPlanFinder) DeletePlans(_param0 string) *MockPendingPlanFinder_DeletePlans_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DeletePlans", params, verifier.timeout)
	return &MockPendingPlanFinder_DeletePlans_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockPendingPlanFinder_DeletePlans_OngoingVerification struct {
	mock              *MockPendingPlanFinder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockPendingPlanFinder_DeletePlans_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockPendingPlanFinder_DeletePlans_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockPendingPlanFinder) Find(_param0 string) *MockPendingPlanFinder_Find_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Find", params, verifier.timeout)
	return &MockPendingPlanFinder_Find_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockPendingPlanFinder_Find_OngoingVerification struct {
	mock              *MockPendingPlanFinder
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockPendingPlanFinder_Find_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockPendingPlanFinder_Find_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}
