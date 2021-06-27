package helper

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

var logger = log.New(os.Stdout, "workout", log.LstdFlags)

func Logf(r *http.Request, format string, values ...interface{}) {
	userID, ok := GetClaim(r, "id").(string)
	if !ok {
		logger.Printf(format+"\n", values)
	} else {
		logger.Printf("[id=%s] "+format+"\n", userID, values)
	}
}

func LogError(req *http.Request, err error, rv interface{}) {
	if err != nil {
		var buf bytes.Buffer
		userID, ok := GetClaim(req, "id").(string)
		if ok {
			fmt.Fprintf(&buf, "Error serving %s: [id=%s]%v\n", req.URL, userID, err)
		} else {
			fmt.Fprintf(&buf, "Error serving %s: %v\n", req.URL, err)
		}

		if rv != nil {
			fmt.Fprintln(&buf, rv)
			buf.Write(debug.Stack())
		}
		log.Print(buf.String())
	}
}
