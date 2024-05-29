package render

import (
	"net/http"
	"testing"

	"github.com/sanjay-xdr/cmd/internals/models"
)

 func TestAddDefaultData(t *testing.T) {

	var td models.TemplateData

	r, err := getSession();


	if err != nil{
		// log.Print("Went wrong ")
		t.Error(err)
	}
	session.Put(r.Context(),"flash","123");

	result := AddDefaultData(&td, r)

	if result.Flash!="123" {
		t.Error("RFlash mil gya bc")
	}
}

func getSession()( *http.Request, error){

	r, err := http.NewRequest("GET", "/someurl", nil) 

	if err != nil{
		// log.Print("Went wrong ")
		return nil,err
	}

	ctx:=r.Context();

	ctx,_=session.Load(ctx, r.Header.Get("X-Session"))

	r=r.WithContext(ctx);
	return r,nil;
}