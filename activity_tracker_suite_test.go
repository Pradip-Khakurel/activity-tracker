package activitytracker_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestActivityTracker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ActivityTracker Suite")
}
