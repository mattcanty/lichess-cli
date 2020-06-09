# Lichess CLI

## Install

```shell
git clone git@github.com:mattcanty/lichess-cli.git
cd lichess-cli
make install
```

## Configure

1. Create a Lichess account
2. Go to [https://lichess.org/account/oauth/token](https://lichess.org/account/oauth/token)
3. Create a new token with `Play games with the board API` selected
4. Copy the token and place in `~/.config/lichess-cli/config.json` like so:

```json
{
  "lichess_api_key": "xxxxxxxxxxxxxxxx"
}
```

## Play

### View Games

```shell
➜  ~ lichess-cli games
┏━━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━┓
┃ ID       ┃ MY TURN ┃ OPPONENT     ┃ LAST MOVE ┃ BOARD           ┃
┣━━━━━━━━━━╋━━━━━━━━━╋━━━━━━━━━━━━━━╋━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━┫
┃ hP09Ep7h ┃ true    ┃ A.I. level 1 ┃ h2h4      ┃ ♜ ♞ ♝ ♚ ♛ ♝ ♞ ♜ ┃
┃          ┃         ┃              ┃           ┃ - ♟ ♟ ♟ - - ♟ ♟ ┃
┃          ┃         ┃              ┃           ┃ - - - - ♟ ♟ - - ┃
┃          ┃         ┃              ┃           ┃ ♟ - - - - - - - ┃
┃          ┃         ┃              ┃           ┃ - - - ♙ ♙ - - - ┃
┃          ┃         ┃              ┃           ┃ - - - - - - - - ┃
┃          ┃         ┃              ┃           ┃ ♙ ♙ ♙ - - ♙ ♙ ♙ ┃
┃          ┃         ┃              ┃           ┃ ♖ ♘ ♗ ♔ ♕ ♗ ♘ ♖ ┃
┣━━━━━━━━━━╋         ┃              ┣━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━┫
┃ mjqGC18w ┃         ┃              ┃ f5c2      ┃ ♖ - - ♕ ♔ ♗ - ♖ ┃
┃          ┃         ┃              ┃           ┃ ♙ ♙ - ♘ ♙ ♙ ♙ ♙ ┃
┃          ┃         ┃              ┃           ┃ - - ♙ - - ♘ - - ┃
┃          ┃         ┃              ┃           ┃ - - - ♙ - - - - ┃
┃          ┃         ┃              ┃           ┃ - - - ♟ - - - - ┃
┃          ┃         ┃              ┃           ┃ ♟ ♟ ♞ - ♟ ♛ - - ┃
┃          ┃         ┃              ┃           ┃ - - ♗ - - ♟ ♟ ♟ ┃
┃          ┃         ┃              ┃           ┃ ♜ - ♝ - ♚ ♝ ♞ ♜ ┃
┗━━━━━━━━━━┻━━━━━━━━━┻━━━━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━┛
```

### Make a Move

`lichess-cli play m c1c2`

`m` is the prefix of the game ID and `c1c2` is your move.
