# handlers

## Table of Contents

- [Package](#package)
- [Functions](#functions)
  - Exported
    - [Advantage](#advantage)
    - [CommonMoves](#commonmoves)
  - Unexported
    - [damageFrom](#damagefrom)
    - [damageTo](#damageto)
    - [getMovesInCommon](#getmovesincommon)
    - [moveInPokemon](#moveinpokemon)
    - [swap](#swap)
    - [translate](#translate)
    - [typeInDamage](#typeindamage)
- [Types](#types)
  - Unexported
    - [hasAdvantage](#hasadvantage)

## Package

```go
package handlers // import "github.com/ferpart/pokerest/handlers"
```

Package handlers has the methods used for handling the calls done to the
api.

## FUNCTIONS

### Advantage

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

### CommonMoves

```go
func CommonMoves(w http.ResponseWriter, r *http.Request)
```

CommonMoves returns as a json format an array of all the moves that the
given pokemon have in common. Depending on the flags given, different
pokemon can be compared, the language of the moves can be set, and there's
also a limit of how many moves will be sent that defaults to 10 if no custom
limit is given.

### damageFrom

```go
func damageFrom(damager helpers.Pokemon, damagee []helpers.Drelations) float64
```

damageFrom recieves a pokemon as a damager and an array of damage relations
as the damagee. Depending if the type of the pokemon is found on the damage
relations, the damge value will be multiplied to see how much damage the
damagee will take from the damaging pokemon. This takes into consideration
the case in which a pokemon could have two types

### damageTo

```go
func damageTo(damager []helpers.Drelations, damagee helpers.Pokemon) float64
```

damageTo recieves an array of damage relations as the damager, and recieves
a pokemon as a damager. Depending if the type of the pokemon is found within
the damage ralations, the damage value will be multiplied to see how much
damage will be done from the damager to the damagee. This takes into
consideretion the case in which a pokemon would have two types.

### getMovesInCommon

```go
func getMovesInCommon(pokeArr []helpers.Pokemon, lan string) []helpers.Elem
```

getMovesInCommon will return an array of moves that appear in all of the
given pokemon movesets. if a language is specified, they'll be translated to
that language, else they'll be left in english.

### moveInPokemon

```go
func moveInPokemon(a string, list []helpers.Move) bool
```

moveInPokemon returns true if a string given to a is found on the given
list.

### swap

```go
func swap(i1 int, i2 int, p *[]helpers.Pokemon)
```

swap swaps two elements of a list.

### translate

```go
func translate(lan string, move helpers.Elem) helpers.Elem
```

translate takes a move, and a language paramater, with that it determines to
which language the move should be translated to, and returns the move in the
new language.

### typeInDamage

```go
func typeInDamage(a string, list []helpers.Elem) bool
```

typeInDamage returns true if a string given to a is found on the given list

## TYPES

### hasAdvantage

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
