## Project : url-shortner
## Author  : Ahmad Baderkhan
## Type    : API => GO LANG 
---

### Description

- This is the api for url shortner , it has a post and a get method . 
  It was written in GO and uses mysql as a db in order to fetch the k-v for the small link to the generated link.
- The Id logic works by using uint in the database , but for the front end it's encoded in base 62 so we can get a smaller looking string. 
   - forexample
      - if the db id for the link was 1000001
      - the return value to the front end client would be 4C93 , this way we can access the url as domain.com/4C93 

---

### Requirements 

- In order to run locally , you must have the following Requirements
  - docker
  - docker-compose
  - go version 1.5
---

### Getting Started

1) ```go install```
2) ```make start```

---

### Supported API Routes


#### Creating a Domain 
- Curl
    ``` bash 
      curl --request POST \
           --url http://localhost:8080/links \
           --header 'content-type: application/json' \
           --data '{"url":"<<your url here>>"}' 
    ```
 - Sample Response
    ```json 
    {"id":"1","url":"<<your url here>>"}
    ```

#### Fetching Created Domain 
- Curl
  ```bash
    curl --request GET \
         --url http://localhost:8080/links/1
   ```
- Sample Response 
  ```json 
    {"id":"1","url":"<<your url here>>"}
    ```
