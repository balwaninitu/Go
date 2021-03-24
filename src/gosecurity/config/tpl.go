package config

import "html/template"

var TPL *template.Template

/*init func will execute before main func and parse all files in
one folder named templates which get executes as once get call through routing*/
func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))

}
