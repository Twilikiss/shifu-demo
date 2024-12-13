package log

import (
	"log"
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel(ErrorLevel)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() != os.Stdout {
		log.Fatalf("[ErrorLevel] is error, please check it again")
	}

	SetLevel(Disabled)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() == os.Stdout {
		log.Fatalf("[Disabled] is error, please check it again")
	}
}
