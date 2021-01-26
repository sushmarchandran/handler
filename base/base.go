package base

import (
	"bytes"
	"context"
	"errors"
	"html/template"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// TaskMeta identifies Library and Name of a task.
type TaskMeta struct {
	// Library where this task is defined.
	Library string `json:"library" yaml:"library"`
	// Name (type) of this task.
	Task string `json:"task" yaml:"task"`
}

// Task defines common method signatures for every task.
type Task interface {
	Run(ctx context.Context) error
	DryRun()
	Extrapolate(tags *Tags) error
}

// Tags supports string extrapolation using tags.
type Tags struct {
	M *map[string]string
}

// Extrapolate str using tags.
func (tags *Tags) Extrapolate(str *string) (string, error) {
	if tags == nil { // return a copy of the string
		return *str, nil
	}
	if templ, err := template.New("").Parse(*str); err == nil {
		buf := bytes.Buffer{}
		if err = templ.Execute(&buf, tags); err == nil {
			return string(buf.Bytes()), nil
		}
		return "", errors.New("cannot extrapolate string")
	}
	return "", errors.New("cannot extrapolate string")
}

// ContextKey is the type of key that will be used to index into context.
type ContextKey string