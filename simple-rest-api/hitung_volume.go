package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Sisi struct {
	Panjang     int    `json:"panjang"`
	Lebar       int    `json:"lebar"`
	Tinggi      int    `json:"tinggi"`
}

type Hasil struct {
	JenisBangun string `json:"jenis_bangun"`
	Volume        int  `json:"volume"`
}

func main() {
	// Mendefine alamat API
	router := mux.NewRouter()
	router.HandleFunc("/api/hitung-volume", Volume)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Memanggil fungsi Volume
func Volume(w http.ResponseWriter, r *http.Request) {
	var sisi Sisi

	if r.Method != "POST" {
		WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		WrapAPIError(w, r, "can't read body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &sisi)
	if err != nil {
		WrapAPIError(w, r, "error Unmarshal :"+err.Error(), http.StatusInternalServerError)
		return
	}

	// Memanggil function untuk memberikan response berhasil
	WrapAPIData(w, r, Hasil{
		JenisBangun: sisi.JenisBangun(),
		Volume:      sisi.RumusVolume(),
	}, "success", http.StatusOK)

}

// Mengetahui jenis bangun ruang
func (s *Sisi) JenisBangun() string {
	if (s.Panjang == s.Lebar && s.Lebar == s.Tinggi) {
		return "Kubus"
	} 
	return "Balok"
	
}

// rumus menghitung Volume bangun ruang
func (s *Sisi) RumusVolume() int {
	return s.Panjang * s.Lebar * s.Tinggi
}

// Untuk memberikan Response Error
func WrapAPIError(w http.ResponseWriter, r *http.Request, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	result, err := json.Marshal(map[string]interface{}{
		"code":          code,
		"error_type":    http.StatusText(code),
		"error_details": message,
	})
	if err == nil {
		w.Write(result)
	} else {
		log.Println(fmt.Sprintf("error wrap API error : %s", err))
	}
}

// Untuk memberikan Response yang berhasil
func WrapAPIData(w http.ResponseWriter, r *http.Request, data interface{}, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	result, err := json.Marshal(map[string]interface{}{
		"code":   code,
		"status": message,
		"data":   data,
	})
	if err == nil {
		w.Write(result)
	} else {
		log.Println(fmt.Sprintf("error wrap API error : %s", err))
	}
}
