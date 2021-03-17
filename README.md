![pwsync](./assets/pwsync.png)

---
![License](https://img.shields.io/github/license/kjcodeacct/pwsync)
[![Go Report Card](https://goreportcard.com/badge/github.com/kjcodeacct/pwsync)](https://goreportcard.com/report/github.com/kjcodeacct/pwsync)
[![Build Status](https://cloud.drone.io/api/badges/kjcodeacct/pwsync/status.svg)](https://cloud.drone.io/kjcodeacct/pwsync)

# Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Installing & Usage](#installing--usage)
    - [Binary Releases](#binary-releases)
    - [Go Get](#go-get)
  - [Building](#building)
    - [Binaries](#binaries)
    - [CI/CD](#cicd)
- [Supported Password Platforms](#supported-password-platforms)
- [Dependencies](#dependencies)

# Overview
Pwsync is a convenient password backup tool to help with the following:

* Backup proprietary password vaults into an encrypted [keepass](https://keepass.info/index.html) database.


If you work with multiple password systems, or value backups for system critical passwords, this could be for you.

# Installing & Usage

### Binary Releases
Binary releases are available in the github releases page found [here](https://github.com/kjcodeacct/pwsync/releases)

### Go Get
golang 1.14+ is required
set GO111MODULE=on

```
$ go get -u github.com/kjcodeacct/pwsync
```

## Building


### Binaries
If you would like to manually build binaries available in the [releases page](https://github.com/kjcodeacct/pwsync/releases), run the following.
```
$ make binaries
```

### CI/CD
If you want to view steps used by <drone.io> for automated builds please view [.drone.yml](.drone.yml)

# Supported Password Platforms

* [Bitwarden](https://bitwarden.com/)
* [Lastpass](https://www.lastpass.com/)

If a password service you use is not availble please feel free to:
* [Create an issue](https://github.com/kjcodeacct/pwsync/issues)
* [Publish a pull request](https://github.com/kjcodeacct/pwsync/pulls)

# Dependencies
Golang version 1.14+
Unix based system (Linux, Mac, BSD, etc)