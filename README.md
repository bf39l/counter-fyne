# Counter - Fyne
A simple counter application created with Fyne + GoLang

*This is a learn by doing project*

### Motivation
I would like to use a counter for counting how many sets I did in gym training. Of course I could use an existing counter from web but I would also want to track how long I took to finish each set which exsting counter does not fit my requirements. Also I want the application has to be controlled by a "remote" (I may build a bluetooth/wireless two buttons macro pad as remote control) therefore I don't need to touch laptop/device with sweaty hands.

#### Why fyne + golang
I am working for a company which use golang to do backend web services. So keeping practising golang is one of purpose. Regarding GUI I did some researches and found out fyne is a rising staring and really easy to use.

## ToDo
- [x] Basic Increase and Decrease buttons
- [ ] Reset button for counter
- [ ] Settings dialog or window for setting up 
    - [x] counting step(s) - default is 1
    - [ ] enable or disable duration calculation
- [ ] Get duration for each count number change
    - more timing related todos will be created
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
