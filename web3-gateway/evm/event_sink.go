package evm

import (
	"github.com/hyperledger/burrow/execution/errors"
	"github.com/hyperledger/burrow/execution/exec"
	"github.com/hyperledger/burrow/txs/payload"
)

var _ exec.EventSink = (*eventSink)(nil)

type eventSink struct {
	txHash []byte
	height uint64
	Events []*exec.Event
}

func newEventSink(txHash []byte, height uint64) *eventSink {
	return &eventSink{
		txHash: txHash,
		height: height,
		Events: []*exec.Event{},
	}
}

func (s eventSink) Log(log *exec.LogEvent) error {
	s.Append(&exec.Event{
		Header: &exec.Header{
			TxType:    payload.TypeCall,
			TxHash:    s.txHash,
			EventType: exec.TypeLog,
			EventID:   exec.EventStringLogEvent(log.Address),
			Height:    s.height,
			Exception: nil,
		},
		Log: log,
	})
	return nil
}

func (s eventSink) Call(call *exec.CallEvent, exception *errors.Exception) error {
	s.Append(&exec.Event{
		Header: &exec.Header{
			TxType:    payload.TypeCall,
			TxHash:    s.txHash,
			EventType: exec.TypeCall,
			EventID:   exec.EventStringAccountCall(call.CallData.Callee),
			Height:    s.height,
			Exception: exception,
		},
		Call: call,
	})
	return nil
}

func (s eventSink) Append(tail ...*exec.Event) {
	for i, ev := range tail {
		if ev != nil && ev.Header != nil {
			ev.Header.Index = uint64(len(s.Events) + i)
		}
	}
	s.Events = append(s.Events, tail...)
}

func (s eventSink) CallTrace() string {
	return exec.Events(s.Events).CallTrace()
}
