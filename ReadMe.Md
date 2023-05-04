# punyc😀der 
[![Circleci Builds](https://circleci.com/gh/simonmittag/punycoder.svg?style=shield)](https://circleci.com/gh/simonmittag/punycoder)
[![Github Issues](https://img.shields.io/github/issues/simonmittag/punycoder)](https://github.com/simonmittag/punycoder/issues)
[![Github Activity](https://img.shields.io/github/commit-activity/m/simonmittag/punycoder)](https://img.shields.io/github/commit-activity/m/simonmittag/punycoder)  
[![Go Version](https://img.shields.io/github/go-mod/go-version/simonmittag/punycoder)](https://img.shields.io/github/go-mod/go-version/simonmittag/punycoder)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://img.shields.io/badge/version-0.1.0-orange)](https://github.com/simonmittag/punycoder)

## TLDR
Punycoder converts international domain names to their punycode representations on the commandline.

## Longer Version
IDNA (Internationalized Domain Name in Applications) is a set of standards and protocols used to support non-ASCII 
characters in domain names. Punycode is a method used in IDNA to encode non-ASCII characters into ASCII characters that 
can be used in domain names, making them backwards compatible with existing implementations. 

This is done by first converting the non-ASCII characters into their Unicode code points, and then encoding those code points using 
a special base-36 encoding scheme. The resulting string of ASCII characters is prefixed with the characters "xn--" to 
indicate that it is a Punycode-encoded domain name. 

For example, the Punycode representation of a legal Chinese domain 
name `中国互联网.com`  is `xn--fiqs8s2qjod38g.com`.

## What's new
### v0.1.0 (05/05/23)
* upgrade to go 1.20

### v0.0.11 (04/05/23)
* added goreleaser and made available on homebrew

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
λ punycoder 中国互联网.com
  xn--fiqs8s2qjod38g.com
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

## Contributions
The punycoder team welcomes all [contributors](https://github.com/simonmittag/punycoder/blob/master/CONTRIBUTING.md). Everyone interacting with the project's codebase, issue trackers, chat rooms and mailing lists
is expected to follow the [code of conduct](https://github.com/simonmittag/punycoder/blob/master/CODE_OF_CONDUCT.md)
