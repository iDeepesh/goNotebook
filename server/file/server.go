package file

import (
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

//ServeFilesInManyWays - serves a file by uisng IO package to copy from input to output stream
func ServeFilesInManyWays() {
	defHandler := func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("file/helloImage.gohtml"))
		e := t.ExecuteTemplate(w, "helloImage.gohtml", nil)
		if e != nil {
			log.Fatal(e)
		}
	}

	imgHandlerIOCopy := func(w http.ResponseWriter, r *http.Request) {
		f, e := os.Open("file/golang.png")
		if e != nil {
			log.Fatal(e)
		}
		defer f.Close()

		io.Copy(w, f)
	}

	imgHandlerServeContent := func(w http.ResponseWriter, r *http.Request) {
		f, e := os.Open("file/gogogo.jpeg")
		if e != nil {
			log.Fatal(e)
		}
		defer f.Close()

		fi, e := os.Stat("file/gogogo.jpeg")
		if e != nil {
			log.Fatal(e)
		}

		http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
	}

	imgHandlerServeFile := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "file/golangticket.png")
	}

	http.HandleFunc("/", defHandler)
	http.HandleFunc("/golang.png", imgHandlerIOCopy)
	http.HandleFunc("/gogogo.jpeg", imgHandlerServeContent)
	http.HandleFunc("/golangticket.png", imgHandlerServeFile)
	http.ListenAndServe(":6080", nil)
}

//ServeWithFileServer - Uses file server to serve files
func ServeWithFileServer() {
	http.Handle("/fs/", http.StripPrefix("/fs", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":6081", nil)
}

//StaticFileServer - serves static files
func StaticFileServer() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./file/"))))
	http.ListenAndServe(":6082", nil)
}
