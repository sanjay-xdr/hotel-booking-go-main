package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors //this error refers to the error in the same package
}

func New(data url.Values) *Form {
	return &Form{

		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	err := r.ParseForm()
	if err != nil {
		fmt.Print("Not able to parse the form")
	}

	x := r.Form.Get(field)

	if x == "" {
		f.Errors.Add(field, "Bhai ye field zaroori hai ")
		return false
	}
	return true

}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {

		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "bhai space mt daal")
		}
	}
}

func (f *Form) Valid() bool {

	return len(f.Errors) == 0

}

func (f *Form) IsEmail(field string) {

	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Please enter valid information")
	}

}
