Goblinpack allows packing of binary assets inside Go executables. Unlike other packers, it:

* Has no binary overhead (some packers use base64 or even ascii-hex
  gzip data. Goblinpack uses byte slices. There is a 5x overhead for source,
  but (practically) no overhead for the compiled executable.
* Allows control of what executables get which assets
* Does not have any clever code to work in dev mode without creating bundles (which makes it simpler)
* Is actively maintained


# Usage

    gobinarypack --name sounds sounds/*


Produces a module, `goblinpack-generated/sounds` which contains the following methods:

* `GetReader(filepath string) (io.Reader, err)`
* `GetByteSlice(filepath string) *bytes[]`


-----

If there's demand I might implement `http.FileSystem` (like
[packr](https://github.com/gobuffalo/packr) for easy use with `http.Server`.
