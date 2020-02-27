// Copyright 2020 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	version = "v0.1.0"
)

type Response struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Messages []string `json:"messages"`
	Status   int      `json:"status"`
}

type VersionResponse struct {
	Version string `json:"version"`
}

func main() {
	log.Println("Starting the twelve service...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message: "Hello world!",
		}

		data, err := json.MarshalIndent(&response, "", " ")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(data)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		response := HealthResponse{
			Messages: []string{},
			Status:   200,
		}

		data, err := json.MarshalIndent(&response, "", " ")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(data)
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		response := VersionResponse{
			Version: version,
		}

		data, err := json.MarshalIndent(&response, "", " ")
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(data)
	})

	s := http.Server{Addr: ":8080"}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}
