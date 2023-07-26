package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TextToImageImage struct {
	Base64       string `json:"base64"`
	Seed         uint32 `json:"seed"`
	FinishReason string `json:"finishReason"`
}

type TextToImageResponse struct {
	Images []TextToImageImage `json:"artifacts"`
}

func TextToImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	engineId := vars["engine_id"]

	textPrompt := r.FormValue("textPrompt") // get the text from the form data

	data := map[string]interface{}{
		"text_prompts": []map[string]string{
			{
				"text": textPrompt,
			},
		},
		"cfg_scale":            7,
		"clip_guidance_preset": "FAST_BLUE",
		"height":               512,
		"width":                512,
		"samples":              1,
		"steps":                30,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := baseURL + "/v1/generation/" + engineId + "/text-to-image"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var resBody map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(resp.StatusCode)
		json.NewEncoder(w).Encode(resBody)
		return
	}

	var resBody TextToImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&resBody); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resBody)
}
