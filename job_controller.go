package yggdrasil

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// Job wraps a playbook along with status and output information.
type Job struct {
	ID       string `json:"id"`
	Playbook string `json:"playbook"`
	Status   string `json:"status"`
	Stdout   string `json:"stdout"`
}

// JobController manages the execution and collects the output of a Job.
type JobController struct {
	job    Job
	client *HTTPClient
	url    string
}

// Start begins execution of the job.
func (j *JobController) Start() error {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	if _, err := f.WriteString(j.job.Playbook); err != nil {
		return err
	}
	defer func() {
		name := f.Name()
		f.Close()
		os.Remove(name)
	}()

	cmd := exec.Command("ansible-playbook", "--ssh-common-args=-oStrictHostKeyChecking=no", f.Name())
	cmd.Stderr = os.Stderr
	output, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	reader := bufio.NewScanner(output)
	for reader.Scan() {
		j.job.Status = "running"
		j.job.Stdout += reader.Text() + "\n"

		if err := j.Update("running", reader.Text()); err != nil {
			return err
		}
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	if err := j.Finish(); err != nil {
		return err
	}

	return nil
}

// Update sends a "running" status update, along with a slice of stdout to the
// job service.
func (j *JobController) Update(status, stdout string) error {
	update := struct {
		Status string `json:"status"`
		Stdout string `json:"stdout"`
	}{
		Status: status,
		Stdout: stdout,
	}
	data, err := json.Marshal(update)
	if err != nil {
		return err
	}

	resp, err := j.client.Patch(j.url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusAccepted {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return &APIResponseError{
			Code: resp.StatusCode,
			body: string(data),
		}
	}
	return nil
}

// Finish completes the job by sending a "complete" status to the job service.
func (j *JobController) Finish() error {
	j.job.Status = "complete"
	data, err := json.Marshal(j.job)
	if err != nil {
		return err
	}

	resp, err := j.client.Put(j.url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusAccepted {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		return &APIResponseError{
			Code: resp.StatusCode,
			body: string(data),
		}
	}
	return nil
}