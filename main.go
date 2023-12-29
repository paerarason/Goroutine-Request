package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
)

//Attibute struct
type attribute struct {
	Value string `json:"value"`
	Type string `json:"type"`
} 

//Attibute struct
type OutgoingData struct {
	Event string `json:"event"`
	EventType string `json:"event_type"`
	AppID string `json:"app_id"`
	UserID string `json:"user_id"`
	MessageID string `json:"message_id"`
	PageTitle string `json:"page_title"`
	PageURL string `json:"page_url"`
	BrowserLanguage string `json:"browser_language"`
	ScreenSize string `json:"screen_size"`
	Attributes map[string]attribute `json:"attributes"`
	UserAttributes map[string]attribute `json:"traits"`
}

type HTTPRequest struct {
	Request *http.Request
	Info    string 
}



//worker Go-routine 
func worker(ReqChannel <-chan HTTPRequest) {
	log.Println("Worker started....")
	for {
	data := <-ReqChannel
	body, err := ioutil.ReadAll(data.Request.Body)
		if err != nil {
			fmt.Println("Failed to read request body:", err)
			continue
		}
		defer data.Request.Body.Close()

		// Process the request body
		var bodyData map[string]interface{}
		err = json.Unmarshal(body, &bodyData)
		if err != nil {
			fmt.Println("Failed to decode request body:", err)
			continue
		}
	transformedData := transformData(bodyData)
	sendToWebhook(transformedData)
	log.Println("Worker Finished....")
}
}

func KeySearch(data map[string]interface{},searchKey string ) string{
	if value, ok := data[searchKey]; ok {
	     return value.(string)
	} 
	     return ""
	}

func transformData(data  map[string]interface{}) OutgoingData {
	
	transformed := OutgoingData{
		Event: KeySearch(data,"ev"),
		EventType: KeySearch(data,"et"),
		AppID: KeySearch(data,"id"),
		UserID: KeySearch(data,"uid"),
		MessageID: KeySearch(data,"mid"),
		PageTitle: KeySearch(data,"t"),
		PageURL: KeySearch(data,"p"),
		BrowserLanguage: KeySearch(data,"l"),
		ScreenSize: KeySearch(data,"sc"),
		Attributes: make(map[string]attribute),
		UserAttributes: make(map[string]attribute),
	}
    
	//Map Attibute
    for key, value := range data {
		if len(key)>4{
            if strings.Contains(key[0:4],"atrk"){
            values:=KeySearch(data,"atrv"+key[4:])
		    types:=KeySearch(data,"atrt"+key[4:])
			attr:=attribute{
		          Value:values,
		          Type :types,	
                }
            transformed.Attributes[value.(string)]=attr
		}
		}
		
	}
   
	//Map UserAttibute
	for key, value := range data {
		if strings.Contains(key,"uatrk"){
            values:=KeySearch(data,"uatrv"+key[5:])
		    types:=KeySearch(data,"uatrt"+key[5:])
			attr:=attribute{
		          Value:values,
		          Type :types,	
                }
          transformed.UserAttributes[value.(string)]=attr
		}
	}
return transformed
}



func sendToWebhook(data OutgoingData) {
	
	log.Println("sendToWebhook Finished....")
	jsonData, err := json.Marshal(data)
	
	if err != nil {
	fmt.Println("Error encoding JSON:", err)
	return
	}

	URL := "https://webhook.site/1f79959d-004c-41ad-b968-44d8e914b4a6" 
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("ERROR",err)
		return
	}

	defer resp.Body.Close()
	log.Println("sendToWebhook Finished",resp.Status)
}


func main() {
    
	ReqChannel := make(chan HTTPRequest)
    go worker(ReqChannel)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		    defer r.Body.Close()
            body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}
			r.Body = ioutil.NopCloser(strings.NewReader(string(body)))
        
		request:=HTTPRequest{
				Request:r,
				Info:"INCOMMING REQUEST",
		   }
		    ReqChannel <-request
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Request Recieved"))
		})

	http.ListenAndServe(":8001", nil)
}