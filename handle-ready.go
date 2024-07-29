package main
import(
	"net/http"
	"fmt"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AJFAJ")
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "okajf aj","lhih":"higg"})
}
