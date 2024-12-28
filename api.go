package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// write json
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// API error type
type ApiError struct {
	Error string
}

// This is the function signature of the function we are using
type apiFunc func(http.ResponseWriter, *http.Request) error

// this is going to decorate our API func and to an HTTP handler function
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// the f is one of the functions of the APIServer that we define below. For example: HandleAccount, etc.
		if err := f(w, r); err != nil {
			// Normally, handle func does not return anything
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIserver(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

// function to start and server up
func (s *APIServer) Run() {
	router := mux.NewRouter()

	// s.handleAccount without parentheses is passing a reference to the function
	// This allows makeHTTPHandleFunc to call the function later when needed.
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountById))

	log.Println("JSON API server running on port: ", s.listenAddr)

	// Starts an HTTP server that listens for incoming requests on the specified
	// address (s.listenAddr) and handles them using the provided handler (router).
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreteAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	/*
		The fmt package in Go provides formatted I/O functions,
		similar to printf and scanf in C. It is used for:
		Formatting strings.
		Printing to the console.
		Creating formatted errors
	*/
	return fmt.Errorf("Method not allowed %s", r.Method)
	// %s -> formatted as a string
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	// // Extracts route variables (parameters) from the HTTP request r.
	// vars := mux.Vars(r)

	// // Retrieves the value of a specific route variable (in this case, id) from the map returned by mux.Vars(r).
	// id := mux.Vars(r)["id"]
	// fmt.Fprintf(w, "Account ID: %s", id)

	account := NewAccount("John", "Doe")

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreteAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	fmt.Printf("Created account: %+v\n", account)

	// Return the created account as a JSON response
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
