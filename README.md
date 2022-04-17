# subjs
A CLI tool written in go to find subdomains and S3 buckets from js files.

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
  -s3
        looks for *.s3.amazonaws.com only (default false)
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

To look for only s3 buckets in the js files:
```
$ cat jsfiles.txt | subjs -s3
$ echo "https://target.tld/somerandomjsfile.js" | subjs -s3
```
