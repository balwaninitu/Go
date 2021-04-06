package util

import "html/template"

var TPL *template.Template

/*init func will execute before main func and parse all files in
one folder named templates which get executes as commanded*/
func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))

}
