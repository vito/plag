// This file was generated by counterfeiter
package outputfakes

import (
	"sync"

	"github.com/tscolari/plag/output"
	"github.com/tscolari/plag/parser"
)

type FakeOutputer struct {
	WriteStub        func(data parser.Data) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		data parser.Data
	}
	writeReturns struct {
		result1 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOutputer) Write(data parser.Data) error {
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		data parser.Data
	}{data})
	fake.recordInvocation("Write", []interface{}{data})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(data)
	} else {
		return fake.writeReturns.result1
	}
}

func (fake *FakeOutputer) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeOutputer) WriteArgsForCall(i int) parser.Data {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].data
}

func (fake *FakeOutputer) WriteReturns(result1 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputer) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeOutputer) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeOutputer) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOutputer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeOutputer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ output.Outputer = new(FakeOutputer)
