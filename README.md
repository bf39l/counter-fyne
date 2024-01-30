# Counter - Fyne
A simple counter application created with Fyne + GoLang

## ToDo
- [x] Basic Increase and Decrease buttons
- [ ] Settings dialog or window for setting up 
    - [x] counting step(s) - default is 1
    - [ ] enable or disable duration calculation
- [ ] Get duration for each count number change
- [ ] Add keyboard shortcut for updating counter therefore
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
## General problem
- macOS block application to launch with err: "Counter - Fyne.app" is damaged and can’t be opened. You should move it to the Bin
    ```bash
    # Could add attribute like follow
    xattr -rc ~/path/to/Counter\ -\ Fyne.app
    ```