# handlers

## Table of Contents

- [Package](#package)
- [Functions](#functions)
  - Exported
    - [Advantage](#advantage)
    - [CommonMoves](#commonmoves)
  - Unexported
    - [damageTo](#damageto)
    - [damageFrom](#damagefrom)
    - [getMovesInCommon](#getmovesincommon)
    - [moveInPokemon](#moveinpokemon)
    - [swap](#swap)
    - [translate](#translate)
    - [typeInDamage](#typeindamage)
- [Types](#types)
  - Unexported
    - [hasAdvantage](#hasadvantage)
    - [hasCommon](#hascommon)

## Package

```go
package handlers // import "github.com/ferpart/pokerest/handlers"
```

Package handlers has the methods used for handling the calls done to the
api.

## FUNCTIONS

### [Advantage](https://github.com/ferpart/pokerest/blob/aea2044ee852502621b9fb25d61e0286982f2f63/handlers/advantage.go#L35)

```go
func Advantage(w http.ResponseWriter, r *http.Request)
```

Advantage given two pokemon we return a json if the first has advantage over
the second given two criteria using the pokemon's type and assuming the
pokemon attacking is using a move of the same type as his:

- Can the first deal double damage to the second?
- Does the first recieve half (or no) damage from the second?
  The DamageTo the second pokemon and the DamageFrom the second pokemon are
  substracted, if we get a number above 0, it means that the 1st pokemon has
  an advantage, if its bellow 0 means that the second pokemon has an advantage
  over the first one, finally if we get 0 we conclude that both pokemon are on
  equal ground.

### [CommonMoves](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/handlers/common.go#L27)

```go
func CommonMoves(w http.ResponseWriter, r *http.Request)
```

CommonMoves returns as a json format an array of all the moves that the
given pokemon have in common. Depending on the flags given, different
pokemon can be compared, the language of the moves can be set, and there's
also a limit of how many moves will be sent that defaults to 10 if no custom
limit is given.

### [damageTo](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/handlers/advantage.go#L77)

```go
func damageTo(damager []helpers.Drelations, damagee helpers.Pokemon) float64
```

damageTo recieves an array of damage relations as the damager, and recieves
a pokemon as a damager. Depending if the type of the pokemon is found within
the damage ralations, the damage value will be multiplied to see how much
damage will be done from the damager to the damagee. This takes into
consideretion the case in which a pokemon would have two types.

### [damageFrom](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/handlers/advantage.go#L103)

```go
func damageFrom(damager helpers.Pokemon, damagee []helpers.Drelations) float64
```

damageFrom recieves a pokemon as a damager and an array of damage relations
as the damagee. Depending if the type of the pokemon is found on the damage
relations, the damge value will be multiplied to see how much damage the
damagee will take from the damaging pokemon. This takes into consideration
the case in which a pokemon could have two types

### [getMovesInCommon](https://github.com/ferpart/pokerest/blob/aea2044ee852502621b9fb25d61e0286982f2f63/handlers/common.go#L46)

```go
func getMovesInCommon(pokeArr []helpers.Pokemon, lan string) []helpers.Elem
```

getMovesInCommon will return an array of moves that appear in all of the
given pokemon movesets. if a language is specified, they'll be translated to
that language, else they'll be left in english.

### [translate](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/handlers/common.go#L96)

```go
func translate(lan string, move helpers.Elem) helpers.Elem
```

translate takes a move, and a language paramater, with that it determines to
which language the move should be translated to, and returns the move in the
new language.

### [swap](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/handlers/common.go#L114)

```go
func swap(i1 int, i2 int, p *[]helpers.Pokemon)
```

swap swaps two elements of a list.

### [typeInDamage](https://github.com/ferpart/pokerest/blob/aea2044ee852502621b9fb25d61e0286982f2f63/handlers/advantage.go#L125)

```go
func typeInDamage(a string, list []helpers.Elem) bool
```

### [moveInPokemon](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/handlers/common.go#L104)

```go
func moveInPokemon(a string, list []helpers.Move) bool
```

moveInPokemon returns true if a string given to a is found on the given
list.

typeInDamage returns true if a string given to a is found on the given list

## TYPES

### [hasAdvantage](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/handlers/advantage.go#L18)

```go
type hasAdvantage struct {
        Name        string  `json:"name"`
        Advantage   bool    `json:"advantage"`
        DamageTo    float64 `json:"damageTo"`
        DamageFrom  float64 `json:"damageFrom"`
        Total       float64 `json:"total"`
        Description string  `json:"description"`
}
```

hasAdvantage is a struct that stores the result of if the first pokemon has
advantage over the second. It has the following fields:

- Advantage: true if first pokemon has advantage over the second.
- DamageTo: is the damage multiplier that the first pokemon can do to the second.
- DamageFrom: is the damage multiplier that the first pokemon can recieve from the second.
- Total: is the difference between the DamageTo and DamageFrom.
- Description: is the description of the shown results.

### [hasCommon](https://github.com/ferpart/pokerest/blob/889085bd71d8f519783f1fa9f8e28287cf99fac8/handlers/common.go#L16)

```go
type hasCommon struct {
	Name string `json:"name"`
	Language string `json:"language"`
	Pokemon []string `json:"pokemon"`
	Moves []helpers.Elem `json:"moves"`
}
```

hasCommon is a struct that stores all the common moves in a struct a list of common moves, as well as the following attributes:

- Name: name of the request type
- Language: language in which the common moves were requested
- Pokemon: pokemon whose common moves were requested
- Moves: moves in common for pokemon
