# gotools

Command line utilities written in Go. Basically an excuse
for me to force myself to learn the more practical aspects
(IO, etc) of Go instead of being stuck in ranting about
the not-as-powerful-as-Rust's type system. Also made because
I am lazy to read Unix documentation and I do not know how
to use the string manipulation commands:

```shell
# Before I uninstall xz, are there any
# packages that depend on it? (Intersect
# the outputs)
$ sctn "$(brew list)" "$(brew uses xz)"

# ls the directory with multiple regexes
# that all must match
$ ark 'a(.+?)' '(.+?)[1-9]{2,4}'
```

**Note [1]:** After I've built `ark` I realized I had just
reimplemented a simpler version of `grep` in Go. Oh
well. Not going to delete it.

**Note [2]:** The ``build.py`` script copies the built
binaries to ``~/.scripts``. If you do not have that in
your path, the script will explode but the intermediate
build files will always be deleted.
