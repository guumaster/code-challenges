package queue_test

import (
	"testing"

	"github.com/guumaster/crawler/queue"
	"github.com/stretchr/testify/mock"
)

type mockJob struct {
	mock.Mock
}

func (m *mockJob) Run() {
	m.Called()
}

func TestQueue_AddJob(t *testing.T) {
	// Create a mock function to pass to the Queue's AddJob method
	mockJob := &mockJob{}
	mockJob.On("Run", mock.Anything).Return()

	q := queue.NewQueue(2)

	// Add 3 jobs to the queue
	q.AddJob(mockJob.Run)
	q.AddJob(mockJob.Run)
	q.AddJob(mockJob.Run)

	q.Start()

	q.Wait()

	// Assert that the inner jobs channel received exactly 3 jobs
	mockJob.AssertNumberOfCalls(t, "Run", 3)
}
