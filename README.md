# :book: :palm_tree: bol

> I wanted a notebook that functioned not as a body but as a mind, a notebook that collected, interposed, collaged: a machine whose components could move, whose cogs, chutes, and levers were air. - [Patricia Lockwood](http://www.newyorker.com/magazine/2016/11/28/finding-poetry-in-a-note-taking-app)

*bol* is a client program for editing and synchronization of encrypted documents. The main utility is a command-line program (`bol`) that lets you write/view encrypted documents using your favorite command-line editor. Synchronization is optional and provided through a server program (`bolserver`), where updates are pushed/pulled. A public server is available at https://bol.schollz.com. Both utilities are available in [the latest release](https://github.com/schollz/bol/releases/latest) as a self-contained executable binary for most popular OSes and there are no requirements, except a text-editor (however, `vim` is bundled for the Windows users).

# Install

[Download from releases](https://github.com/schollz/bol/releases/latest).

Or, if you have Go1.7+, just install via `go get`:

```
go get -u -v github.com/schollz/bol/...
```

# Usage

Just run

```
bol
```


**Delete an entry** by editing the entry and replacing the entire entry with ```ignore entry```. **Delete a document** by making a new entry in that document that says ```ignore document```. However, nothing is really deleted, as *bol* will save a copy of every entry and change committed to it.

**Edit a specific document/entry** using the command `bol DocumentName/EntryName`.

**Backup everything** to an AES-encrypted JSON file using `bol -dump`. The dump file `user-20ZZ-YY-XX.bol` can be read using the boltool, `boltool -decrypt user-20ZZ-YY-XX.bol`.

**Erase all local files** using `bol -clean`. This will not remove remote files, there is no way to remove remote files.

**Change user** or the server, by using `bol -config`.

**Summarize** a document using `bol -summary`.

# Server

The default server is a public server, https://bol.schollz.com. You can also run your own server simply running `bolserver`. Then, use `bol -config` and type in the server address, now `http://localhost:9095` or whichever you set it to.

The server is optional, and only required if you want synchronized entires on multiple computers. *bol* will function locally with or without Internet. The server is useful if you wish to edit your document on multiple computers, and also will allow you to add new entries via the website.

# Implementation

*bol* uses `ssed` as a backend for the encrypted storage and synchronization. For more information, [see the `ssed` README](https://github.com/schollz/bol/blob/master/ssed/README.md). Essentially, *bol* stores all entries as AES-encrypted JSON files which are then bundled as a `tar.bz2` archive which is synchronized with a server. Nothing can ever be deleted (only ignored), and all changes are stored.

# Inspiration

Here is some software which are similar to *bol*, but often require other software or system-specific utilities. I enjoy these software, and used a lot of inspiration from them, but ultimately I found that *bol* could provide some functionality or utility that was still absent.

-   [**13 lines of shell**](https://gist.github.com/schollz/27b4ffe562b0b74bf8ee1e8055680d22)- git-based journal  
    No encryption, no editing past entries, no version control, no deletion - but only 13 lines!

-   [**Cryptpad**](https://beta.cryptpad.fr/pad/)  - zero knowledge realtime encrypted editor in your web browser  
    Requires internet access, and a browser, difficult to reconstruct many documents.

-   [**jrnl.sh**](http://jrnl.sh/)- command line journal application with encryption  
    Requires Python. Syncing only available via Dropbox which won’t support merging encrypted files if editing offline.

-   [**ntbk**](https://www.npmjs.com/package/ntbk) - command line journal  
    Requires Node.js. Syncing only available via Dropbox which won’t support merging encrypted files if editing offline.

-   [**vimwiki**](http://vimwiki.github.io/) - command line editor with [capability of distributed encryption](http://www.stochasticgeometry.ie/2012/11/23/vimwiki/)  
    Requires system-specific filesystem encryption (e.g. [eCryptFS](http://www.stochasticgeometry.ie/2012/11/23/vimwiki/)). Merges are not handled.

-   [**Org mode**](http://orgmode.org/) - fast and effective plain-text system for authoring documents  
    Requires `git`.

-   [**gojot**](http://gojot.schollz.com/) - git-based system for distribution editing of encryption of documents (no longer supported)  
    Requires `git`.


# Contributing

I'm open to contributions, submit a pull request.

# License

MIT
