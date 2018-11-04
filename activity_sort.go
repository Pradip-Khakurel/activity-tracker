package activitytracker

type byDate []*activityPart

func (activities byDate) Len() int {
	return len(activities)
}

func (activities byDate) Swap(i, j int) {
	activities[i], activities[j] = activities[j], activities[i]
}
func (activities byDate) Less(i, j int) bool {
	return activities[i].date < activities[j].date
}
