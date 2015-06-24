#!/usr/bin/env python
from contextlib import contextmanager
from shutil import copyfile
import os
import os.path as path


@contextmanager
def cleanup(pth):
    try:
        yield pth
    finally:
        if path.exists(pth):
            os.unlink(pth)


@contextmanager
def cd(pth):
    prev = os.getcwd()
    os.chdir(pth)
    try:
        yield
    finally:
        os.chdir(prev)


def build(name):
    src = path.join('%s-bin' % name)
    with cleanup(src):
        os.system('go build -o %s ./%s' % (src, name))
        dst = path.join(path.expanduser('~/.scripts'), name)
        link(src, dst)


def link(src, dst):
    if path.exists(dst):
        os.unlink(dst)
    os.rename(
        path.abspath(src),
        path.abspath(dst),
    )


if __name__ == '__main__':
    for item in ['sctn', 'ark']:
        build(item)
