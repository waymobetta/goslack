## goslack

Sometimes tasks take a long time to complete.

Often I like to move on to other things while long tasks take time to complete.

I'd also like to be notified when said long tasks actually do complete so that I may be able to resume work.

### Download
```
go get github.com/waymobetta/goslack
```


### Slack API key
Create new bot user and obtain API key
```
https://{team}.slack.com/apps/manage/custom-integrations
```

### Usage

build
```bash
$ go build -o goslack example/main.go
$ goslack <message>
```
### Example
Notify me when nmap scan completes
```bash
$ nmap -sT XX.XX.XX.XX-XX;goslack 'nmap scan complete'
```

More basic sleep example (\*nix)
```bash
sleep 5;goslack finished
```

MIT
