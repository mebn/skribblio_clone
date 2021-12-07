# DD1349_indaprojekt_skribblio
En kopia av [skribble.io](https://skribbl.io)!

## Vad är det här?
Det här är en kopia av det populära ritspelet [skribble.io](https://skribbl.io). Spelarna roterar om vem som ska få rita. Ritaren får då välja mellan 3 ord och sedan rita så bra man kan. De andra spelarna ska sedan gissa ordet och om de gissar rätt får de poäng!

## Hur fungerar det?
Vi använder [Gorilla](https://github.com/gorilla/websocket) som är ett paket till Go som tar hand om websockets. Detta använder vi för att kommunicera med alla spelare i olika rum. På frontend använder vi JS inbyggda websockets som då enkelt kan prata med vår server.

## Hur man kör
### Installera alla nödvändiga paket
Gorilla: `go get github.com/gorilla/websocket`.

English: `go get github.com/gregoryv/english`.

### Starta servern
Efter det är det bara att köra `go run .` och gå in på [localhost:8080](http://localhost:8080/).

