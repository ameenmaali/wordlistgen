# wordlistgen

## What and why?
wordlistgen is a tool to pass a list of URLs and get back a list of relevant words for your wordlists. Wordlists are much more
effective when you take the application's context into consideration. wordlistgen pulls out URL components, such as subdomain names,
paths, query strings, etc. and spits them back to stdout so you can easily add them to your wordlists

# Installation
If you don't have Go installed, "go" do that!
```go get -u github.com/ameenmaali/wordlistgen```

## Usage
wordlistgen takes URLs and paths from stdin, of which you will most likely want in a file such as:
```
$ cat file.txt
https://google.com/home/?q=2&d=asd
http://my.site
/api/v2/auth/me?id=123
```
### Help
```
$ wordlistgen -h
Usage of cookiesplz:
  -fq
    	If enabled, filter out query strings (e.g. if enabled /?q=123 - q would NOT be included in results
  -qv
    	If enabled, include query string values (e.g. if enabled /?q=123 - 123 would be included in results
```

## Examples

Get unique URL components from a file of URLs and/or paths:

`cat hosts.txt | wordlistgen`

Get unique URL components from a file of URLs and/or paths, including query string values, and save to a file:

`cat hosts.txt | wordlistgen -qv > urlComponents.txt`

wordlistgen works at it's best when chained with other tools, such as [@tonnomnom's](https://tomnomnom) [waybackurls](https://github.com/tomnomnom/waybackurls) :

`cat hosts.txt | waybackurls | wordlistgen`