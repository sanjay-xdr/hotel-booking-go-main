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

func TestRenderTemplate(t *testing.T){
	pathToTemplates="./../../../templates";

	tc, err:= CreateTemplateCache();

	if err != nil{
		t.Error(err)
	}
	app.TemplateCache=tc;
	r,err:=getSession();
	if err!= nil {
		t.Error(err)
	}

	var ww myWriter


	err=RenderTemplate(&ww,r,"home.page.html", &models.TemplateData{});

	if err != nil {
		t.Error("Failed Render templateś")
	}

	err=RenderTemplate(&ww,r,"tewting.page.html", &models.TemplateData{});

	if err==nil {
		t.Error("This page doesnt exist")
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

func TestNewTemplates(t *testing.T){

	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T){
	pathToTemplates="./../../../templates";
	_,err:=CreateTemplateCache();

	if err !=nil {
		t.Error("Error")
	}

}