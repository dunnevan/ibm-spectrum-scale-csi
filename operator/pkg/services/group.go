package services

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

type Group struct {
	pending int
	Ch      chan Outcome
}

func NewGroup() Group {
	return Group{
		Ch: make(chan Outcome, 10),
	}
}

func (g *Group) Wait() Outcome {
	out := Outcome{}
	for g.pending > 0 {
		result := <-g.Ch
		out = out.Foldl(result)
		g.pending--
		if out.Err != nil {
			//unreconcilable error, so stop waiting
			return out
		}
	}
	return out
}

func (g *Group) WaitBlockErr() Outcome {
	out := Outcome{}
	for g.pending > 0 {
		out = out.Foldl(<-g.Ch)
		g.pending--
	}
	return out
}

func (g *Group) Do(functions ...func() (ctrl.Result, error)) *Group {
	for _, f := range functions {
		g.pending++
		go func(f func() (ctrl.Result, error)) {
			result, err := f()
			g.Ch <- Outcome{result, err}
		}(f)
	}
	return g
}

// Outcome is a Reconcile return, wrapper of tuple: (Result, error)
type Outcome struct {
	//Result
	Result ctrl.Result
	//Err is a hard, unreconcilable error.  Reconcile will not be reattempted
	Err error
}

// Foldl an Outcome into an existing Outcome.
func (o1 Outcome) Foldl(o2 Outcome) Outcome {
	if o1.Err == nil {
		o1.Err = o2.Err
	}
	if !o1.Result.Requeue {
		o1.Result.Requeue = o2.Result.Requeue
	}
	if o1.Result.RequeueAfter == 0 || (o2.Result.RequeueAfter < o1.Result.RequeueAfter && o2.Result.RequeueAfter != 0) {
		o1.Result.RequeueAfter = o2.Result.RequeueAfter
	}

	return o1
}

func (o Outcome) HasEffect() bool {
	return o.Err != nil || !o.Result.IsZero()
}
