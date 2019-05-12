//----------------------------------------------------------------------------
// ZGork - a playful imitation of Zork.
// Copyright (C) 2018 Michael D Henderson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//----------------------------------------------------------------------------

// Package main implements the web server for ZGork.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mdhender/zgork/xii"
)

func main() {
	// load the environment
	port, err := xii.AsInt("PORT", xii.Int{Required: false, DefaultValue: 8080, Help: "port to serve on"})
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}

	http.Handle("/command", &CommandController{})
	http.Handle("/", http.FileServer(http.Dir("./public")))

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// serve the content.
	log.Printf("listening on port %d; press Ctrl+C to terminate...\n", port)
	log.Fatal(srv.ListenAndServe())
}

// CommandController is
type CommandController struct{}

func (ctl *CommandController) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		postCommand(rw, req)
	default:
		http.Error(rw, "method not allowed", http.StatusBadRequest)
	}
}

func postCommand(rw http.ResponseWriter, req *http.Request) {
	r := struct {
		ID int `json:"id"`
	}{
		ID: 5,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(r)
}
