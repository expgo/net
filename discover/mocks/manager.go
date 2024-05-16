// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/expgo/net/discover"
	"github.com/expgo/net/protocol"
)

type Manager struct {
	CacheStub        func() map[protocol.DeviceID]discover.CacheEntry
	cacheMutex       sync.RWMutex
	cacheArgsForCall []struct {
	}
	cacheReturns struct {
		result1 map[protocol.DeviceID]discover.CacheEntry
	}
	cacheReturnsOnCall map[int]struct {
		result1 map[protocol.DeviceID]discover.CacheEntry
	}
	ChildErrorsStub        func() map[string]error
	childErrorsMutex       sync.RWMutex
	childErrorsArgsForCall []struct {
	}
	childErrorsReturns struct {
		result1 map[string]error
	}
	childErrorsReturnsOnCall map[int]struct {
		result1 map[string]error
	}
	ErrorStub        func() error
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
	}
	errorReturns struct {
		result1 error
	}
	errorReturnsOnCall map[int]struct {
		result1 error
	}
	LookupStub        func(context.Context, protocol.DeviceID) ([]string, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		arg1 context.Context
		arg2 protocol.DeviceID
	}
	lookupReturns struct {
		result1 []string
		result2 error
	}
	lookupReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ServeStub        func(context.Context) error
	serveMutex       sync.RWMutex
	serveArgsForCall []struct {
		arg1 context.Context
	}
	serveReturns struct {
		result1 error
	}
	serveReturnsOnCall map[int]struct {
		result1 error
	}
	StringStub        func() string
	stringMutex       sync.RWMutex
	stringArgsForCall []struct {
	}
	stringReturns struct {
		result1 string
	}
	stringReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Manager) Cache() map[protocol.DeviceID]discover.CacheEntry {
	fake.cacheMutex.Lock()
	ret, specificReturn := fake.cacheReturnsOnCall[len(fake.cacheArgsForCall)]
	fake.cacheArgsForCall = append(fake.cacheArgsForCall, struct {
	}{})
	stub := fake.CacheStub
	fakeReturns := fake.cacheReturns
	fake.recordInvocation("Cache", []interface{}{})
	fake.cacheMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Manager) CacheCallCount() int {
	fake.cacheMutex.RLock()
	defer fake.cacheMutex.RUnlock()
	return len(fake.cacheArgsForCall)
}

func (fake *Manager) CacheCalls(stub func() map[protocol.DeviceID]discover.CacheEntry) {
	fake.cacheMutex.Lock()
	defer fake.cacheMutex.Unlock()
	fake.CacheStub = stub
}

func (fake *Manager) CacheReturns(result1 map[protocol.DeviceID]discover.CacheEntry) {
	fake.cacheMutex.Lock()
	defer fake.cacheMutex.Unlock()
	fake.CacheStub = nil
	fake.cacheReturns = struct {
		result1 map[protocol.DeviceID]discover.CacheEntry
	}{result1}
}

func (fake *Manager) CacheReturnsOnCall(i int, result1 map[protocol.DeviceID]discover.CacheEntry) {
	fake.cacheMutex.Lock()
	defer fake.cacheMutex.Unlock()
	fake.CacheStub = nil
	if fake.cacheReturnsOnCall == nil {
		fake.cacheReturnsOnCall = make(map[int]struct {
			result1 map[protocol.DeviceID]discover.CacheEntry
		})
	}
	fake.cacheReturnsOnCall[i] = struct {
		result1 map[protocol.DeviceID]discover.CacheEntry
	}{result1}
}

func (fake *Manager) ChildErrors() map[string]error {
	fake.childErrorsMutex.Lock()
	ret, specificReturn := fake.childErrorsReturnsOnCall[len(fake.childErrorsArgsForCall)]
	fake.childErrorsArgsForCall = append(fake.childErrorsArgsForCall, struct {
	}{})
	stub := fake.ChildErrorsStub
	fakeReturns := fake.childErrorsReturns
	fake.recordInvocation("ChildErrors", []interface{}{})
	fake.childErrorsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Manager) ChildErrorsCallCount() int {
	fake.childErrorsMutex.RLock()
	defer fake.childErrorsMutex.RUnlock()
	return len(fake.childErrorsArgsForCall)
}

func (fake *Manager) ChildErrorsCalls(stub func() map[string]error) {
	fake.childErrorsMutex.Lock()
	defer fake.childErrorsMutex.Unlock()
	fake.ChildErrorsStub = stub
}

func (fake *Manager) ChildErrorsReturns(result1 map[string]error) {
	fake.childErrorsMutex.Lock()
	defer fake.childErrorsMutex.Unlock()
	fake.ChildErrorsStub = nil
	fake.childErrorsReturns = struct {
		result1 map[string]error
	}{result1}
}

