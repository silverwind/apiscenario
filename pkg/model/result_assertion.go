package model

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/silverwind/api-scenario/pkg/log"
)

type ResultAssertion struct {
	Success  bool   `json:"success"`
	Source   Source `json:"source,omitempty"`
	Property string `json:"property,omitempty"`
	Err      error  `json:"error,omitempty"`
	Message  string `json:"message,omitempty"`
}

func NewResultAssertion(comparison Comparison, success bool, v ...interface{}) ResultAssertion {
	msg := comparison.GetMessage()
	var message string
	if success {
		message = fmt.Sprintf(msg.Success, v...)
	} else {
		message = fmt.Sprintf(msg.Failure, v...)
	}
	return ResultAssertion{
		Success: success,
		Message: message,
		Err:     nil,
	}
}

var sourceDisplayName = map[Source]string{
	ResponseJson:   "body",
	ResponseTime:   "Response time",
	ResponseStatus: "status",
	ResponseHeader: "header",
}

func (ar *ResultAssertion) Print() {
	source := sourceDisplayName[ar.Source]
	if len(ar.Property) > 0 {
		source += "." + ar.Property
	}

	if ar.Success {
		logrus.Infof(log.SuccessColor.Sprint("\u2713\t")+"%s - %s", source, ar.Message)
		return
	}

	logrus.Errorf("X\t%s - %s", source, ar.Message)
	if logrus.IsLevelEnabled(logrus.DebugLevel) && ar.Err != nil {
		logrus.Debug(ar.Err)
	}
}
