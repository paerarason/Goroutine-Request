# Goroutine-Request

### To run the Appliction
```bash
go run main.go
```
### API TESTING 
POST:8001
```
http
http://localhost:8001/
```
#### Sample Input 1:
```bash
{
"ev": "contact_form_submitted",
"et": "form_submit",
"id": "cl_app_id_001",
"uid": "cl_app_id_001-uid-001",
"mid": "cl_app_id_001-uid-001",
"t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"l": "en-US",
"sc": "1920 x 1080",
"atrk1": "form_varient",
"atrv1": "red_top",
"atrt1": "string",
"atrk2": "ref",
"atrv2": "XPOWJRICW993LKJD",
"atrt2": "string",
"uatrk1": "name",
"uatrv1": "iron man",
"uatrt1": "string",
"uatrk2": "email",
"uatrv2": "ironman@avengers.com",
"uatrt2": "string",
"uatrk3": "age",
"uatrv3": "32",
"uatrt3": "integer"
}
```
#### Sample Output 1:
```bash
{
"event": "contact_form_submitted",
"event_type": "form_submit",
"app_id": "cl_app_id_001",
"user_id": "cl_app_id_001-uid-001",
"message_id": "cl_app_id_001-uid-001","page_title": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"page_url": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"browser_language": "en-US",
"screen_size": "1920 x 1080",
"attributes": { "form_varient": {
"value": "red_top",
"type": "string" },
"ref": {
"value": "XPOWJRICW993LKJD",
"type": "string"} },
"traits": {
"name": {
"value": "iron man",
"type": "string"},
"email": {
"value": "ironman@avengers.com",
"type": "string"},
"age": {
"value": "32",
"type": "integer"}}
}
```

#### Webhook Response
![OUTPUT 1](static/Output1.png)


#### Sample Input 2:
```bash
{
"ev": "top_cta_clicked",
"et": "clicked",
"id": "cl_app_id_001",
"uid": "cl_app_id_001-uid-001",
"mid": "cl_app_id_001-uid-001",
"t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"l": "en-US",
"sc": "1920 x 1080",
"atrk1": "button_text",
"atrv1": "Free trial",
"atrt1": "string",
"atrk2": "color_variation",
"atrv2": "ESK0023",
"atrt2": "string",
"uatrk1": "user_score",
"uatrv1": "1034",
"uatrt1": "number",
"uatrk2": "gender",
"uatrv2": "m",
"uatrt2": "string",
"uatrk3": "tracking_code","uatrv3": "POSERK093",
"uatrt3": "string"
}
```
#### Sample Output 2:
```bash
{
"event": "top_cta_clicked",
"event_type": "clicked",
"app_id": "cl_app_id_001",
"user_id": "cl_app_id_001-uid-001",
"message_id": "cl_app_id_001-uid-001",
"page_title": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"page_url": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"browser_language": "en-US",
"screen_size": "1920 x 1080",
"attributes": {
"button_text": {
"value": "Free trial",
"type": "string"},
"color_variation": {
"value": "ESK0023",
"type": "string"}},
"traits": {
"user_score": {
"value": "1034",
"type": "number"},
"gender": {
"value": "m",
"type": "string"},
"tracking_code": {
"value": "POSERK093",
"type": "string"}}
}
```

#### Webhook Response
![OUTPUT 2](static/Output2.png)



#### Sample Input 3:
```bash
{
"ev": "top_cta_clicked",
"et": "clicked",
"id": "cl_app_id_001",
"uid": "cl_app_id_001-uid-001",
"mid": "cl_app_id_001-uid-001",
"t": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"p": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"l": "en-US",
"sc": "1920 x 1080",
"atrk1": "button_text",
"atrv1": "Free trial",
"atrt1": "string",
"atrk2": "color_variation",
"atrv2": "ESK0023",
"atrt2": "string",
"atrk3": "page_path",
"atrv3": "/blog/category_one/blog_name.html",
"atrt3": "string",
"atrk4": "source",
"atrv4": "facebook",
"atrt4": "string",
"uatrk1": "user_score",
"uatrv1": "1034",
"uatrt1": "number",
"uatrk2": "gender",
"uatrv2": "m",
"uatrt2": "string",
"uatrk3": "tracking_code",
"uatrv3": "POSERK093",
"uatrt3": "string",
"uatrk4": "phone",
"uatrv4": "9034432423",
"uatrt4": "number",
"uatrk5": "coupon_clicked",
"uatrv5": "true",
"uatrt5": "boolean",
"uatrk6": "opt_out","uatrv6": "false",
"uatrt6": "boolean"
}
```
#### Sample Output 3:
```bash
{
"event": "top_cta_clicked",
"event_type": "clicked",
"app_id": "cl_app_id_001",
"user_id": "cl_app_id_001-uid-001",
"message_id": "cl_app_id_001-uid-001",
"page_title": "Vegefoods - Free Bootstrap 4 Template by Colorlib",
"page_url": "http://shielded-eyrie-45679.herokuapp.com/contact-us",
"browser_language": "en-US",
"screen_size": "1920 x 1080",
"attributes": {
"button_text": {
"value": "Free trial",
"type": "string"},
"color_variation": {
"value": "ESK0023",
"type": "string"},
"page_path": {
"value": "/blog/category_one/blog_name.html",
"type": "string"},
"source": {
"value": "facebook",
"type": "string"}},
"traits": {
"user_score": {
"value": "1034",
"type": "number"},
"gender": {
"value": "m",
"type": "string"},
"tracking_code": {
"value": "POSERK093",
"type": "string"},
"phone": {
"value": "9034432423",
"type": "number"},
"coupon_clicked": {
"value": "true",
"type": "boolean"},
"opt_out": {
"value": "false",
"type": "boolean"}}
}
```

#### Webhook Response
![OUTPUT 3](static/output3.png)