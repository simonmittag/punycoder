![Punycoder Logo](punycoder_80.png)
Punycoder converts international domain names to their ([RFC1035](https://www.rfc-editor.org/rfc/rfc1035.txt) compatible)
punycode representations on the commandline, so you can use them with `nslookup` or `dig`.

This is done by first converting the non-ASCII characters into their Unicode code points, and then encoding those code 
points using a special base-36 encoding scheme. The resulting string of ASCII characters is prefixed with the 
characters "xn--" to indicate that it is a Punycode-encoded domain name.

[![Circleci Builds](https://circleci.com/gh/simonmittag/punycoder.svg?style=shield)](https://circleci.com/gh/simonmittag/punycoder)
[![Github Issues](https://img.shields.io/github/issues/simonmittag/punycoder)](https://github.com/simonmittag/punycoder/issues)
[![Github Activity](https://img.shields.io/github/commit-activity/m/simonmittag/punycoder)](https://img.shields.io/github/commit-activity/m/simonmittag/punycoder)  
[![Go Version](https://img.shields.io/github/go-mod/go-version/simonmittag/punycoder)](https://img.shields.io/github/go-mod/go-version/simonmittag/punycoder)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://img.shields.io/badge/version-0.2.1-orange)](https://github.com/simonmittag/punycoder)

## What's new
### v0.2.2 (16/06/23)
* new logo

### v0.2.1 (25/05/23)
* refactored cli args.

### v0.2.0 (22/05/23)
* support unicode encoding of punycode

## Up and running
### Homebrew
```
brew tap simonmittag/cli &&
  brew install punycoder &&
  punycoder 
```

### Golang
```bash
git clone https://github.com/simonmittag/punycoder && cd punycoder && 
go install github.com/simonmittag/punycoder/cmd/punycoder && 
punycoder 
```

## Usage
```
λ Usage: punycoder host | [-u] host | [-v] | [-h]
  -h    print usage
  -u    convert to unicode
  -v    print punycoder version
```

## Examples

Use punycoder with bash pipes reading from stdin
```
echo 1❤️.ws | punycoder | nslookup
Server:		8.8.8.8
Address:	8.8.8.8#53

Non-authoritative answer:
Name:	xn--1-7iqv272g.ws
Address: 64.70.19.203
```

```
echo xn--1-7iqv272g.ws | punycoder -u 
1❤️.ws
```

## Contributions
The punycoder team welcomes all [contributors](https://github.com/simonmittag/punycoder/blob/master/CONTRIBUTING.md). Everyone interacting with the project's codebase, issue trackers, chat rooms and mailing lists
is expected to follow the [code of conduct](https://github.com/simonmittag/punycoder/blob/master/CODE_OF_CONDUCT.md)
