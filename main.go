package main

import (
	"errors"
	"fmt"
)

type Car struct {
	engine string
	color  string
	model  string
}

func (c Car) ToString() string {
	return fmt.Sprintf(`{
"Engine": "%s",
"Color": "%s",
"Model": "%s"
}`, c.engine, c.color, c.model)
}

type CarBuilder interface {
	SetEngine(eng string) CarBuilder
	SetColor(col string) CarBuilder
	SetModel(mod string) CarBuilder
}
type CarObjectBuilder struct {
	engine string
	color  string
	model  string
}

func (c *CarObjectBuilder) SetEngine(eng string) CarBuilder {
	c.engine = eng
	return c
}
func (c *CarObjectBuilder) SetColor(col string) CarBuilder {
	c.color = col
	return c
}
func (c *CarObjectBuilder) SetModel(mod string) CarBuilder {
	c.model = mod
	return c
}
func (c *CarObjectBuilder) GetResult() (Car, error) {
	if c.model == "" || c.engine == "" || c.color == "" {
		return Car{}, errors.New("model , engine and color required")
	}

	return Car{
		engine: c.engine,
		color:  c.color,
		model:  c.model,
	}, nil
}

type CarCommandBuilder struct {
	result    []string
	hasEngine bool
	hasColor  bool
	hasModel  bool
}

func (c *CarCommandBuilder) SetEngine(e string) CarBuilder {
	c.result = append(c.result, "Engine: "+e)
	c.hasEngine = true
	return c
}

func (c *CarCommandBuilder) SetColor(e string) CarBuilder {
	c.result = append(c.result, "Color: "+e)
	c.hasColor = true
	return c
}

func (c *CarCommandBuilder) SetModel(e string) CarBuilder {
	c.result = append(c.result, "Model: "+e)
	c.hasModel = true
	return c
}

func (c *CarCommandBuilder) GetResult() ([]string, error) {
	if !c.hasEngine || !c.hasColor || !c.hasModel {
		return nil, errors.New("model, engine and color required")
	}
	return c.result, nil
}

type Director struct{}

func (d *Director) MakeSportsCar(builder CarBuilder) {
	builder.SetModel("McLaren").SetEngine("V8").SetColor("red")
}

func (d *Director) MakeCityCar(builder CarBuilder) {
	builder.SetModel("Golf").SetEngine("V4").SetColor("blue")
}

func main() {
	director := &Director{}
	object := &CarObjectBuilder{}

	director.MakeSportsCar(object)
	i, err := object.GetResult()
	fmt.Println("1 representation")
	//fmt.Println(i.ToString())
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(i.ToString())
	}

	fmt.Println()

	fmt.Println("2 representation")
	obj := &CarCommandBuilder{}
	director.MakeCityCar(obj)
	res, err := obj.GetResult()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for _, i := range res {
		fmt.Println(i)
	}
}
