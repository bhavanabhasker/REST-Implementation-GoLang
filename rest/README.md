Download all the files from github


POST INSTALLATION : 


Run the below dependencies 


go get  github.com/gorilla/mux


go get gopkg.in/mgo.v2



How to execute : 


Step 1 : go build


Run the rest application 


Step 2 : ./rest.exe


OR


Step 1: go run *.go 



Creating the new locations


1. POST/locations 


Go to Postman


Request : 
{
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113"
}


HTTP Response 201


The response will be in the below format,
{
   "id" : 12345,
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113",
   "coordinate" : { 
      "lat": 37.4220352,
      "long": -122.0841244
   }
}


2. GET locations


a. When only locations is entered 


Eg: Request :


 http://localhost:8080/locations
 
 
 HTTP Response 200
 
 
 All the entries in the table will be displayed. 
 
 
 Response 
 [
  {
    "id": 12345,
    "name": "John Smith",
    "address": "1600 Amphitheatre Parkway",
    "city": "Mountain View",
    "state": "CA",
    "zip": "94043",
    "Coordinates": {
      "lat": 37.4220352,
      "long": -122.0841244
    }
  }
]



b. When the location id is also entered 


Request GET /locations/12345


Response :
{
   "id" : 12345,
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113",
   "coordinate" : { 
      "lat": 37.4220352,
      "long": -122.0841244
   }
} 


3. Update a location 



Request : /locations/12345

{
   "address" : "1600 Amphitheatre Parkway",
   "city" : "Mountain View",
   "state" : "CA",
   "zip" : "94043"
}


Response 


HTTP Response 201


{
   "id" : 12345,
   "name" : "John Smith",
   "address" : "1600 Amphitheatre Parkway",
   "city" : "Mountain View",
   "state" : "CA",
   "zip" : "94043"
   "coordinate" : { 
      "lat" : 37.4220352,
     "lng" : -122.0841244
   }
}


4. Delete a location 


Request :


 /locations/12345
 
 
Response :


HTTP Response 200





