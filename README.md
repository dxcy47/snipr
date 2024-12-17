# snipr

snipr (snipper) is a simple screenshoting tool for linux (only works on x11, and requires libx11 to work).
After running, click on your screen wherever you want the first corner of your screenshot to be, and then click wherever you want the second corner to be, and the result will be copied to your clipboard.

# TODO List
- [ ] add config functionality
- [ ] possibly add some sort of visual to show which area is being copied

# Installation
**You need go on your system to install snipr.**

It is possible to either run it as is, by running `go run snipr.go` while in the files directory, running `go build snipr.go` yourself to compile an executable, or run install.sh to autoamtically compile it, move it to PATH, and (whenever i get around to implementing it) create a config directory in ~/.config
