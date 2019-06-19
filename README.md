# Shibafu Programming Language
[![CircleCI](https://circleci.com/gh/Code-Hex/shibafu.svg?style=svg&circle-token=42af5120e1edd375241967a09e303d2a4064b840)](https://circleci.com/gh/Code-Hex/shibafu)

<p align="center">
  <img width="460" src="https://user-images.githubusercontent.com/6500104/59754663-7fae8d00-92c1-11e9-9b11-a9a3ec172967.jpg">
</p>

Online in Japan, people have been using "w" is from "warau" (笑う) or "warai" (笑い), the Japanese word for laugh or smile. they call "w" "kusa" (草) which means grass. 
If you want to know more details of this culture. See "[In Japan, People Do Not LOL. They wwww.](https://kotaku.com/in-japan-people-do-not-lol-they-wwww-5986170)"

When this grass (草) gathers it is called lawn (芝生, Shibafu).

This language is named `shibafu` because this programming language will be growing many grasses.

The syntax of this language is based on [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck). you can use these tokens.

|  Token  |                                                                                                                                                                            Meaning                                                                                                                                                                              |
|---------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|  `www`  | increment the data pointer (to point to the next cell to the right).                                                                                                                                                                                                                                                                                              |
|  `WWW`  | decrement the data pointer (to point to the next cell to the left).                                                                                                                                                                                                                                                                                               |
|  `wWw`  | increment (increase by one) the byte at the data pointer.                                                                                                                                                                                                                                                                                                         |
|  `WwW`  | decrement (decrease by one) the byte at the data pointer.                                                                                                                                                                                                                                                                                                         |
|  `wwW`  | output the byte at the data pointer.                                                                                                                                                                                                                                                                                                                              |
|  `Www`  | accept one byte of input, storing its value in the byte at the data pointer.                                                                                                                                                                                                                                                                                      |
|  `wWW`  | if the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next token, jump it forward to the token after the matching `WWw` token.                                                                                                                                                                                   |
|  `WWw`  | if the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching `wWW` command.                                                                                                                                                                                 |

You can see [examples](https://github.com/Code-Hex/shibafu/tree/master/example).
