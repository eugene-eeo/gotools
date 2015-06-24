#!/usr/bin/env python
from contextlib import contextmanager
from shutil import copyfile, rmtree
import os
import os.path as path


@contextmanager
def cleanup(pth='build'):
    os.mkdir(pth)
    try:
        yield pth
    finally:
        rmtree(pth)


@contextmanager
def cd(pth):
    prev = os.getcwd()
    os.chdir(pth)
    try:
        yield
    finally:
        os.chdir(prev)


def build(name):
    with cleanup() as pth:
        copyfile(name, os.path.join(pth, name))
        with cd(pth):
            os.system('go build %s' % name)
        binname = name.rstrip('.go')
        src = path.join(pth, binname)
        dst = path.join(path.expanduser('~/.scripts'), binname)
        link(src, dst)


def link(src, dst):
    if path.exists(dst):
        os.unlink(dst)
    os.rename(
        path.abspath(src),
        path.abspath(dst),
    )


if __name__ == '__main__':
    for item in ['sctn.go', 'ark.go']:
        build(item)
