// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Hello struct {
	Name      string `json:"name"`
	Secondary string `json:"secondary"`
}

func (Hello) IsEntity() {}

type HelloMultiSingleKeys struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func (HelloMultiSingleKeys) IsEntity() {}

type HelloWithErrors struct {
	Name string `json:"name"`
}

func (HelloWithErrors) IsEntity() {}

type MultiHello struct {
	Name string `json:"name"`
}

func (MultiHello) IsEntity() {}

type MultiHelloByNamesInput struct {
	Name string `json:"Name"`
}

type MultiHelloMultipleRequires struct {
	Name string `json:"name"`
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	Key3 string `json:"key3"`
}

func (MultiHelloMultipleRequires) IsEntity() {}

type MultiHelloMultipleRequiresByNamesInput struct {
	Name string `json:"Name"`
}

type MultiHelloMultipleRequiresWithKey1AndKey2sInput struct {
	Key1 string `json:"Key1"`
	Key2 string `json:"Key2"`
}

type MultiHelloRequires struct {
	Name string `json:"name"`
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func (MultiHelloRequires) IsEntity() {}

type MultiHelloRequiresByNamesInput struct {
	Name string `json:"Name"`
}

type MultiHelloRequiresWithKey1sInput struct {
	Key1 string `json:"Key1"`
}

type MultiHelloWithError struct {
	Name string `json:"name"`
}

func (MultiHelloWithError) IsEntity() {}

type MultiHelloWithErrorByNamesInput struct {
	Name string `json:"Name"`
}

type MultiPlanetRequiresNested struct {
	Name  string `json:"name"`
	World *World `json:"world"`
	Size  int    `json:"size"`
}

func (MultiPlanetRequiresNested) IsEntity() {}

type MultiPlanetRequiresNestedByNamesInput struct {
	Name string `json:"Name"`
}

type MultiPlanetRequiresNestedWithWorldFoosInput struct {
	WorldFoo string `json:"WorldFoo"`
}

type PlanetMultipleRequires struct {
	Name     string `json:"name"`
	Diameter int    `json:"diameter"`
	Density  int    `json:"density"`
	Weight   int    `json:"weight"`
}

func (PlanetMultipleRequires) IsEntity() {}

type PlanetRequires struct {
	Name     string `json:"name"`
	Size     int    `json:"size"`
	Diameter int    `json:"diameter"`
}

func (PlanetRequires) IsEntity() {}

type PlanetRequiresNested struct {
	Name   string   `json:"name"`
	World  *World   `json:"world"`
	Worlds []*World `json:"worlds,omitempty"`
	Size   int      `json:"size"`
	Sizes  []int    `json:"sizes,omitempty"`
}

func (PlanetRequiresNested) IsEntity() {}

type Query struct {
}

type World struct {
	Foo   string `json:"foo"`
	Bar   int    `json:"bar"`
	Hello *Hello `json:"hello,omitempty"`
}

func (World) IsEntity() {}

type WorldName struct {
	Name string `json:"name"`
}

func (WorldName) IsEntity() {}

type WorldWithMultipleKeys struct {
	Foo   string `json:"foo"`
	Bar   int    `json:"bar"`
	Hello *Hello `json:"hello,omitempty"`
}

func (WorldWithMultipleKeys) IsEntity() {}
