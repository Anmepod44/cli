package executor

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/cli-command-assistant/internal/types"
	"github.com/google/uuid"
)

// Executor implements CommandExecutor interface
type Executor struct {
	maxConcurrent int
	running       map[string]*runningCommand
	mu            sync.RWMutex
	workerPool    chan struct{}
}

// runningCommand tracks a command being executed
type runningCommand struct {
	cmd    *exec.Cmd
	cancel context.CancelFunc
}

// NewExecutor creates a new command executor
func NewExecutor(maxConcurrent int) *Executor {
	if maxConcurrent <= 0 {
		maxConcurrent = 5 // default
	}

	return &Executor{
		maxConcurrent: maxConcurrent,
		running:       make(map[string]*runningCommand),
		workerPool:    make(chan struct{}, maxConcurrent),
	}
}

// Execute runs a command asynchronously and returns a result channel
func (e *Executor) Execute(cmd *types.Command) <-chan types.ExecutionResult {
	resultChan := make(chan types.ExecutionResult, 1)

	go func() {
		defer close(resultChan)

		// Generate unique ID for this command
		commandID := uuid.New().String()

		// Acquire worker slot
		e.workerPool <- struct{}{}
		defer func() { <-e.workerPool }()

		// Create context for cancellation
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Parse command (use sh -c to handle pipes and redirects)
		execCmd := exec.CommandContext(ctx, "sh", "-c", cmd.Raw)

		// Track running command
		e.mu.Lock()
		e.running[commandID] = &runningCommand{
			cmd:    execCmd,
			cancel: cancel,
		}
		e.mu.Unlock()

		// Cleanup tracking when done
		defer func() {
			e.mu.Lock()
			delete(e.running, commandID)
			e.mu.Unlock()
		}()

		// Capture stdout and stderr
		var stdout, stderr bytes.Buffer
		execCmd.Stdout = &stdout
		execCmd.Stderr = &stderr

		// Record start time
		startTime := time.Now()

		// Execute command
		err := execCmd.Run()

		// Calculate duration
		duration := time.Since(startTime)

		// Prepare result
		result := types.ExecutionResult{
			CommandID: commandID,
			Duration:  duration,
		}

		if err != nil {
			// Command failed
			result.Error = err
			result.ExitCode = execCmd.ProcessState.ExitCode()

			// Include stderr in output if available
			if stderr.Len() > 0 {
				result.Output = stderr.String()
			} else {
				result.Output = fmt.Sprintf("Command failed: %v", err)
			}
		} else {
			// Command succeeded
			result.ExitCode = 0
			result.Output = stdout.String()
		}

		resultChan <- result
	}()

	return resultChan
}

// ExecuteBatch runs multiple commands concurrently
func (e *Executor) ExecuteBatch(cmds []*types.Command) <-chan types.ExecutionResult {
	resultChan := make(chan types.ExecutionResult, len(cmds))

	go func() {
		defer close(resultChan)

		var wg sync.WaitGroup

		for _, cmd := range cmds {
			wg.Add(1)

			go func(c *types.Command) {
				defer wg.Done()

				// Execute command
				cmdResultChan := e.Execute(c)

				// Forward result
				if result, ok := <-cmdResultChan; ok {
					resultChan <- result
				}
			}(cmd)
		}

		wg.Wait()
	}()

	return resultChan
}

// Cancel stops a running command by ID
func (e *Executor) Cancel(id string) error {
	e.mu.RLock()
	running, exists := e.running[id]
	e.mu.RUnlock()

	if !exists {
		return fmt.Errorf("command %s not found or already completed", id)
	}

	// Cancel the context, which will kill the process
	running.cancel()

	return nil
}
