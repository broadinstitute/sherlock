# `/internal`

This directory holds the majority of the code for a golang application. Packages containing application logic should live in this folder.

The go compiler enforces that packages in `/internal` can not be imported from other go modules via `go get`.

When writing library code that is intended to be importable by other go modules, make sure not to place any public facing packages in this directory.
There is not a definitive standard for organizining public facing packages. One convention is to place them in their directories in the repo root.
Another is to place each publically importable package in it's own directory under `/pkg`. Reccomend using the later pattern in order to reduce clutter
in the repo root.
