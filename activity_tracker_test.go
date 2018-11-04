package activitytracker

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("computeTransitions", func() {
	var activities []*Activity

	BeforeEach(func() {
		activities = []*Activity{}
	})

	It("should return activity none if empty activity list", func() {
		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(1))
		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[0].State).To(Equal(StateNone))
	})

	It("should return activity none if start equals end for the same activity", func() {
		activities = append(activities, &Activity{State: State1, Start: 10, End: 10})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(1))
		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[0].State).To(Equal(StateNone))
	})

	It("should return none if start is superior to end for the same activity", func() {
		activities = append(activities, &Activity{State: State1, Start: 15, End: 10})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(1))
		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[0].State).To(Equal(StateNone))
	})

	It("should return none then state1 then none", func() {
		activities = append(activities, &Activity{State: State1, Start: 10, End: 15})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(3))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(10)))
		Expect(trs[2].Start).To(Equal(float64(15)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(State1))
		Expect(trs[2].State).To(Equal(StateNone))
	})

	It("should return none in between two activities", func() {
		activities = append(activities, &Activity{State: State1, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State2, Start: 20, End: 30})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(5))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(10)))
		Expect(trs[2].Start).To(Equal(float64(15)))
		Expect(trs[3].Start).To(Equal(float64(20)))
		Expect(trs[4].Start).To(Equal(float64(30)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(State1))
		Expect(trs[2].State).To(Equal(StateNone))
		Expect(trs[3].State).To(Equal(State2))
		Expect(trs[4].State).To(Equal(StateNone))
	})

	It("should return two states simultaneously", func() {
		activities = append(activities, &Activity{State: State1, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State2, Start: 10, End: 15})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(3))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(10)))
		Expect(trs[2].Start).To(Equal(float64(15)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(State1 | State2))
		Expect(trs[2].State).To(Equal(StateNone))
	})

	It("should return all states simultaneously", func() {
		activities = append(activities, &Activity{State: State1, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State2, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State3, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State4, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State5, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State6, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State7, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State8, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State9, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State10, Start: 10, End: 15})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(3))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(10)))
		Expect(trs[2].Start).To(Equal(float64(15)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(StateAll))
		Expect(trs[2].State).To(Equal(StateNone))
	})

	It("should merge two identical activities into one bigger activity", func() {
		activities = append(activities, &Activity{State: State1, Start: 10, End: 15})
		activities = append(activities, &Activity{State: State1, Start: 15, End: 30})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(3))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(10)))
		Expect(trs[2].Start).To(Equal(float64(30)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(State1))
		Expect(trs[2].State).To(Equal(StateNone))
	})

	It("should return one activity with both states when two different activities overlap", func() {
		activities = append(activities, &Activity{State: State1, Start: 10, End: 20})
		activities = append(activities, &Activity{State: State2, Start: 15, End: 30})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(5))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(10)))
		Expect(trs[2].Start).To(Equal(float64(15)))
		Expect(trs[3].Start).To(Equal(float64(20)))
		Expect(trs[4].Start).To(Equal(float64(30)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(State1))
		Expect(trs[2].State).To(Equal(State1 | State2))
		Expect(trs[3].State).To(Equal(State2))
		Expect(trs[4].State).To(Equal(StateNone))
	})

	It("should ignore an activity if a bigger actvity with the same state contains this smaller activity", func() {
		activities = append(activities, &Activity{State: State1, Start: 10, End: 30})
		activities = append(activities, &Activity{State: State1, Start: 15, End: 20})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(3))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(10)))
		Expect(trs[2].Start).To(Equal(float64(30)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(State1))
		Expect(trs[2].State).To(Equal(StateNone))
	})

	It("should return one unique activity if activities with same state overlap", func() {
		activities = append(activities, &Activity{State: State1, Start: 5, End: 20})
		activities = append(activities, &Activity{State: State1, Start: 10, End: 30})
		activities = append(activities, &Activity{State: State1, Start: 15, End: 25})
		activities = append(activities, &Activity{State: State1, Start: 20, End: 100})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(3))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(5)))
		Expect(trs[2].Start).To(Equal(float64(100)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(State1))
		Expect(trs[2].State).To(Equal(StateNone))
	})

	It("should return multiple activities when activities with different states overlap", func() {
		activities = append(activities, &Activity{State: State1, Start: 5, End: 20})
		activities = append(activities, &Activity{State: State2, Start: 10, End: 30})
		activities = append(activities, &Activity{State: State3, Start: 15, End: 25})
		activities = append(activities, &Activity{State: State4, Start: 20, End: 100})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(8))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(5)))
		Expect(trs[2].Start).To(Equal(float64(10)))
		Expect(trs[3].Start).To(Equal(float64(15)))
		Expect(trs[4].Start).To(Equal(float64(20)))
		Expect(trs[5].Start).To(Equal(float64(25)))
		Expect(trs[6].Start).To(Equal(float64(30)))
		Expect(trs[7].Start).To(Equal(float64(100)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(State1))
		Expect(trs[2].State).To(Equal(State1 | State2))
		Expect(trs[3].State).To(Equal(State1 | State2 | State3))
		Expect(trs[4].State).To(Equal(State2 | State3 | State4))
		Expect(trs[5].State).To(Equal(State2 | State4))
		Expect(trs[6].State).To(Equal(State4))
		Expect(trs[7].State).To(Equal(StateNone))
	})

	It("should return multiple activities when activities with different states overlap, despite having unsorted activities", func() {
		activities = append(activities, &Activity{State: State2, Start: 10, End: 30})
		activities = append(activities, &Activity{State: State1, Start: 5, End: 20})
		activities = append(activities, &Activity{State: State4, Start: 20, End: 100})
		activities = append(activities, &Activity{State: State3, Start: 15, End: 25})

		trs := ComputeTransitions(activities)
		size := len(trs)

		Expect(size).To(Equal(8))

		Expect(trs[0].Start).To(Equal(EPOCH))
		Expect(trs[1].Start).To(Equal(float64(5)))
		Expect(trs[2].Start).To(Equal(float64(10)))
		Expect(trs[3].Start).To(Equal(float64(15)))
		Expect(trs[4].Start).To(Equal(float64(20)))
		Expect(trs[5].Start).To(Equal(float64(25)))
		Expect(trs[6].Start).To(Equal(float64(30)))
		Expect(trs[7].Start).To(Equal(float64(100)))

		Expect(trs[0].State).To(Equal(StateNone))
		Expect(trs[1].State).To(Equal(State1))
		Expect(trs[2].State).To(Equal(State1 | State2))
		Expect(trs[3].State).To(Equal(State1 | State2 | State3))
		Expect(trs[4].State).To(Equal(State2 | State3 | State4))
		Expect(trs[5].State).To(Equal(State2 | State4))
		Expect(trs[6].State).To(Equal(State4))
		Expect(trs[7].State).To(Equal(StateNone))
	})
})
