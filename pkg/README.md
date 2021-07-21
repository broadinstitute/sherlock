# `/pkg`

This directory contains packages that are public and can be imported by other go projects. It is usually used for shared library
code intended to be importable from other go modules.

Usage of this directory is optional depending on needs of the project namely doesn't it include importable packages relevant outside of a particular project. There are varying standards on how to organize `go get`able packages, however go projects generally tend to favor shallow and wide directory layouts.
