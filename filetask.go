package filetask

type Task func() error

type Task interface {
	Do() error
}

type TaskFunc func() error

func (f TaskFunc) Do() error {
	return f()
}

type TaskSet []Task

func (ts TaskSet) Do() error {
	for _, task := range ts {
		if err := task.Do(); err != nil {
			return err
		}
	}
	return nil
}

type Worker <-chan Task

func (w Worker) Run() (done chan<- error) {
	go func() {
		for task := range w {
			done <- task.Do()
		}
	}()
	return
}

type FileTask struct {
	Task
	Input  string
	Output string
}

func (ft *FileTask) Do() error {
}
