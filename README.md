# TreeSize

Print folders in tree format and display file's size.

## Install

```bash
go install github.com/GoToUse/TreeSize@latest
```

## Examples

```bash
╰─± go run main.go -f /Users/dapeng/Desktop/code/python3/learn 
Total size: 43.5KiB
Output in tree format:
.
└── learn
    ├── zero.py (704 B)
    ├── brackets.py (551 B)
    ├── celery_learn
    ├── main.py (1.9KiB)
    ├── spawn_fork.py (457 B)
    ├── generic_type.py (993 B)
    ├── demo.py (512 B)
    ├── demo1.py (1.5KiB)
    ├── iterator_folder
    │   ├── test.py (1.8KiB)
    │   ├── __init__.py (71 B)
    │   ├── __pycache__
    │   │   ├── _base.cpython-37.pyc (640 B)
    │   │   └── __init__.cpython-37.pyc (221 B)
    │   ├── product_i.py (1010 B)
    │   ├── permutation_i.py (788 B)
    │   ├── _base.py (380 B)
    │   ├── qtlogfh.log (3.4KiB)
    │   ├── accumulate_i.py (557 B)
    │   └── compress_i.py (558 B)
    ├── hw
    │   ├── sort_str.py (437 B)
    │   ├── div2.py (248 B)
    │   ├── hex.py (596 B)
    │   ├── hj22.py (692 B)
    │   ├── hj31.py (265 B)
    │   ├── hj35.py (819 B)
    │   ├── hj54.py (3.7KiB)
    │   ├── merge_value.py (310 B)
    │   └── number.py (346 B)
    ├── zaro1.py (153 B)
    ├── IndexPy
    │   ├── demo.py (338 B)
    │   ├── .idea
    │   │   ├── workspace.xml (8.8KiB)
    │   │   ├── .gitignore (176 B)
    │   │   ├── IndexPy.iml (327 B)
    │   │   ├── misc.xml (195 B)
    │   │   ├── modules.xml (266 B)
    │   │   ├── inspectionProfiles
    │   │   │   ├── profiles_settings.xml (174 B)
    │   │   │   └── Project_Default.xml (1.2KiB)
    │   │   └── codeStyles
    │   │       └── codeStyleConfig.xml (149 B)
    │   ├── test.json (421 B)
    │   ├── demo1.py (453 B)
    │   └── __pycache__
    │       └── demo1.cpython-39.pyc (530 B)
    ├── socket
    │   ├── client.py (612 B)
    │   ├── server.py (715 B)
    │   └── socket.io.folder
    ├── descriptor
    │   ├── d1.py (1.1KiB)
    │   └── d3.py (872 B)
    ├── goCode
    │   ├── sortStr.go (732 B)
    │   ├── go.mod (23 B)
    │   ├── hj34.go (345 B)
    │   ├── main.go (468 B)
    │   └── reverseStr.go (573 B)
    └── __pycache__
        └── zero.cpython-310.pyc (696 B)
```
