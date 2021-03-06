## WKB CLI

Command line tool for dumping WKB metadata. Interface for [nathancahill/wkb](https://github.com/nathancahill/wkb).

## Installation

```
$ go install github.com/nathancahill/wkb-cli
```

## Usage

```
$ wkb-cli rast.bin
```

With PostGIS, dump WKB from PostGIS at any point to a hex file:

```
> COPY (SELECT ST_AsBinary(rast) FROM table LIMIT 1) to 'rast.hex';
```

Then convert the hex file to binary:

```
$ xxd -r -p rast.hex > rast.bin
```

The WKB can then be inspected:

```
$ wkb-cli rast.bin
```
