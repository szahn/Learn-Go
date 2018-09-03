package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 8080

	args := os.Args[1:]
	if len(args) > 0 {
		port, _ = strconv.Atoi(args[0])
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<!DOCTYPE html>
			<html>
				<head>
					<title>Simple Go Web Server</title>
					<style>
						body{
							margin: auto 10px;
							text-align: center;
							font-size:2em;
						}
					</style>
				</head>
				<body>
					<h1>Hello World!</h1>
					<script>
						(function(){
							var captions = [
								"Hello Wêreld",
								"مرحبا بالعالم",
								"你好，世界",
								"Hej Verden",
								"Hei maailma",
								"Bonjour le monde",
								"Hallo Welt",
								"Γειά σου Κόσμε",
								"Hello Aloha",
								"Halló heimur",
								"Ciao mondo",
								"こんにちは世界",
								"ສະ​ບາຍ​ດີ​ຊາວ​ໂລກ",
								"Hej världen",
								"สวัสดีชาวโลก",
								"Привіт Світ",
								"Chào thế giới",
								"Sawubona Mhlaba",
								"Hello World"
							];

							var i = 0;
							setInterval(function(){
								var el = document.querySelectorAll("h1");
								el[0].innerHTML = captions[i]
								i +=1;
								if (i > captions.length -1) i =0;
							}, 3000);

						})();
					</script>
				</body>
			</html		
		`))
	})

	serverAddress := ":" + strconv.Itoa(port)
	log.Printf("Starting server on %s\r\n", serverAddress)
	if err := http.ListenAndServe((serverAddress), nil); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
