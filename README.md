# activity-tracker

Given a list of activities return a list of transitions.

## usage

```
activities := []*Activity{}
activities = append(activities, &Activity{State: State1, Start: 10, End: 50})

transitions := ComputeTransitions(activities)

```

## tests

Tests use gingko.

```
go test
```

or
```
ginkgo
```