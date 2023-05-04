# punyc😀der 

Converts international domain names ("idna") to their punycode representations on the commandline

[![Circleci Builds](https://circleci.com/gh/simonmittag/punycoder.svg?style=shield)](https://circleci.com/gh/simonmittag/punycoder)
[![Github Issues](https://img.shields.io/github/issues/simonmittag/punycoder)](https://github.com/simonmittag/punycoder/issues)
[![Github Activity](https://img.shields.io/github/commit-activity/m/simonmittag/punycoder)](https://img.shields.io/github/commit-activity/m/simonmittag/punycoder)  
[![Go Version](https://img.shields.io/github/go-mod/go-version/simonmittag/punycoder)](https://img.shields.io/github/go-mod/go-version/simonmittag/punycoder)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## What's new
added goreleaser

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
λ punycoder d😀main.com
  xn--dmain-z224d.com
```

## Contributions
The punycoder team welcomes all [contributors](https://github.com/simonmittag/punycoder/blob/master/CONTRIBUTING.md). Everyone interacting with the project's codebase, issue trackers, chat rooms and mailing lists
is expected to follow the [code of conduct](https://github.com/simonmittag/punycoder/blob/master/CODE_OF_CONDUCT.md)
