package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const MONGO_DB_NAME string = "cat-api-dev"
const CAT_COLLECTION string = "cats"

type Cat struct {
	Id          bson.ObjectId `json:"id" bson:"_id",omitempty`
	Name        string        `json:"name" bson:"name",omitempty`
	Age         int32         `json:"age" bson: "age",omitempty`
	AdoptionFee float32       `json:"adoptionFee" bson: "adoptionFee", omitempty`
	Image       string        `json:"image" bson: "image", omitempty`
	Likes       []string      `json:"likes" bson: "likes", omitempty`
	Dislikes    []string      `json:"dislikes" bson: "dislikes", omitempty`
}

type CatAPI struct {
	*iris.Context
}

// GET /api/cats
func (c CatAPI) Get() {
	var cats []Cat
	err := catDb().Find(nil).All(&cats)
	if err != nil {
		panic(err)
	}

	if cats == nil {
		cats = []Cat{}
	}

	c.JSON(iris.StatusOK, cats)
}

// GET /api/cats/:param1 which its value passed to the id argument
func (c CatAPI) GetBy(id string) { // id equals to c.Param("param1")
	c.Write("Get from /api/cats/%s", id)

	oid := bson.ObjectIdHex(id)

	cat := Cat{}
	if err := catDb().FindId(oid).One(&cat); err != nil {
		c.NotFound()
		return
	}

	c.JSON(iris.StatusOK, cat)
}

// POST /api/cats
func (c CatAPI) Post() {
	cat := &Cat{}
	err := c.ReadJSON(cat)
	if err != nil {
		panic(err)
	}

	cat.Id = bson.NewObjectId()

	if err := catDb().Insert(cat); err != nil {
		c.JSON(iris.StatusInternalServerError, err)
	} else {
		c.JSON(iris.StatusCreated, cat)
	}
}

// PUT /api/cats/:param1
func (c CatAPI) PutBy(id string) {
	oid := bson.ObjectIdHex(id)

	cat := &Cat{}
	err := c.ReadJSON(cat)
	if err != nil {
		panic(err)
	}

	cat.Id = oid

	if err := catDb().UpdateId(oid, cat); err != nil {
		c.JSON(iris.StatusInternalServerError, err)
	} else {
		c.JSON(iris.StatusOK, cat)
	}
}

// DELETE /api/cats/:param1
func (c CatAPI) DeleteBy(id string) {
	oid := bson.ObjectIdHex(id)

	if err := catDb().RemoveId(oid); err != nil {
		c.JSON(iris.StatusInternalServerError, err)
	} else {
		c.JSON(iris.StatusNoContent, nil)
	}
}

func seedDb(c *iris.Context) {
	c.JSON(iris.StatusNotImplemented, nil)
}

func catDb() *mgo.Collection {
	return getSession().DB(MONGO_DB_NAME).C(CAT_COLLECTION)
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27017/" + MONGO_DB_NAME)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	return s
}

func Logger(c *iris.Context) {
	println(c.MethodString() + " - " + c.PathString())
	c.Next()
}

func main() {
	iris.UseFunc(Logger)
	iris.API("/api/cats", CatAPI{})
	iris.Post("/api/cats/reset", seedDb)

	iris.Config.Render.Template.Engine = config.MarkdownEngine
	iris.Config.Render.Template.Extensions = []string{".md"}
	iris.Config.Render.Template.Directory = "./"

	iris.Get("/", func(ctx *iris.Context) {
		err := ctx.Render("readme.md", nil)
		if err != nil {
			panic(err)
		}
	})

	iris.Listen(":8080")
}
