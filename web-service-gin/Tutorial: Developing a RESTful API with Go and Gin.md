This tutorial introduces the basics of writing a RESTful web service API with Go and the Gin Web Framework (Gin).

Gin simplifies many coding tasks associated with building web applications, including web services. In this tutorial, you'll use Gin to route requests, retrieve request details, and marshal JSON for responses.

# Design API endpoints
You'll build an API that provides access to store selling vintage recordings on vinyl. So you'll need to provide endpoints through which a client can get and add albums for users. 

When developing an API, you typically begin by designing the endpoints. Your API's users will have more success if the endpoints are easy to understand. 

Here are the endpoints you'll create in this tutorial. 

/albums
- GET - Get a list of all albums, returned as JSON.
- POST - Add a new album from request data sent as JSON. 

/albums/:id
- GET - Get an album by its ID, returning the album data as JSON. 

# Create a folder for your code 

To begin, create a project for the code you'll write.

# Create the data 
To keep things simple for the tutorial, you'll store data in memory. A more typical API would interact whit a database.