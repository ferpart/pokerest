# helpers

## Table of Contents

- [Package](#package)
- [Functions](#functions)
- [Types](#types)

## Package

```go
package helpers // import "github.com/ferpart/pokerest/helpers"
```

Package helpers stores all the helper functions that are used for parsing
requests to the api.

## FUNCTIONS

```go
func GetLanguage(res interface{}) string
```

GetLanguage verifies if language exists and returns the name of the
language. In case the language does not exist, it will return an empty
string.

```go
func SendJSON(w http.ResponseWriter, status int, body interface{})
```

SendJSON recieves a ResponseWriter, a status, and a body then sends a json
to the ResponseWriter

```go
func writeHeader(w *http.ResponseWriter, status int)
```

writeHeader writes the headers required for setting the application to type json and allowing cors

```go
func GetDamageRelations(url string) Drelations
```

GetDamageRelations gets da damage relations from the passed url and returns
it in a Drelations struct

```go
func GetMove(url string, lan string) Elem
```

GetMove gets a move with the provided url and translates it to the language
given in lan

```go
func GetPokemon(pokeNames []string) []Pokemon
```

GetPokemon given an array of two pokemon, will return an array of type
Pokemon

## TYPES

```go
type CompleteMove struct {
        Name  string     `json:"name"`
        Names []language `json:"names"`
}
```

CompleteMove saves the name, and all the names in different locales of a
move

```go
type Drelations struct {
        Name      string    `json:"name"`
        Drelation drelation `json:"damage_relations"`
}
```

Drelations stores a type, and the damage relations of said type

```go
type Elem struct {
        Name string `json:"name"`
        URL  string `json:"url"`
}
```

Elem stores a name as well as a url

```go
type Move struct {
        Move Elem `json:"move"`
}
```

Move stores an Elem containing a move's data

```go
type Pokemon struct {
        Name   string  `json:"name"`
        Ptypes []ptype `json:"types"`
        Moves  []Move  `json:"moves"`
}
```

Pokemon stores a pokemon's name, type, and moveset

```go
type drelation struct {
        DoubleDfrom []Elem `json:"double_damage_from"`
        DoubleDto   []Elem `json:"double_damage_to"`
        HalfDfrom   []Elem `json:"half_damage_from"`
        HalfDto     []Elem `json:"half_damage_to"`
        NoDfrom     []Elem `json:"no_damage_from"`
        NoDto       []Elem `json:"no_damage_to"`
}
```

drelation is a structure that stores all the damages made from, and two the
type in question

```go
type language struct {
        Name     string `json:"name"`
        Language Elem   `json:"language"`
}
```

language stores in Name the name of a move in a given language, and Language
stores the name, and url of the language being used

```go
type ptype struct {
        Ptype Elem `json:"type"`
}
```

ptype stores a types name as well as its url
