package main

import (
	"log"
	"net/http"
	"html/template"
	"io/ioutil"
)

var(
	Templates *template.Template
	GError error
)

func HandleError(err error, ErrMessage, Succmessage string){
	if err!=nil{
		log.Fatalln(err)
		log.Println(ErrMessage)
	} else if Succmessage != ""{
		log.Println(Succmessage)
	}
}

func init(){

	Templates, GError = template.ParseGlob("templates/*")
	HandleError(GError, "Error parsing glob from templates", "Parsed glob from templates")
}

func main(){
	
	files, err := ioutil.ReadDir("./storage/")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
            log.Println(f.Sys())
    }

	http.HandleFunc("/", Home)
	http.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir("storage"))))

	log.Println("Listening on  :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))



}

func Home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Cache-Control", "max-age=1000000")
	err := Templates.ExecuteTemplate(w,"home.html", nil)
	HandleError(err, "Error executing template", "")
}