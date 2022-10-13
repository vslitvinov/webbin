package modules

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
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

func LoadConfigPage(c *Cache, name string, url string) error {
    data, err := LoadConfigFile(url)
    if err != nil {
        log.Println(err)
        return err 
    }

    c.Set(name, data)
    return nil

}

func LoadConfigFile(url string) (interface{}, error){
    var data interface{}
    file, err := os.ReadFile("./web/data/home.json")
    if err != nil {
        log.Println(err)
        return nil, err
    }
    
    if err := json.Unmarshal(file,&data); err != nil {
        log.Println(err)
        return nil, err
    }

    return data, nil

}