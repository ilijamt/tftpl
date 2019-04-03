package cmd

import (
	"github.com/ilijamt/tftpl"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

func openFile(filename string) (reader io.ReadCloser, err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return nil, err
	}
	return file, nil
}

func readDataFile(filename string) (data *tftpl.Template, err error) {
	var rc io.ReadCloser
	if rc, err = openFile(filename); err != nil {
		return nil, err
	}
	defer rc.Close()

	buff, _ := ioutil.ReadAll(rc)

	data = new(tftpl.Template)
	err = yaml.Unmarshal(buff, &data)
	return
}
