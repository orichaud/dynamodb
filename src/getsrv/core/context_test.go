package core

import (
	"testing"
)

func TestOnSuccess(t *testing.T) {
	context := &Context{
		Stats:        Stats{InvocationCount: 0, SuccessCount: 0, FailedCount: 0},
		DynamoClient: nil}
	context.OnSuccess()
	stats := context.GetStats()
	if stats.SuccessCount != 1 {
		t.Errorf("Invalid success count: %v", stats.SuccessCount)
	}
	if stats.InvocationCount != 1 {
		t.Errorf("Invalid invocation count: %v", stats.InvocationCount)
	}
}

func TestOnFailure(t *testing.T) {
	context := &Context{
		Stats:        Stats{InvocationCount: 0, SuccessCount: 0, FailedCount: 0},
		DynamoClient: nil}
	context.OnFailure()
	stats := context.GetStats()
	if stats.FailedCount != 1 {
		t.Errorf("Invalid failed count: %v", stats.FailedCount)
	}
	if stats.InvocationCount != 1 {
		t.Errorf("Invalid invocation count: %v", stats.InvocationCount)
	}
}
