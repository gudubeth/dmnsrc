Domain Search Tool
===============================

[![CircleCI](https://circleci.com/gh/ozgio/dmn.svg?style=svg)](https://circleci.com/gh/ozgio/dmn)

Command line tool for checking domain name availability. 

Usage:

```bash
# check availability of a domain
dmn check example.com

# check multiple domains
dmn check example.com example.org

# check domains from a file
echo "example.com\nexample.org" > names.txt
cat names.txt | dmn check 

# get help
dmn --help
```

-----------

### Prerequisites:

- dep: https://golang.github.io/dep/
- make: https://www.gnu.org/software/make/

### Setup

    # Get dependencies
    make dep

For online code documentation see [Godocs](https://godoc.org/github.com/ozgio/dmn).

TODO
----
* Improve domain check mechanism and add different checking methods based on a flag
* Add name generation support


