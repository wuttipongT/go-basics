package restapi

/*

# Project Description
A Go-woered "Event Booking" REST API
GET /events									Get a list of avaiable events
GET /events/<id>							Get a list of avaiable events
POST /events								Create a new bookabel event | Auth Required
PUT /events/<id>							Update an event | Auth Required | Only by creator
DELETE /events/<id>							Delete an event | Auth Required | Only by creator
POST /singup								Create new user
POST /login									Authenticate user | Auth Token(JWT)
POST /events/<id>/register					Register user for event | Auth Required
DELETE /events/<id>/register				Cancel registration | Auth Required

*/

func New() {

}

func main() {

}
