package core

import (
	"testing"
)

func TestOnSuccess(t *testing.T) {
	context := &Context{
		Stats:        Stats{InvocationCount: 0, SuccessCount: 0, FailedCount: 0},
		DynamoClient: nil}
	context.OnSuccess()
	if count := context.GetSuccessCount(); count != 1 {
		t.Errorf("Invalid success count: %v", count)
	}
	if count := context.GetInvocationCount(); count != 1 {
		t.Errorf("Invalid invocation count: %v", count)
	}
}

func TestOnFailure(t *testing.T) {
	context := &Context{
		Stats:        Stats{InvocationCount: 0, SuccessCount: 0, FailedCount: 0},
		DynamoClient: nil}
	context.OnFailure()
	if count := context.GetFailedCount(); count != 1 {
		t.Errorf("Invalid failed count: %v", count)
	}
	if count := context.GetInvocationCount(); count != 1 {
		t.Errorf("Invalid invocation count: %v", count)
	}
}
