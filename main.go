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
