# Counter - Fyne
A simple counter application created with Fyne + GoLang

## ToDo
- [x] Basic Increase and Decrease buttons
- [ ] Settings dialog or window for setting up counting step(s) default is 1
- [ ] Get duration for each count number change
- [ ] More...

## Local run and build
### Local run
```bash
go run main.go
```
### Local build
```bash
# Make sure $GOBIN in $PATH
go install fyne.io/fyne/v2/cmd/fyne@latest

# macOS
fyne package --os darwin
```