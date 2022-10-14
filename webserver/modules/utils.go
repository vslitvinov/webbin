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

func LoadConfigPage(c *Cache, name string, url string,data interface{}) error {
    data, err := LoadConfigFile(data,url)
    if err != nil {
        log.Println(err)
        return err 
    }

    c.Set(name, data)
    return nil

}

func LoadConfigFile(data interface{}, url string) (interface{}, error){
    data = &models.PageHome{}
    file, err := os.ReadFile(url)
    if err != nil {
        log.Println(err)
        return nil, err
    }
    
    if err := json.Unmarshal(file,&data ); err != nil {
        log.Println(err)
        return nil, err
    }
log.Println(data)
    return data, nil

}