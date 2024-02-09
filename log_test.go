package log

import (
	"testing"
)

var localPath = []string{"/tmp/test.log"}

func TestLog(t *testing.T) {
	log, _ := NewDevelopmentLogger(localPath)
	log.Debug("this is Debug log wrappered")
	log.Info("this is Info log wrappered")
	log.Warn("this is Warn log wrappered")
	log.Error("this is Error log wrappered")
	log.Fatal("this is Fatal log wrappered")
}
