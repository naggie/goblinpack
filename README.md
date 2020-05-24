Goblinpack allows packing of binary assets inside Go executables. Unlike other packers, it:

* Has no binary overhead (some packers use base64 or even ascii-hex
  gzip data. Goblinpack uses byte slices. There is a 6x overhead for source,
  but (practically) no overhead for the compiled executable.
* Does not have any clever code to work in dev mode without creating bundles (which makes it simpler)
* Does not compress the files. I suggest using [UPX](https://upx.github.io/) to
  compress the entire binary, with `-s -w` LDFLAGS to strip debugging symbols.
* Produces self-contained data files with no dependencies


# Usage

    goblinpack <module path> <files...>

Goblinpack will then generate a golang module, creating directories as necessary:

    <module path>/data.go
    <module path>/decoders.go

Let's say you have a directory called `sounds` containing `.wav` files.

    GO111MODULE=off go get -u github.com/naggie/goblinpack/goblinpack
    goblinpack _data/sounds sounds/*

You can then use the data:

    import github.com/username/repository/_data/sounds

    r, err := _data.GetReader("test.wav")


Produces a module, `_data/sounds` which contains the following methods:

* `GetReader(filepath string) (io.Reader, err)`
* `GetByteSlice(filepath string) *bytes[]`


# Tips

* Add `_data` to your `.gitignore` file. Generated data file should not be checked in
* Make sure you don't accidentally import data files that aren't needed
* Use [UPX](https://upx.github.io/) to compress the entire binary
* Use [`-s -w` LDFLAGS](https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/)
  to strip debugging symbols to make your exe even smaller
* Use [go generate](https://blog.carlmjohnson.net/post/2016-11-27-how-to-use-go-generate/)
  to activate goblinpack as part of your build pipeline


-----

If there's demand I might implement `http.FileSystem` (like
[packr](https://github.com/gobuffalo/packr) for easy use with `http.Server`.)

Packr's implementation: https://github.com/gobuffalo/packd/blob/master/file.go


----

Tests to do:
* go fmt on data files should do nothing
* roundtrip test


# Large files

See https://github.com/golang/go/wiki/GcToolchainTricks . Note appending a zip
archive is more suitable for larger (>10MB) files as the go compiler struggles
with large source files.
