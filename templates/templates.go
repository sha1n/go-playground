package templates

import (
	"os"
	"text/template"
	"time"
	"path"
	"fmt"
)

type User struct {
	Name     string
	Password string
	Roles    []string
	CreationTime time.Time
	Active bool
}

func Demo() {

	fmt.Println("*** Go templates ***")

	desc := User{
		Name:     "jhendrix",
		Password: "secret",
		Roles:    []string{"user", "admin"},
		CreationTime: time.Now(),
		Active: false,
	}

	functions := template.FuncMap{
		"yesno": func(b bool) string {
			if b {
				return "Yes"
			} else {
				return "No"
			}
		},
	}

	templatePath := "templates/user.tpl"
	templateName := path.Base(templatePath) // yes, it's weird.. see https://golang.org/pkg/text/template/#Template.ParseFiles

	err := template.Must(
		template.
			New(templateName).
			Funcs(functions).
			ParseFiles(templatePath),
	).Execute(os.Stdout, desc)

	if err != nil {
		panic(err)
	}
}
