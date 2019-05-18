package str

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
)

func ConnString(strs ...string) string {
	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func GetPageCount(pageSize, total int) (sizes int) {
	if total%pageSize == 0 {
		sizes = total / pageSize
	} else {
		sizes = total/pageSize + 1
	}
	return
}

func OsKill(closeAction func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	close(c)
	if sig == os.Interrupt || sig == os.Kill {
		fmt.Println("Close signal!")
		closeAction()
	}
}
