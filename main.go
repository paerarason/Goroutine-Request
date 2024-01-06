package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"sync"
	"context"
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

	URL := "https://webhook.site/197aa273-0162-43d0-9234-40adf4dfff58" 
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("ERROR",err)
		return
	}

	defer resp.Body.Close()
	log.Println("sendToWebhook Finished",resp.Status)
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		defer r.Body.Close()
		ctx:=context.Background()
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err!=nil{
			log.Println("Failed to decode request body:", err)
		}
		ctx=context.WithValue(ctx,"body",bodyBytes)
        var wg sync.WaitGroup

		wg.Add(1)
		go func (ctx context.Context,w http.ResponseWriter){
			  var bodyData map[string]interface{}
			  if body := ctx.Value("body"); body!= nil {
			  err := json.Unmarshal( body.([]byte), &bodyData)
			  if err != nil {
				  	log.Println("Failed to decode request body:", err)
			     }
				 transformedData := transformData(bodyData)
				 sendToWebhook(transformedData)
				 w.WriteHeader(http.StatusOK)
				 w.Write([]byte("Request Recieved"))	
				 wg.Done()
				return
		      }
			 log.Println("BODY EMPTY")
		}(ctx,w)
        //GONE
	    wg.Wait()
	    })
       
		
	http.ListenAndServe(":8001", nil)
}