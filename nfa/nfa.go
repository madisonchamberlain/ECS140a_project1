package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// if you the length is 0 and you start at the final state return true
	if len(input) == 0 {
		if start == final {
			return true
		} else {
			return false
		}
	} else {
		// find the states reachable from current state
		CurReachable := transitions(start, input[0])
		// for each reachable state recurse to see if you can reach the end
		for _, reachable := range CurReachable {
			if Reachable(transitions, reachable, final, input[1:len(input)]) {
				// if the previous recursion returns true; return true
				return true
			}
		}
		// if the previous recursion returns false; return false
		return false
	}
}
