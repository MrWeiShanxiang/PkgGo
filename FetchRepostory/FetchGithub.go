package FetchRepostory

import (
	"io/ioutil"

	"github.com/robertkrimen/otto"
)

func ListRepostory(url string) ([]string, error) {
	JsFilepath := "./js/index.js"
	bytes, err := ioutil.ReadFile(JsFilepath)
	if err != nil {
		return nil, err
	}
	vm := otto.New()
	ValueObject, err := vm.Run(bytes)
	if err != nil {
		return nil, err
	}
	return ValueObject.Object().Keys(), nil
}
