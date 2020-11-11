package utils

import (
	"bufio"
	"os"
)

func Writer(s string, fp string) error {
	file, err := os.OpenFile(fp, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	// _, err = writer.WriteString(s)
	_, err = writer.WriteString(s + "\n")
	if err != nil {
		return err
	}

	writer.Flush()
	return nil

}
