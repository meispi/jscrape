# jscrape
A CLI tool written in go to find subdomains and S3 buckets from js files.

## Installation
```
go install github.com/meispi/jscrape@latest
```

## How to use

```
  -all
        looks for *target* (and not just *.target.*) (default false)
  -c int
        number of threads (default 8)
  -s3
        looks for *.s3.amazonaws.com only (default false)
  -w string
        Enter your word (e.g. uber)
```

```
$ cat jsfiles.txt | jscrape -w <target_name>
$ echo "https://target.tld/somerandomjsfile.js" | jscrape -w <target>
```

e.g.:
`cat jsfiles.txt | jscrape -w uber`

You can pipe the output of jscrape to [httprobe](https://github.com/tomnomnom/httprobe) to filter the domains that resolve.

```
cat jsfiles.txt | jscrape -w uber | httprobe
```

To look for only s3 buckets in the js files:
```
$ cat jsfiles.txt | jscrape -s3
$ echo "https://target.tld/somerandomjsfile.js" | jscrape -s3
```
