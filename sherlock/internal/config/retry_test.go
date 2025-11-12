package config

import (
	"errors"
	"testing"

	"github.com/avast/retry-go/v4"
	"github.com/stretchr/testify/assert"
)

func TestRetryOptions(t *testing.T) {
	LoadTestConfig()

	t.Run("retryable error", func(t *testing.T) {
		functionCalls := 0
		functionThatErrorsAtFirst := func() error {
			switch functionCalls {
			case 0:
				functionCalls++
				return errors.New("some sherlock retryable error: blah")
			case 1:
				functionCalls++
				return nil
			default:
				panic("function called too many times!")
			}
		}

		err := retry.Do(functionThatErrorsAtFirst, RetryOptions...)
		assert.NoError(t, err)
		assert.Equal(t, 2, functionCalls)
	})

	t.Run("exhausted attempts", func(t *testing.T) {
		functionCalls := 0
		functionThatErrorsAtFirst := func() error {
			functionCalls++
			return errors.New("some sherlock retryable error: blah")
		}

		err := retry.Do(functionThatErrorsAtFirst, RetryOptions...)
		assert.Error(t, err)
		assert.Equal(t, 2, functionCalls)
	})

	t.Run("non-retryable error", func(t *testing.T) {
		functionCalls := 0
		functionThatErrorsAtFirst := func() error {
			switch functionCalls {
			case 0:
				functionCalls++
				return errors.New("some non-retryable error: blah")
			case 1:
				functionCalls++
				return nil
			default:
				panic("function called too many times!")
			}
		}

		err := retry.Do(functionThatErrorsAtFirst, RetryOptions...)
		assert.Error(t, err)
		assert.Equal(t, 1, functionCalls)
	})

	t.Run("unrecoverable error", func(t *testing.T) {
		functionCalls := 0
		functionThatErrorsAtFirst := func() error {
			switch functionCalls {
			case 0:
				functionCalls++
				return retry.Unrecoverable(errors.New("some sherlock retryable error: blah"))
			case 1:
				functionCalls++
				return nil
			default:
				panic("function called too many times!")
			}
		}

		err := retry.Do(functionThatErrorsAtFirst, RetryOptions...)
		assert.Error(t, err)
		assert.Equal(t, 1, functionCalls)
	})
}
