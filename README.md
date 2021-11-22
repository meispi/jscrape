# subjs
A CLI tool written in go to find subdomains from js files.

## Installation
```
go install github.com/meispi/subjs@latest
```

## How to use
```
cat jsfiles.txt | subjs -w <target_name>
```

e.g.:
`cat jsfiles.txt | subjs -w uber`

You can pipe the output of subjs to [httprobe](https://github.com/tomnomnom/httprobe) to filter the domains that resolve.

```
cat jsfiles.txt | subjs -w uber | httprobe
```