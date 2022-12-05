package main

import (
	"fmt"
	"time"
	"math/rand"
	"net/http"
)

func main() {
	people := []string{"kuba", "seba", "zywiec", "doma", "maja", "ola"}
	 rand.Seed(time.Now().UnixNano())
	 rand.Shuffle(len(people), func(i, j int) { people[i], people[j] = people[j], people[i] })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if len(people) == 0 {
			fmt.Fprintf(w, "juz nie ma nie klikaj tyle debilu")
			return
		}
		num := rand.Intn(len(people))
		person := people[num]
		people = append(people[:num], people[num+1:]...)
		fmt.Fprintf(w, "masz: %s", person)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
