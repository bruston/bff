bff
===

Yet another brainfuck interpreter written in Go. Just for fun.

## Usage

Supply bff with the path to a file containing a brainfuck program. 

```
bff filename.bf
```

The starting number of memory cells defaults to 1048576 (1024 * 1024) and is adjusted with the -cells flag.

```
bff -cells=1024 filename.bf
```
