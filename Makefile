# must run "go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo" first
app = clickss
bin = Clickss.exe

build:
	del *.syso *.exe *.zip *.log
	go generate
	go build -ldflags "-H=windowsgui" -o $(bin)

run:
	go run .

7zzip: 
	7z a -tzip "$(app).zip" $(bin)

release:
	build zip