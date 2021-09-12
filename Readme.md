# S2k (Send to Kindle)
A little cli script to send files to your Amazon Kindle(TM) via internet.


## How to use the tool:
* After Downloading the go file you can run:

- Add env variables:

```sh

export EMAIL_USERNAME=yyy       # Your smtp username
export EMAIL_PASSWORD=xxx       # Your smtp password
export EMAIL_PORT=xxx           # Your smtp port default is: smtp.gmail.com
export EMAIL_HOST=xxx           # Your smtp port default is: 465

```
- Run (after cloning this project) the script on your favorite shell:

```sh
$ git clone git@github.com:AbderrahimSoubaiElidrissi/s2k.git

$ cd s2k
```

1. You can run it directly with golan in your machine:

`go run ./main.go <URL> <KINDLE_EMAIL>` with URL is the url to the file and `KINDLE_EMAIL` as your send-to-Kindle e-mail address.

So you run something like: `go run ./main.go https://www.w3.org/Press/98Folio.pdf soubai_fgsufgsu@kindle.com`

2. You can install it 

```sh
$ go build
$ go install
$ $GOPATH/bin/s2k https://www.w3.org/Press/98Folio.pdf soubai_fgsufgsu@kindle.com

```


and that's it.

## License:
[The MIT License](./LICENSE)