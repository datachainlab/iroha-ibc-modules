package evm

import (
	"github.com/hyperledger/burrow/execution/errors"
	"github.com/hyperledger/burrow/execution/exec"
	"github.com/hyperledger/burrow/txs/payload"
)

var _ exec.EventSink = (*eventSink)(nil)

type eventSink struct {
	height uint64
	Events []*exec.Event
}

func newEventSink(height uint64) *eventSink {
	return &eventSink{
		height: height,
		Events: []*exec.Event{},
	}
}

func (s eventSink) Print(print *exec.PrintEvent) error {
	s.Append(&exec.Event{
		Header: &exec.Header{
			TxType:    payload.TypeCall,
			TxHash:    nil,
			EventType: exec.TypePrint,
			EventID:   exec.EventStringLogEvent(print.Address),
			Height:    s.height,
			Exception: nil,
		},
		Print: print,
	})

	return nil
}

func (s eventSink) Log(log *exec.LogEvent) error {
	s.Append(&exec.Event{
		Header: &exec.Header{
			TxType:    payload.TypeCall,
			TxHash:    nil,
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
			TxHash:    nil,
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
