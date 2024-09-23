Simple API code to help demonstrate making REST API calls using GoLang

Using Postman, you can run this code, then make a GET api call to your localhost:
http://localhost:8000/account/coins/?username=<username>
pass in the Authorization token (found in mockdb.go) for the corresponding user, and it will return the balance.
Handles errors for improper authorization/user