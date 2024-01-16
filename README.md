# TreeSize

Print folders in tree format and display file's size.

## Install

```bash
go install github.com/GoToUse/TreeSize@latest
```

## Examples

```bash
╰─± TreeSize --help
Usage of TreeSize:
  -e value
    	Exclude directories.
  -f string
    	Folder path. (default ".")
  -h	Print the size in a more human readable way.

### Usage Case
╰─± TreeSize -f /Users/dapeng/Desktop/code/python3/learn -h
/Users/dapeng/Desktop/code/python3/learn
.
├── zero.py (704 B)
├── __pycache__
│   └── zero.cpython-310.pyc (696 B)
├── demo.py (512 B)
├── demo1.py (1.5KiB)
├── celery_learn
├── brackets.py (551 B)
├── hw
│   ├── div2.py (248 B)
│   ├── hex.py (596 B)
│   ├── sort_str.py (437 B)
│   ├── hj35.py (819 B)
│   ├── hj31.py (265 B)
│   ├── hj22.py (692 B)
│   ├── number.py (346 B)
│   └── hj54.py (3.7KiB)
├── goCode
│   ├── main.go (468 B)
│   ├── sortStr.go (732 B)
│   ├── go.mod (23 B)
│   ├── hj34.go (345 B)
│   └── reverseStr.go (573 B)
├── spawn_fork.py (457 B)
├── generic_type.py (993 B)
├── zaro1.py (862 B)
├── cv
│   └── demo.py (812 B)
├── logg
│   ├── test.py (96 B)
│   ├── main.py (262 B)
│   ├── t.log (100 B)
│   ├── __init__.py (0 B)
│   └── log.py (782 B)
├── iterator_folder
│   ├── compress_i.py (558 B)
│   ├── _base.py (380 B)
│   ├── __init__.py (71 B)
│   ├── product_i.py (1010 B)
│   ├── accumulate_i.py (557 B)
│   ├── permutation_i.py (788 B)
│   ├── __pycache__
│   │   ├── __init__.cpython-37.pyc (221 B)
│   │   └── _base.cpython-37.pyc (640 B)
│   ├── qtlogfh.log (3.4KiB)
│   └── test.py (1.8KiB)
├── descriptor
│   ├── d1.py (1.1KiB)
│   ├── d2.py (1.2KiB)
│   └── d3.py (872 B)
├── main.py (1.9KiB)
└── socket
    ├── client.py (612 B)
    ├── server.py (715 B)
    ├── sf
    │   ├── server.py (2.4KiB)
    │   ├── client.py (1.8KiB)
    │   └── client1.py (1.8KiB)
    └── socket.io.folder

Summary: Total folders: 12 Total files: 47 Total size: 39.2KiB
```
