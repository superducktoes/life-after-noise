package main

import (
	"net/http"
	"os"
)

func main() {
	finish := make(chan bool)

	server38000 := http.NewServeMux()
	server38000.HandleFunc("/", foo38000)

	server56962 := http.NewServeMux()
	server56962.HandleFunc("/", foo56962)

	server44042 := http.NewServeMux()
	server44042.HandleFunc("/", foo44042)

	server50880 := http.NewServeMux()
	server50880.HandleFunc("/", foo50880)

	server65535 := http.NewServeMux()
	server65535.HandleFunc("/", foo65535)

	server60416 := http.NewServeMux()
	server60416.HandleFunc("/", foo60416)

	server64977 := http.NewServeMux()
	server64977.HandleFunc("/", foo64977)

	server13269 := http.NewServeMux()
	server13269.HandleFunc("/", foo13269)

	server55364 := http.NewServeMux()
	server55364.HandleFunc("/", foo55364)

	server6066 := http.NewServeMux()
	server6066.HandleFunc("/", foo6066)

	server7550 := http.NewServeMux()
	server7550.HandleFunc("/", foo7550)

	server4588 := http.NewServeMux()
	server4588.HandleFunc("/", foo4588)

	server42043 := http.NewServeMux()
	server42043.HandleFunc("/", foo42043)

	server12060 := http.NewServeMux()
	server12060.HandleFunc("/", foo12060)

	serverShell := http.NewServeMux()
	serverShell.HandleFunc("/shell", serveShell)

	go func() {
		http.ListenAndServe(":38000", server38000)
	}()

	go func() {
		http.ListenAndServe(":56962", server56962)
	}()

	go func() {
		http.ListenAndServe(":44042", server44042)
	}()

	go func() {
		http.ListenAndServe(":50880", server50880)
	}()

	go func() {
		http.ListenAndServe(":65535", server65535)
	}()

	go func() {
		http.ListenAndServe(":60416", server60416)
	}()

	go func() {
		http.ListenAndServe(":64977", server64977)
	}()

	go func() {
		http.ListenAndServe(":13269", server13269)
	}()

	go func() {
		http.ListenAndServe(":55364", server55364)
	}()

	go func() {
		http.ListenAndServe(":6066", server6066)
	}()

	go func() {
		http.ListenAndServe(":7550", server7550)
	}()

	go func() {
		http.ListenAndServe(":4588", server4588)
	}()

	go func() {
		http.ListenAndServe(":42043", server42043)
	}()

	go func() {
		http.ListenAndServe(":12060", server12060)
	}()

	<-finish
}

func serveShell(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))

	err := os.WriteFile("/tmp/dat1", r, 0644)
	check(err)

	defer f.Close()

}

func foo38000(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo56962(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo44042(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo50880(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo65535(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo60416(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo64977(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo13269(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo55364(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo6066(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo7550(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo4588(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo42043(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

func foo12060(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}

/*
	server12060 := http.NewServeMux()
	server12060.HandleFunc("/", foo12060)

	go func() {
		http.ListenAndServe(":12060", server12060)
	}()

	func foo12060(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test"))
}
*/
