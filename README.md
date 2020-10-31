# File Tree
This is a small utility to make a tree of a directory. App stores tree in bolt database. It also possible to find duplicates by exactly match of the size.

## Usage
 ```
Usage:
  --help [OPTIONS] <collect | dups | print>

Help Options:
  -h, --help  Show this help message

Available commands:
  collect  gather data to database
  dups     find duplicates in database basing on file size
  print    print content of database
```

## Example
```
$ mkdir -p {test, test/a, test/b} && echo 123 > test/123 && echo 123 > test/a/123 && echo 2222 > test/2222                                                                               Sat Oct 31 20:32:45 2020
$ ./ftree collect -d ex.db test/                                                                                                                                                         Sat Oct 31 20:32:46 2020
$ ./ftree print -d ex.db                                                                                                                                                                 Sat Oct 31 20:32:51 2020
test/ 192 [directory]
test/123 4 [file]
test/2222 5 [file]
test/a 96 [directory]
test/a/123 4 [file]
test/b 64 [directory]
$ /ftree dups -d ex.db                                                                                                                                                                  Sat Oct 31 20:32:55 2020
test/a/123
test/123
```