# bol

> I wanted a notebook that functioned not as a body but as a mind, a notebook that collected, interposed, collaged: a machine whose components could move, whose cogs, chutes, and levers were air. - [Patricia Lockwood](http://www.newyorker.com/magazine/2016/11/28/finding-poetry-in-a-note-taking-app)

*bol* is [a command-line program](https://github.com/schollz/bol/releases) that lets you write/view encrypted documents and [a webpage](https://bol.schollz.com/) that lets you write (not view) documents.

*bol* uses `ssed` as a backend for the encrypted storage and synchronization. For more information, [see the white paper](https://github.com/schollz/bol/blob/master/WHITEPAPER.md).

## Install

```
go get -u -v github.com/schollz/bol/...
```

## Run

```
bol
```

To delete entry, just delete the entire entry and replace with ```ignore entry```.

To delete a document, just make a new entry that says ```ignore document```.

The files `bol` creates can be inspected with the `boltool`,

```
boltool -decrypt e53a4a99301c71e6039cb80d52db09a6083bb0913df83fe2343db36d9edf4aae.json
```
