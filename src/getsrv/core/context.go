package core

import (
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Stats maintained
type Stats struct {
	InvocationCount uint32
	SuccessCount    uint32
	FailedCount     uint32
}

// Global Context
type Context struct {
	DynamoClient *dynamodb.DynamoDB
	Stats
}

// Create new Context
// Credentials from the shared credentials file ~/.aws/credentials
// Region from the shared configuration file ~/.aws/config.
// The dynamo client is attached to this context
func NewContext() *Context {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable}))

	// Create DynamoDB client
	context := &Context{
		Stats:        Stats{InvocationCount: 0, SuccessCount: 0, FailedCount: 0},
		DynamoClient: dynamodb.New(sess)}

	return context
}

func (context *Context) GetInvocationCount() uint32 {
	return atomic.LoadUint32(&context.InvocationCount)
}

func (context *Context) GetSuccessCount() uint32 {
	return atomic.LoadUint32(&context.SuccessCount)
}

func (context *Context) GetFailedCount() uint32 {
	return atomic.LoadUint32(&context.FailedCount)
}

func (context *Context) OnSuccess() {
	atomic.AddUint32(&context.InvocationCount, 1)
	atomic.AddUint32(&context.SuccessCount, 1)
}

func (context *Context) OnFailure() {
	atomic.AddUint32(&context.InvocationCount, 1)
	atomic.AddUint32(&context.FailedCount, 1)
}
