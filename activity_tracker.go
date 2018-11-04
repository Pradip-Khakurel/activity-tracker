package activitytracker

import "sort"

type activityPart struct {
	state   State
	date    float64
	isStart bool
}

// given a list of activities compute a list of transitions
func ComputeTransitions(activities []*Activity) []*Transition {
	transitions := []*Transition{}
	tr := &Transition{State: StateNone, Start: EPOCH}
	transitions = append(transitions, tr)

	var prev *Transition
	var cur *Transition

	prev = transitions[0]
	parts := createSortedParts(cleanActivities(activities))
	opened := make(map[State]int)

	for _, prt := range parts {

		if prt.date > prev.Start {
			cur = &Transition{State: prev.State, Start: prt.date}
			transitions = append(transitions, cur)
			prev = cur
		}

		if prt.isStart {
			i, ok := opened[prt.state]

			if ok {
				opened[prt.state] = i + 1
			} else {
				opened[prt.state] = 1
			}

			cur.State |= prt.state // add state

		} else {
			opened[prt.state]--

			if opened[prt.state] == 0 {
				cur.State &= StateAll &^ prt.state // remove state iff it is not "opened" anymore
			}
		}

	}

	return distincts(transitions)
}

// remove inconsistent activities
func cleanActivities(activities []*Activity) []*Activity {
	cleaned := []*Activity{}

	for _, act := range activities {
		if act.Start >= act.End {
			continue
		}
		cleaned = append(cleaned, act)
	}

	return cleaned
}

func createSortedParts(activities []*Activity) []*activityPart {
	parts := []*activityPart{}

	for _, act := range activities {

		parts = append(parts, &activityPart{date: act.Start, state: act.State, isStart: true})
		parts = append(parts, &activityPart{date: act.End, state: act.State, isStart: false})
	}

	sort.Sort(byDate(parts))

	return parts
}

func distincts(transitions []*Transition) []*Transition {
	uniques := []*Transition{}
	uniques = append(uniques, transitions[0])

	for i := 1; i < len(transitions); i++ {
		prev := (transitions)[i-1]
		cur := (transitions)[i]

		if cur.State == prev.State {
			continue
		}

		uniques = append(uniques, cur)
	}

	return uniques
}
