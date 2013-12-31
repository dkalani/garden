package job_tracker

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/vito/garden/backend"
	"github.com/vito/garden/command_runner"
)

type JobTracker struct {
	containerPath string
	runner        command_runner.CommandRunner

	jobs      map[uint32]*Job
	nextJobID uint32

	sync.RWMutex
}

type UnknownJobError struct {
	JobID uint32
}

func (e UnknownJobError) Error() string {
	return fmt.Sprintf("unknown job: %d", e.JobID)
}

func New(containerPath string, runner command_runner.CommandRunner) *JobTracker {
	return &JobTracker{
		containerPath: containerPath,
		runner:        runner,

		jobs: make(map[uint32]*Job),
	}
}

func (t *JobTracker) Spawn(cmd *exec.Cmd, discardOutput, autoLink bool) (uint32, error) {
	t.Lock()

	jobID := t.nextJobID
	t.nextJobID++

	job := NewJob(jobID, discardOutput, t.containerPath, t.runner)

	t.jobs[jobID] = job

	t.Unlock()

	ready, active := job.Spawn(cmd)

	err := <-ready
	if err != nil {
		return 0, err
	}

	if autoLink {
		go t.Link(jobID)

		err = <-active
		if err != nil {
			return 0, err
		}
	}

	return jobID, nil
}

func (t *JobTracker) Restore(jobID uint32, discardOutput bool) {
	t.Lock()

	job := NewJob(jobID, discardOutput, t.containerPath, t.runner)

	t.jobs[jobID] = job

	if jobID >= t.nextJobID {
		t.nextJobID = jobID + 1
	}

	t.Unlock()

	go t.Link(jobID)
}

func (t *JobTracker) Link(jobID uint32) (uint32, []byte, []byte, error) {
	t.RLock()
	job, ok := t.jobs[jobID]
	t.RUnlock()

	if !ok {
		return 0, nil, nil, UnknownJobError{jobID}
	}

	defer t.unregister(jobID)

	return job.Link()
}

func (t *JobTracker) Stream(jobID uint32) (chan backend.JobStream, error) {
	t.RLock()
	job, ok := t.jobs[jobID]
	t.RUnlock()

	if !ok {
		return nil, UnknownJobError{jobID}
	}

	go t.Link(jobID)

	return job.Stream(), nil
}

func (t *JobTracker) ActiveJobs() []*Job {
	t.RLock()
	defer t.RUnlock()

	jobs := []*Job{}

	for _, job := range t.jobs {
		jobs = append(jobs, job)
	}

	return jobs
}

func (t *JobTracker) unregister(jobID uint32) {
	t.Lock()
	defer t.Unlock()

	delete(t.jobs, jobID)
}
