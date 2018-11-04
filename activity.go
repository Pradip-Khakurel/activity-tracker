package activitytracker

const EPOCH float64 = 0

type Activity struct {
	State State
	Start float64
	End   float64
}
