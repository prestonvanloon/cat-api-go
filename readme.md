# Cat API
A simple API that allows for demonstration of a CRUD model.

## Cat Model

```
Cat = {
  _id: MongoID,
  __v: MongoVersionKey
  name: String,
  age: Number,
  adoptionFee: Number,
  image: String,
  likes: [ String ],
  dislikes: [ String ]
}
```
Example cat:
```
{
  "_id": "57489ed5f88d18f12128b25b",
  "name": "James",
  "image": "http://25.media.tumblr.com/tumblr_lwib0wCyny1qbhms5o1_500.jpg",
  "age": 5,
  "adoptionFee": 54.33,
  "__v": 0,
  "dislikes": [
    "water",
    "vacuums"
  ],
  "likes": [
    "yarn",
    "lasers"
  ]
}
```

## Routes
### List cats
 - GET /api/cats

Returns an array of cats.

### Create a cat
- POST /api/cats

Creates a cat. The payload should be in the same format as the Cat model. Returns newly created cat.

### Get a cat
 - GET /api/cats/:id

Returns a single cat or 404 if no cat found.

### Update a cat
- PUT /api/cats/:id

Replaces a cat with the request payload. The payload should be in the same format as the Cat model. Returns updated cat.

### Delete a cat
- DELETE /api/cats/:id

Deletes a cat. Returns status 204.

### Reset DB and seed cats
- POST /api/cats/reset

This endpoint removes all existing cats and seeds the database with some new cats. No body is required. Returns status code 201.

## Running and developing
### Requirements
  - MongoDB
  - Go

### Installation
```
go get -u github.com/kataras/iris/iris
go get go get gopkg.in/mgo.v2
```

### Running the server
```
go run main.go
```
The server will run on localhost:8080 by default.
