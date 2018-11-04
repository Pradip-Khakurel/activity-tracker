package activitytracker

type State int

const (
	StateNone State = 0
	State1    State = 1
	State2    State = 2
	State3    State = 4
	State4    State = 8
	State5    State = 16
	State6    State = 32
	State7    State = 64
	State8    State = 128
	State9    State = 256
	State10   State = 512
	StateAll  State = 1023
)
