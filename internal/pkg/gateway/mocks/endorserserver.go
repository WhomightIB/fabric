// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/hyperledger/fabric-protos-go-apiv2/peer"
)

type EndorserServer struct {
	ProcessProposalStub        func(context.Context, *peer.SignedProposal) (*peer.ProposalResponse, error)
	processProposalMutex       sync.RWMutex
	processProposalArgsForCall []struct {
		arg1 context.Context
		arg2 *peer.SignedProposal
	}
	processProposalReturns struct {
		result1 *peer.ProposalResponse
		result2 error
	}
	processProposalReturnsOnCall map[int]struct {
		result1 *peer.ProposalResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *EndorserServer) ProcessProposal(arg1 context.Context, arg2 *peer.SignedProposal) (*peer.ProposalResponse, error) {
	fake.processProposalMutex.Lock()
	ret, specificReturn := fake.processProposalReturnsOnCall[len(fake.processProposalArgsForCall)]
	fake.processProposalArgsForCall = append(fake.processProposalArgsForCall, struct {
		arg1 context.Context
		arg2 *peer.SignedProposal
	}{arg1, arg2})
	stub := fake.ProcessProposalStub
	fakeReturns := fake.processProposalReturns
	fake.recordInvocation("ProcessProposal", []interface{}{arg1, arg2})
	fake.processProposalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *EndorserServer) ProcessProposalCallCount() int {
	fake.processProposalMutex.RLock()
	defer fake.processProposalMutex.RUnlock()
	return len(fake.processProposalArgsForCall)
}

func (fake *EndorserServer) ProcessProposalCalls(stub func(context.Context, *peer.SignedProposal) (*peer.ProposalResponse, error)) {
	fake.processProposalMutex.Lock()
	defer fake.processProposalMutex.Unlock()
	fake.ProcessProposalStub = stub
}

func (fake *EndorserServer) ProcessProposalArgsForCall(i int) (context.Context, *peer.SignedProposal) {
	fake.processProposalMutex.RLock()
	defer fake.processProposalMutex.RUnlock()
	argsForCall := fake.processProposalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *EndorserServer) ProcessProposalReturns(result1 *peer.ProposalResponse, result2 error) {
	fake.processProposalMutex.Lock()
	defer fake.processProposalMutex.Unlock()
	fake.ProcessProposalStub = nil
	fake.processProposalReturns = struct {
		result1 *peer.ProposalResponse
		result2 error
	}{result1, result2}
}

func (fake *EndorserServer) ProcessProposalReturnsOnCall(i int, result1 *peer.ProposalResponse, result2 error) {
	fake.processProposalMutex.Lock()
	defer fake.processProposalMutex.Unlock()
	fake.ProcessProposalStub = nil
	if fake.processProposalReturnsOnCall == nil {
		fake.processProposalReturnsOnCall = make(map[int]struct {
			result1 *peer.ProposalResponse
			result2 error
		})
	}
	fake.processProposalReturnsOnCall[i] = struct {
		result1 *peer.ProposalResponse
		result2 error
	}{result1, result2}
}

func (fake *EndorserServer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.processProposalMutex.RLock()
	defer fake.processProposalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *EndorserServer) recordInvocation(key string, args []interface{}) {
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
