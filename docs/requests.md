# Requests

The api specification can be seen in both `json` and `yaml` format in the following urls:

- [json](https://github.com/ferpart/pokerest/blob/master/Docs/pokerest.json)
- [yaml](https://github.com/ferpart/pokerest/blob/master/docs/pokerest.yaml)

## Endpoints

In this project we currently have only two endpointsand these are:

- [/api/advantage](#advantage)
  - [Examples](#examples)
- [/api/commonMoves](#common-moves)
  - [Examples](#examples-1)

We are asumming you'll be mounting the project locally so we'll use:

```text
http://localhost:5000
```

For the requests

## Advantage

```text
http://localhost:5000/api/advantage
```

This api recieves a param with name: `pokemon`, and two arguments separated by a comma. Like the following text:

```text
http://localhost:5000/api/advantage?pokemon=pikachu,bulbasaur
```

### Examples

`GET` http://localhost:5000/api/advantage?pokemon=squirtle,rhydon

```json
{
  "name": "Advantage",
  "advantage": true,
  "damageTo": 4,
  "damageFrom": 1,
  "total": 3,
  "description": "squirtle has advantage over rhydon"
}
```

`GET` http://localhost:5000/api/advantage

```json
{
  "name": "Advantage",
  "advantage": false,
  "damageTo": 0,
  "damageFrom": 0,
  "total": 0,
  "description": "No pokemon were given"
}
```

## Common Moves

```text
http://localhost:5000/api/commonMoves
```

This api recieves 3 params with the following values:

- `pokemon`: these can be two or more pokemon separated by a comma
- `language`: this is the language in which the common moves will be displayed if none is given, or the language is not found, will default to english.
- `limit`: this is a number which will limit the ammount of moves shown if given 0 or nothing, will default to 10 elements.

```text
http://localhost:5000/api/commonMoves?pokemon=charizard,squirtle,lucario,lugia,rayquaza,onix&language=es&limit=5
```

### Examples

`GET` http://localhost:5000/api/commonMoves?pokemon=charizard,squirtle,lucario,lugia,rayquaza,onix&language=es&limit=5

```json
{
  "name": "Common",
  "language": "es",
  "pokemon": ["charizard", "squirtle", "lucario", "lugia", "rayquaza", "onix"],
  "moves": [
    {
      "name": "Golpe Cabeza",
      "url": "https://pokeapi.co/api/v2/move/29/"
    },
    {
      "name": "Fuerza",
      "url": "https://pokeapi.co/api/v2/move/70/"
    },
    {
      "name": "Tóxico",
      "url": "https://pokeapi.co/api/v2/move/92/"
    },
    {
      "name": "Doble Equipo",
      "url": "https://pokeapi.co/api/v2/move/104/"
    },
    {
      "name": "Descanso",
      "url": "https://pokeapi.co/api/v2/move/156/"
    }
  ]
}
```

`GET` http://localhost:5000/api/commonMoves?pokemon=charizard,squirtle&language=ja&limit=5

```json
{
  "name": "Common",
  "language": "ja",
  "pokemon": ["charizard", "squirtle"],
  "moves": [
    {
      "name": "メガトンパンチ",
      "url": "https://pokeapi.co/api/v2/move/5/"
    },
    {
      "name": "メガトンキック",
      "url": "https://pokeapi.co/api/v2/move/25/"
    },
    {
      "name": "ずつき",
      "url": "https://pokeapi.co/api/v2/move/29/"
    },
    {
      "name": "のしかかり",
      "url": "https://pokeapi.co/api/v2/move/34/"
    },
    {
      "name": "とっしん",
      "url": "https://pokeapi.co/api/v2/move/36/"
    }
  ]
}
```

`GET` http://localhost:5000/api/commonMoves?pokemon=charizard,squirtle&limit=5

```json
{
  "name": "Common",
  "language": "en",
  "pokemon": ["charizard", "squirtle"],
  "moves": [
    {
      "name": "mega-punch",
      "url": "https://pokeapi.co/api/v2/move/5/"
    },
    {
      "name": "mega-kick",
      "url": "https://pokeapi.co/api/v2/move/25/"
    },
    {
      "name": "headbutt",
      "url": "https://pokeapi.co/api/v2/move/29/"
    },
    {
      "name": "body-slam",
      "url": "https://pokeapi.co/api/v2/move/34/"
    },
    {
      "name": "take-down",
      "url": "https://pokeapi.co/api/v2/move/36/"
    }
  ]
}
```
