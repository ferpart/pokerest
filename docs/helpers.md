# helpers

## Table of Contents

- [Package](#package)
- [Functions](#functions)
  - Exported
    - [GetDamageRelations](#getdamagerelations)
    - [GetLanguage](#getlanguage)
    - [GetMove](#getmove)
    - [GetPokemon](#getpokemon)
    - [SendJSON](#sendjson)
  - Unexported
    - [writeHeader](#writeheader)
- [Types](#types)
  - Exported
    - [CompleteMove](#completemove)
    - [Drelations](#drelations)
    - [Elem](#elem)
    - [Move](#move)
    - [Pokemon](#pokemon)
  - Unexported
    - [drelation](#drelation)
    - [language](#language)
    - [ptype](#ptype)

## Package

```go
package helpers // import "github.com/ferpart/pokerest/helpers"
```

Package helpers stores all the helper functions that are used for parsing
requests to the api.

## FUNCTIONS

### [GetDamageRelations](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/relationparser.go#L27)

```go
func GetDamageRelations(url string) Drelations
```

GetDamageRelations gets da damage relations from the passed url and returns
it in a Drelations struct

### [GetLanguage](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/languageparser.go#L12)

```go
func GetLanguage(res interface{}) string
```

GetLanguage verifies if language exists and returns the name of the
language. In case the language does not exist, it will return an empty
string.

### [GetMove](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/languageparser.go#L44)

```go
func GetMove(url string, lan string) Elem
```

GetMove gets a move with the provided url and translates it to the language
given in lan

### [GetPokemon](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/pokeparser.go#L34)

```go
func GetPokemon(pokeNames []string) []Pokemon
```

GetPokemon given an array of two pokemon, will return an array of type
Pokemon

### [SendJSON](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/sendjson.go#L11)

```go
func SendJSON(w http.ResponseWriter, status int, body interface{})
```

SendJSON recieves a ResponseWriter, a status, and a body then sends a json
to the ResponseWriter

### [writeHeader](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/sendjson.go#L22)

```go
func writeHeader(w *http.ResponseWriter, status int)
```

writeHeader writes the headers required for setting the application to type json and allowing cors

## TYPES

### [CompleteMove](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/languageparser.go#L31)

```go
type CompleteMove struct {
        Name  string     `json:"name"`
        Names []language `json:"names"`
}
```

CompleteMove saves the name, and all the names in different locales of a
move

### [Drelations](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/relationparser.go#L10)

```go
type Drelations struct {
        Name      string    `json:"name"`
        Drelation drelation `json:"damage_relations"`
}
```

Drelations stores a type, and the damage relations of said type

### [Elem](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/pokeparser.go#L28)

```go
type Elem struct {
        Name string `json:"name"`
        URL  string `json:"url"`
}
```

Elem stores a name as well as a url

### [Move](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/pokeparser.go#L23)

```go
type Move struct {
        Move Elem `json:"move"`
}
```

Move stores an Elem containing a move's data

### [Pokemon](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/pokeparser.go#L11)

```go
type Pokemon struct {
        Name   string  `json:"name"`
        Ptypes []ptype `json:"types"`
        Moves  []Move  `json:"moves"`
}
```

Pokemon stores a pokemon's name, type, and moveset

### [drelation](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/relationparser.go#L16)

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

### [language](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/languageparser.go#L38)

```go
type language struct {
        Name     string `json:"name"`
        Language Elem   `json:"language"`
}
```

language stores in Name the name of a move in a given language, and Language
stores the name, and url of the language being used

### [ptype](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/helpers/pokeparser.go#L18)

```go
type ptype struct {
        Ptype Elem `json:"type"`
}
```

ptype stores a types name as well as its url
