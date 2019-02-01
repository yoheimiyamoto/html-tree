# Overview
Parse html and print tree map

```
.
├── head
│    ├── meta
│    ├── title
│    └── link
└── body
     ├── div (header)
     │    └── div (logo)
     ├── div (main)
     │    ├── h2 (main-title)
     │    └── div (contents)
     │         ├── h3 (sub-title)
     │         ├── div (contents-item)
     │         │    ├── a
     │         │    │    └── p (title)
     │         │    └── p (date)
     │         ├── div (contents-item)
     │         │    ├── a
     │         │    │    └── p (title)
     │         │    └── p (date)
     │         └── div (contents-item)
     │              ├── a
     │              │    └── p (title)
     │              └── p (date)
     └── script
```

# Install
```
go get github.com/yoheimiyamoto/html-tree
```

# Usage
```
go run main.go {URL}
```