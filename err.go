package irdata

import "fmt"

func generateError(msg string, trigger error) string {
	if trigger != nil && trigger.Error() != "" && msg != "" {
		return fmt.Sprintf("%s: %s", msg, trigger.Error())
	} else if trigger != nil && trigger.Error() != "" {
		return trigger.Error()
	} else {
		return msg
	}
}

type ConfigurationError struct {
	Msg     string
	Trigger error
}

func (e *ConfigurationError) Error() string {
	return generateError(e.Msg, e.Trigger)
}

type AuthenticationError struct {
	Msg     string
	Trigger error
}

func (e *AuthenticationError) Error() string {
	return generateError(e.Msg, e.Trigger)
}

type ServiceUnavailableError struct {
	Msg     string
	Trigger error
}

func (e *ServiceUnavailableError) Error() string {
	return generateError(e.Msg, e.Trigger)
}

type LinkError struct {
	Msg     string
	Trigger error
}

func (e *LinkError) Error() string {
	return generateError(e.Msg, e.Trigger)
}

type RateLimitExceededError struct {
	Msg     string
	Trigger error
}

func (e *RateLimitExceededError) Error() string {
	return generateError(e.Msg, e.Trigger)
}
