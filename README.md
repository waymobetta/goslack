## goslack

Sometimes tasks take a long time to complete.

Often I like to move on to other things while long tasks take time to complete.

I'd also like to be notified when said long tasks actually do complete so that I may be able to resume work.

---

### install
```
go get github.com/waymobetta/goslack
```

### slack API key
_Create new bot user and obtain API key_
```
https://{team}.slack.com/apps/manage/custom-integrations
```

### usage

_Build CLI based on example_
```bash
$ go build -o $GOPATH/bin/goslack example/main.go
$ goslack <message>
```
### examples
_Notify me when nmap scan completes_
```bash
$ nmap -sC -sV XX.XX.XX.XX-XX > scsv.nmap; goslack 'nmap scan complete'
```

_Basic sleep example (\*nix)_
```bash
sleep 3; goslack finished
```

MIT
