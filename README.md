# subjs
A CLI tool written in go to find subdomains from js files.

## Installation
```
go install github.com/meispi/subjs@latest
```

## How to use

```
  -all
        looks for *target* (and not just *.target.*) (default false)
  -c int
        number of threads (default 8)
  -w string
        Enter your word (e.g. uber)
```

```
$ cat jsfiles.txt | subjs -w <target_name>
$ echo "https://target.tld/somerandomjsfile.js" | subjs -w <target>
```

e.g.:
`cat jsfiles.txt | subjs -w uber`

You can pipe the output of subjs to [httprobe](https://github.com/tomnomnom/httprobe) to filter the domains that resolve.

```
cat jsfiles.txt | subjs -w uber | httprobe
```