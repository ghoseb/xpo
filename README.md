```
__  __ _ __    ___
\ \/ /| '_ \  / _ \
 >  < | |_) || (_) |
/_/\_\| .__/  \___/
      | |
      |_|
Highlight strings by regular expressions in the terminal
```

# What is xpo?

Remember staring at large log files and trying to find all IP addresses in a wall of text? It's not only difficult
to parse such files visually, it's also difficult to find if they are all the same or if they are changing.

`xpo` solves this problem by colorizing regexp matches while making sure same matches always have the same color
and similar matches have similar colors. If `127.0.0.1` is green-ish, `127.0.0.2` ought to be close enough.

Colors are assigned kind-of arbitrarily, but it's also possible to assign explicit colors to matched strings.

Color space: `6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)`

# Building

```bash
# you'll need the Go toolchain to compile the project
$ brew install go

# cd src/xpo
$ make build

# move the binary from bin to a suitable location
$ mv bin/xpo ~/bin/
```

# Homebrew package

TODO

# Usage

```bash
# highlight any `ERR`, `WARN` or `INFO` in the log file
$ cat logfile | xpo -r "ERR|WARN|INFO" | less -R

# highlight IP addresses and HTTP methods
$ cat logfile | xpo -r "^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}|GET|POST" | less -R

# use specific colors using -e
# specified as "r,g,b:string" where (0 ≤ r, g, b ≤ 5)
# the -e option can be used multiple times
$ cat logfile | xpo -r "^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}|GET|POST" -e "0,5,0:GET" -e "5,0,0:POST" | less -R

# if we are on a light background, -l may be used to exclude light colors
$ cat logfile | xpo -r "ERR|WARN|INFO" -l | less -R
```

# Useful regular expressions

* IPv4 - `((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}`
* UUIDv4 - `\b[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}\b`

# Credits
* Steve Losh for the idea and [wonderful implementation in Common Lisp](https://stevelosh.com/blog/2021/03/small-common-lisp-cli-programs/)
