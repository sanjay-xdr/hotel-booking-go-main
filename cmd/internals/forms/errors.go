package forms

// import "crypto/internal/edwards25519/field"

// import "golang.org/x/text/message"

type errors map[string][]string

// adds an error
func (e errors) Add(field, message string) {

	e[field] = append(e[field], message)

}

//this returns the first error message
func (e errors) Get(field string) string {

	es := e[field]

	if len(es) == 0 {
		return ""
	}

	return es[0]

}
