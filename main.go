package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mario/nyoba/models"
	"mario/nyoba/user"
	"mario/nyoba/utils"
	"net/http"
	"strconv"
)

func main(){

	http.HandleFunc("/user", GetUser)
	http.HandleFunc("/user/create", PostUser)
	http.HandleFunc("/user/update", UpdateUser)
	http.HandleFunc("/user/delete", DeleteUser)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

// GetUser ...
func GetUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		users, err := user.GetAll(ctx)
		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, users, http.StatusOK)
		return
	}
	http.Error(w, "Permission Rejected", http.StatusNotFound)
	return
}

// PostUser ...
func PostUser(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		if r.Header.Get("Content-Type") != "application/json"{
			http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var use models.User

		if err:= json.NewDecoder(r.Body).Decode(&use); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := user.Insert(ctx, use); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Not Permmited", http.StatusMethodNotAllowed)
	return
}

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "PUT" {
 
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
            return
        }
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var use models.User
 
        if err := json.NewDecoder(r.Body).Decode(&use); err != nil {
            utils.ResponseJSON(w, err, http.StatusBadRequest)
            return
        }
 
        fmt.Println(use)
 
        if err := user.Update(ctx, use); err != nil {
            utils.ResponseJSON(w, err, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusCreated)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}

// DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "DELETE" {
 
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
            return
        }
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var use models.User
		
		id := r.URL.Query().Get("id")
 
        if id == "" {
            utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
            return
        }
        use.ID, _ = strconv.Atoi(id)
 
        if err := user.Delete(ctx, use); err != nil {
            utils.ResponseJSON(w, err, http.StatusInternalServerError)
            return
        }
 
        res := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, res, http.StatusCreated)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}