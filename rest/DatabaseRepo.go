package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var todos Locations

// This struct contains selected fields from Google's Geocoding Service response
type googleGeocodeResponse struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64
				Lng float64
			}
		}
	}
}

// Find from the db by ID
func RepoFindTodo(id int) Location {
	session, err := mgo.Dial("mongodb://bhavana:bhavana@ds037244.mongolab.com:37244/tests")
	c := session.DB("tests").C("locations")
	result := Location{}
	err = c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	// return empty Todo if not found
	return result
}

// Fetch all the results for the handler 'Get/'
func RepoFindAll() []Location {
	session, err := mgo.Dial("mongodb://bhavana:bhavana@ds037244.mongolab.com:37244/tests")
	c := session.DB("tests").C("locations")
	result := []Location{}
	err = c.Find(nil).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// Handler for the POST
func RepoCreateTodo(t Location) Location {
	var data googleGeocodeResponse
	session, err := mgo.Dial("mongodb://bhavana:bhavana@ds037244.mongolab.com:37244/tests")

	//todos = append(todos, t)
	c := session.DB("tests").C("locations")
	result := []Location{}
	err = c.Find(nil).All(&result)
	if len(result) == 0 {
		t.Id = 12345
	} else {
		t.Id = 12345 + len(result)
	}
	address := t.Address
	address += t.City
	address += t.State
	address += t.Zip
	connection := "https://maps.googleapis.com/maps/api/geocode/json"
	connection += "?address="
	url_safe_query := url.QueryEscape(address)
	connection += url_safe_query
	connection += "&key=AIzaSyBYfZue8Wu5PW65uJN05n6a-uTzq7gQmOc"
	req, _ := http.NewRequest("GET", connection, nil)
	client := &http.Client{}
	resp, requestErr := client.Do(req)
	if requestErr != nil {
		log.Fatal(err)
	}
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			json.Unmarshal(body, &data)
		}
	}
	t.Coordinates.Lat = data.Results[0].Geometry.Location.Lat
	t.Coordinates.Long = data.Results[0].Geometry.Location.Lng
	err = c.Insert(&Location{Id: t.Id, Name: t.Name, Address: t.Address,
		City: t.City, State: t.State, Zip: t.Zip, Coordinates: t.Coordinates})
	if err != nil {
		log.Fatal(err)
	}
	return t
}
func RepoUpdateTodo(id int, t Location) Location {
	var data googleGeocodeResponse
	session, err := mgo.Dial("mongodb://bhavana:bhavana@ds037244.mongolab.com:37244/tests")

	c := session.DB("tests").C("locations")
	result := Location{}
	err = c.Find(bson.M{"id": id}).One(&result)
	t.Id = id
	var address string
	if t.Address != "" {
		address = t.Address
	} else {
		address = result.Address
		t.Address = result.Address
	}
	if t.City != "" {
		address += t.City
	} else {
		address += result.City
		t.City = result.City
	}
	if t.State != "" {
		address += t.State
	} else {
		address += result.State
		t.State = result.State
	}
	if t.Zip != "" {
		address += t.Zip
	} else {
		address += result.Zip
		t.Zip = result.Zip
	}
	connection := "https://maps.googleapis.com/maps/api/geocode/json"
	connection += "?address="
	url_safe_query := url.QueryEscape(address)
	connection += url_safe_query
	connection += "&key=AIzaSyBYfZue8Wu5PW65uJN05n6a-uTzq7gQmOc"
	req, _ := http.NewRequest("GET", connection, nil)
	client := &http.Client{}
	resp, requestErr := client.Do(req)
	if requestErr != nil {
		log.Fatal(err)
	}
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			json.Unmarshal(body, &data)
		}
	}

	t.Coordinates.Lat = data.Results[0].Geometry.Location.Lat
	t.Coordinates.Long = data.Results[0].Geometry.Location.Lng

	err = c.Remove(bson.M{"id": t.Id})
	if err != nil {
		log.Fatal(err)
	}
	if t.Name == "" {
		t.Name = result.Name
	}
	err = c.Insert(&Location{Id: t.Id, Name: t.Name, Address: t.Address,
		City: t.City, State: t.State, Zip: t.Zip, Coordinates: t.Coordinates})
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func RepoDestroyTodo(id int) error {
	session, err := mgo.Dial("mongodb://bhavana:bhavana@ds037244.mongolab.com:37244/tests")
	c := session.DB("tests").C("locations")
	err = c.Remove(bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
