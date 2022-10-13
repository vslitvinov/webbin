package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"html/template"
)



func TempLoad() (*template.Template, error){


 var allFiles []string
    files, err := ioutil.ReadDir("./webserver/template/auth/")
    if err != nil {
        fmt.Println(err)
    }

    for _, file := range files {
        filename := file.Name()
        if strings.HasSuffix(filename, ".tmpl") {
            allFiles = append(allFiles, "./webserver/template/auth/"+filename)
        }
    }



    templates, err := template.ParseFiles(allFiles...) 
    if err != nil {
    	log.Println(err)
    	return nil, err
    }
    return templates, nil 

    // s1 := templates.Lookup("header.tmpl")
    // s1.ExecuteTemplate(os.Stdout, "header", nil)
    // fmt.Println()
    // s2 := templates.Lookup("content.tmpl")
    // s2.ExecuteTemplate(os.Stdout, "content", nil)
    // fmt.Println()
    // s3 := templates.Lookup("footer.tmpl")
    // s3.ExecuteTemplate(os.Stdout, "footer", nil)
    // fmt.Println()
    // s3.Execute(os.Stdout, nil)


}