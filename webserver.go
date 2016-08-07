// web server basico
package main

import (
   "fmt"
   "log"
   "net/http"
   //  "strings"
   )

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
func appHandler(w http.ResponseWriter, r *http.Request) {
   urlc:=r.URL.Path
   //urln:=strings.ToLower(urlc[1:])
   fmt.Fprintf(w, "App URL = %q\n", urlc)
   /*	   
   if strings.HasPrefix(urln,"index") || urln=="" {
       http.FileServer(http.Dir("/tmp"))
   }else{
       //fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)   

   }
   */
}
   
func main() {
	http.HandleFunc("/otro/", appHandler)
	http.Handle("/fijo/",http.StripPrefix("/fijo/", http.FileServer(http.Dir("/tmp"))))
    http.HandleFunc("/", staticHandler) 
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
    //log.Fatal( http.ListenAndServe("localhost:8000", http.FileServer(http.Dir("/tmp") ) ) )
}