func (fake *Manager) ChildErrorsReturnsOnCall(i int, result1 map[string]error) {
	fake.childErrorsMutex.Lock()
	defer fake.childErrorsMutex.Unlock()
	fake.ChildErrorsStub = nil
	if fake.childErrorsReturnsOnCall == nil {
		fake.childErrorsReturnsOnCall = make(map[int]struct {
			result1 map[string]error
		})
	}
	fake.childErrorsReturnsOnCall[i] = struct {
		result1 map[string]error
	}{result1}
}

func (fake *Manager) Error() error {
	fake.errorMutex.Lock()
	ret, specificReturn := fake.errorReturnsOnCall[len(fake.errorArgsForCall)]
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
	}{})
	stub := fake.ErrorStub
	fakeReturns := fake.errorReturns
	fake.recordInvocation("Error", []interface{}{})
	fake.errorMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Manager) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *Manager) ErrorCalls(stub func() error) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = stub
}

func (fake *Manager) ErrorReturns(result1 error) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = nil
	fake.errorReturns = struct {
		result1 error
	}{result1}
}

func (fake *Manager) ErrorReturnsOnCall(i int, result1 error) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = nil
	if fake.errorReturnsOnCall == nil {
		fake.errorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.errorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Manager) Lookup(arg1 context.Context, arg2 protocol.DeviceID) ([]string, error) {
	fake.lookupMutex.Lock()
	ret, specificReturn := fake.lookupReturnsOnCall[len(fake.lookupArgsForCall)]
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		arg1 context.Context
		arg2 protocol.DeviceID
	}{arg1, arg2})
	stub := fake.LookupStub
	fakeReturns := fake.lookupReturns
	fake.recordInvocation("Lookup", []interface{}{arg1, arg2})
	fake.lookupMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Manager) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *Manager) LookupCalls(stub func(context.Context, protocol.DeviceID) ([]string, error)) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = stub
}

func (fake *Manager) LookupArgsForCall(i int) (context.Context, protocol.DeviceID) {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	argsForCall := fake.lookupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Manager) LookupReturns(result1 []string, result2 error) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *Manager) LookupReturnsOnCall(i int, result1 []string, result2 error) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = nil
	if fake.lookupReturnsOnCall == nil {
		fake.lookupReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.lookupReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *Manager) Serve(arg1 context.Context) error {
	fake.serveMutex.Lock()
	ret, specificReturn := fake.serveReturnsOnCall[len(fake.serveArgsForCall)]
	fake.serveArgsForCall = append(fake.serveArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ServeStub
	fakeReturns := fake.serveReturns
	fake.recordInvocation("Serve", []interface{}{arg1})
	fake.serveMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Manager) ServeCallCount() int {
	fake.serveMutex.RLock()
	defer fake.serveMutex.RUnlock()
	return len(fake.serveArgsForCall)
}

func (fake *Manager) ServeCalls(stub func(context.Context) error) {
	fake.serveMutex.Lock()
	defer fake.serveMutex.Unlock()
	fake.ServeStub = stub
}

func (fake *Manager) ServeArgsForCall(i int) context.Context {
	fake.serveMutex.RLock()
	defer fake.serveMutex.RUnlock()
	argsForCall := fake.serveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Manager) ServeReturns(result1 error) {
	fake.serveMutex.Lock()
	defer fake.serveMutex.Unlock()
	fake.ServeStub = nil
	fake.serveReturns = struct {
		result1 error
	}{result1}
}

func (fake *Manager) ServeReturnsOnCall(i int, result1 error) {
	fake.serveMutex.Lock()
	defer fake.serveMutex.Unlock()
	fake.ServeStub = nil
	if fake.serveReturnsOnCall == nil {
		fake.serveReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.serveReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Manager) String() string {
	fake.stringMutex.Lock()
	ret, specificReturn := fake.stringReturnsOnCall[len(fake.stringArgsForCall)]
	fake.stringArgsForCall = append(fake.stringArgsForCall, struct {
	}{})
	stub := fake.StringStub
	fakeReturns := fake.stringReturns
	fake.recordInvocation("String", []interface{}{})
	fake.stringMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Manager) StringCallCount() int {
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	return len(fake.stringArgsForCall)
}

func (fake *Manager) StringCalls(stub func() string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = stub
}

func (fake *Manager) StringReturns(result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	fake.stringReturns = struct {
		result1 string
	}{result1}
}

func (fake *Manager) StringReturnsOnCall(i int, result1 string) {
	fake.stringMutex.Lock()
	defer fake.stringMutex.Unlock()
	fake.StringStub = nil
	if fake.stringReturnsOnCall == nil {
		fake.stringReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.stringReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *Manager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cacheMutex.RLock()
	defer fake.cacheMutex.RUnlock()
	fake.childErrorsMutex.RLock()
	defer fake.childErrorsMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	fake.serveMutex.RLock()
	defer fake.serveMutex.RUnlock()
	fake.stringMutex.RLock()
	defer fake.stringMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Manager) recordInvocation(key string, args []interface{}) {
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

var _ discover.Manager = new(Manager)
