package main

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", Handle)

	err := http.ListenAndServe("localhost:8080", router)//ffff
	if err != nil {
		panic(err)
	}
}
}
