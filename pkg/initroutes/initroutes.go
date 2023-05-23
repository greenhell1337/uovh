package initroutes

import (
	"net/http"
	"uovh/pkg/handlers"
	"uovh/pkg/cfs"
)

func InitRotes(mux *http.ServeMux) http.ServeMux{
	fileServer := http.FileServer(cfs.CloseFileSystem{Fs: http.Dir("./ui/static/")})

	mux.HandleFunc("/home", handlers.MainPage)
	
	mux.HandleFunc("/video=1", handlers.FirstVideoPage)
	mux.HandleFunc("/video=2", handlers.SecondVideoPage)
	mux.HandleFunc("/video=3", handlers.ThirdVideoPage)
	
	mux.Handle("/static", http.NotFoundHandler())
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return *mux
}