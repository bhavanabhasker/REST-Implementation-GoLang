# REST-Implementation-GoLang
Sample REST interface application in golang 

##Objective 
This application enables the user to store the personal information of his clients/family. 
The information may contain address details including the latitude and longitude details. 

##How to Execute?

1. Download/Clone the files from the github repo 
2. Run the following dependencies 
<pre>
go get github.com/gorilla/mux

go get gopkg.in/mgo.v2

</pre>
3. Execute the below commands after step 2
<pre>
 go build
 ./rest.exe
 go run *.go
</pre>

## Request and Responses 

1. use the route /POST/locations for creating new user
Sample request and response as shown below,
<pre>
/POST/locations
{ "name" : "John Smith", "address" : "123 Main St", "city" : "San Francisco", "state" : "CA", "zip" : "94113" }

The response:
{ "id" : 12345, "name" : "John Smith", "address" : "123 Main St", "city" : "San Francisco", "state" : "CA", "zip" : "94113", "coordinate" : { "lat": 37.4220352, "long": -122.0841244 } }
</pre>

2. use the route GET/locations to get a list of all the saved records 
<pre>
http://localhost:8080/locations
Response:
[ { "id": 12345, "name": "John Smith", "address": "1600 Amphitheatre Parkway", "city": "Mountain View", "state": "CA", "zip": "94043", "Coordinates": { "lat": 37.4220352, "long": -122.0841244 } } ]
</pre>

3. The route GET/locations/<locationID> can be used to retrieve a specific record from the database 
<pre>
GET /locations/12345

Response:
{ "id" : 12345, "name" : "John Smith", "address" : "123 Main St", "city" : "San Francisco", "state" : "CA", "zip" : "94113", "coordinate" : { "lat": 37.4220352, "long": -122.0841244 } }
</pre>

4. Use PUT to update the entry in the database 
<pre>
/locations/12345
{ "address" : "1600 Amphitheatre Parkway", "city" : "Mountain View", "state" : "CA", "zip" : "94043" }
Response:
{ "id" : 12345, "name" : "John Smith", "address" : "1600 Amphitheatre Parkway", "city" : "Mountain View", "state" : "CA", "zip" : "94043" "coordinate" : { "lat" : 37.4220352, "lng" : -122.0841244 } }
</pre>
5. An entry in the database can be deleted with /DELETE 
<pre>
/locations/12345
Response:
HTTP Response 200
</pre>


## Application Framework and Additional Notes

1. Client dispatches the request to server using the REST calls 
2. Gorilla MUX is used for dispatching/handling the requests 
3. Mongo DB is used to persist the client information. 
