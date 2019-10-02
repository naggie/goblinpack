Goblinpack allows packing of binary assets inside Go executables. Unlike other packers, it:

* Has no binary overhead (some packers use base64 or even ascii-hex
  gzip data. Goblinpack uses byte slices. There is a 6x overhead for source,
  but (practically) no overhead for the compiled executable.
* Allows control of what executables get which assets
* Does not have any clever code to work in dev mode without creating bundles (which makes it simpler)
* Does not compress the files. I suggest using [UPX](https://upx.github.io/) to
  compress the entire binary, with `-s -w` LDFLAGS to strip debugging symbols.
* Produces self-contained data files with no dependencies
* Is actively maintained


# Usage

    goblinpack <name> <files...>

Goblinpack will then generate a golang module:

    _data/<name>/data.go
    _data/<name>/decoders.go

Let's say you have a directory called `sounds` containing `.wav` files.

    go get github.com/naggie/goblinpack
    goblinpack sounds sounds/*

You can then use the data:

    import _data

    r, err := _data.GetReader("test.wav")


Produces a module, `_data/sounds` which contains the following methods:

* `GetReader(filepath string) (io.Reader, err)`
* `GetByteSlice(filepath string) *bytes[]`


-----

If there's demand I might implement `http.FileSystem` (like
[packr](https://github.com/gobuffalo/packr) for easy use with `http.Server`.)

Packr's implementation: https://github.com/gobuffalo/packd/blob/master/file.go


----

* TODO tests
* TODO go fmt on data files should do nothing
