package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func TestLocation(t *testing.T) {
	log := log.New(os.Stdout, "TEST: ", log.Lshortfile|log.LstdFlags)
	sd := make(chan os.Signal, 1)

	app := NewAPI(log, sd)

	c := struct {
		X   string
		Y   string
		Z   string
		Vel string
	}{
		X:   "4.0",
		Y:   "10.0",
		Z:   "3.0",
		Vel: "2.0",
	}

	data, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest("GET", "/v1/atlas/dns/location/10", bytes.NewBuffer(data))
	w := httptest.NewRecorder()

	app.ServeHTTP(w, r)
	resp := struct {
		Loc float64 `json:"loc"`
	}{}
	t.Log("Given section id: 10, X: 4, Y: 10, Z: 3, Vel: 2.")
	{
		t.Log("\tTest 0:\tVanila path.")
		{
			if w.Code != http.StatusOK {
				t.Fatalf("\t%s\tShould receive a status code of 200 for the response : %v", Failed, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of 200 for the response.", Success)
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatal(err)
			}
			want := float64(10*4 + 10*10 + 10*3 + 10*2)

			if resp.Loc != want {
				t.Logf("\t%s\twant %v got %v", Failed, want, resp.Loc)
				t.Fatalf("\tAborting Test")
			}
			t.Logf("\t%s\twant %v got %v", Success, want, resp.Loc)
		}
	}

}
