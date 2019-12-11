package main

import (
	"fmt"
	"keylime_verifier/app"
	"keylime_verifier/config"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func get(w http.ResponseWriter, r *http.Request) {
	myurl := r.URL.RequestURI()
	//fmt.Println(myurl)
	rest_params := app.GetRestfulParams(myurl)
	fmt.Println(rest_params)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
	// w.Write([]byte(rest_params))
}

func useAgents(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"Please use /agents"}`))
}

func getAgents(w http.ResponseWriter, r *http.Request) {
	val := chi.URLParam(r, "agent_uuid")
	if val != "" {
		fmt.Fprintf(w, "Hello %s!\n", val)
	} else {
		fmt.Fprintf(w, "Hello ... you.\n")
	}
}

func AgentActivate(w http.ResponseWriter, r *http.Request) {
	val := chi.URLParam(r, "agent_uuid")
	if val != "" {
		fmt.Fprintf(w, "Hello %s! going to activate you!\n", val)
		// do something with specific
	} else {
		fmt.Fprintf(w, "Hello ... you.\n")
		// get list of agents
	}
}

func main() {
	// db.init
	// db.update_all_agents('operational_state', cloud_verifier_common.CloudAgent_Operational_State.SAVED)
	app.InitTLS()
	port := config.GetConfig("general", "cloudverifier_port")
	log.Printf("Starting Cloud Verifier on port %v,  use <Ctrl-C> to stop", port)

	router := chi.NewRouter()
	router.Get("/", useAgents)
	router.Route("/agents", func(router chi.Router) {
		router.Get("/{agent_uuid}", getAgents)
		router.Get("/", getAgents)
	})

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
