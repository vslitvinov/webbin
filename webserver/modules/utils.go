package modules

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"site/webserver/models"
	"strings"
)



func TempLoad(url string) (*template.Template, error){


 var allFiles []string
    files, err := ioutil.ReadDir(url)
    if err != nil {
        fmt.Println(err)
    }

    for _, file := range files {
        filename := file.Name()
        if strings.HasSuffix(filename, ".html") {
            allFiles = append(allFiles, url+filename)
        }
    }


    templates, err := template.ParseFiles(allFiles...) 
    if err != nil {
    	log.Println(err)
    	return nil, err
    }
    return templates, nil 

}


func LoadConfigFile( url string) (map[string]models.Page, error){
    data := map[string]models.Page{}
    file, err := os.ReadFile(url)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    
    if err := json.Unmarshal(file,&data); err != nil {
        log.Println(err)
        return nil, err
    }
// log.Println(data)
    return data, nil

}