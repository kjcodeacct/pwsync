![pwsync](./assets/pwsync.png)

---
![License](https://img.shields.io/github/license/kjcodeacct/pwsync)
[![Go Report Card](https://goreportcard.com/badge/github.com/kjcodeacct/pwsync)](https://goreportcard.com/report/github.com/kjcodeacct/pwsync)
[![Build Status](https://cloud.drone.io/api/badges/kjcodeacct/pwsync/status.svg)](https://cloud.drone.io/kjcodeacct/pwsync)

# Table of Contents

- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Quick Start](#quick-start)
- [Installing & Usage](#installing--usage)
    - [Binary Releases](#binary-releases)
    - [Go Get](#go-get)
  - [Building](#building)
    - [Binaries](#binaries)
    - [CI/CD](#cicd)
  - [Usage](#usage)
  - [Configuration](#configuration)
    - [Fields](#fields)
- [Supported OS](#supported-os)
- [Supported Password Platforms](#supported-password-platforms)
- [Dependencies](#dependencies)
- [Issues & Updates](#issues--updates)

# Overview
Pwsync is a convenient password backup tool to help with the following:

* Backup proprietary password vaults into an encrypted [keepass](https://keepass.info/index.html) database.


If you work with multiple password systems, or want backups of system critical passwords, this is for you.

---

Please see the list of [supported password platforms](#supported-password-platforms).

---

**Security Note**

**ALL** interaction with a password service is done with it's native command line application, pwsync **does not** make API calls directly to a password service. This is enables far more flexibility and reduces security issues.

**ALL** exported password vaults from a given service **MUST** be in a unencrypted, CSV format.

By default any unencrypted csv files are cleaned up in a secure manner.

---
# Quick Start

First install your password services CLI application.
* See links available in [supported password platforms](#supported-password-platforms).

Next initialize your working directory.
```
$ pwsync init --platform=bitwarden
```

Login to your password service
```
$ pwsync login
```

Fetch any updates from you password service
```
$ pwsync fetch
```


Pull and backup your latest passwords
```
$ pwsync pull
```

Once a backup has been created, feel free to logout
```
$ pwsync logout
```

---

![cli demo](./assets/cli_demo.gif)

---
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


## Usage

Commands
* **init**
  * Initialize the working directory for pwsync, and create the configuration file. (default .pwsync.yaml)
  * Flags:

    **--platform**  platform to create a default cfg (lastpass,bitwarden)

* **login**
  * Login to the desired password platform, executes the 'login' command desginated by the password manager.
* **logout**
  * Logout of the desired password platform, executes the 'logout' command desginated by the password manager.
* **fetch**
  *  Fetches the latest password vault from desired password platform, executes the 'sync/fetch/update' command designated by the password manager.
* **pull**
  * Exports the latest password vault into a CSV, then converts it to a keepass database.
  * Flags:

    **--cleanup** :  auto cleanup pulled files (default true)


## Configuration

Pwsync uses a yaml file for configuration of commands.

### Fields
* **platform**

  Specifies your password platform to be used (bitwarden, lastpass, etc)

* **timeout**

  Specify the timeout waiting for vault updates, and exports, in seconds.

* **cmdList**

  List of commands, and their configurations, mapping to pwsync (login, logout, fetch, pull)

  * **name**

    Name of the command, this *must* match to a support subcommand of pwsync (login, logout, fetch, pull).

  * **cmd**

    This is the most critical configuration component, and matches to a password services command line application.
    If this is not specified, you will *not* be able to successfully backup anything.

    List of currently supported command line applications can be found here:

    * [Supported Password Platforms](#supported-password-platforms)

    Please note it is a *requirement* you have your password services command line application installed before configuring pwsync.

    For advanced usage with command line parameters, please see the detailed example below.

  *

A detailed example is below:

```yaml
# platform - specify the password platform used
platform: bitwarden
# timeout - timeout for waiting for password vault updates and exports, in seconds, defaults to 10 seconds
timeout: 10
# cmdList - list of commands mapping to pwsync (login, logout, fetch, pull)
cmdList:
- name: login
  # cmd - this specifies what executable to run, and the parameters to provide it
  cmd:
  # this executes the 'bw' command, either specify by name if it is in your PATH, or by absolute path
  - bw
  # this passes the command line argument 'login' to the executable 'bw'
  - login
  # this uses the environment variable 'PWSYNC_USERNAME' as an additional parameter
  # all environment variables used *MUST* be between curly braces {}
  - '{PWSYNC_USERNAME}'
- name: logout
  cmd:
  - bw
  - logout
- name: pull
  cmd:
  - bw
  - export
  # in the event your password management application exports to stdout instead of a file, this can be supplemented here
  stdoutFile: password_export.csv
- name: fetch
  cmd:
  - bw
  - sync

```

# Supported OS
Currently pwsync is limited to the following Operating Systems:

* Linux
* BSD
* Mac/Darwin

Unfortunately there is currently no windows support, as this application requires pseudo terminals, and the current package in use 'pty' does not support windows.

If you are an avid windows user and want to add this feature, please [publish a pull request](https://github.com/kjcodeacct/pwsync/pulls).

# Supported Password Platforms

* [Bitwarden](https://bitwarden.com/)
  * [Bitwarden CLI application](https://bitwarden.com/help/article/cli/)
* [Lastpass](https://www.lastpass.com/)
  * [Lastpass CLI application](https://github.com/lastpass/lastpass-cli)

If a password service you use is not availble please feel free to:
* [Create an issue](https://github.com/kjcodeacct/pwsync/issues)
* [Publish a pull request](https://github.com/kjcodeacct/pwsync/pulls)

# Dependencies
* Golang version 1.14+
* Unix based system (Linux, Mac, BSD, etc)

# Issues & Updates
If there is a breaking change with a platforms commandline application, please:
* [Create an issue](https://github.com/kjcodeacct/pwsync/issues)
  * Document your error in the issue

I will try to maintain this application, but I cannot account for every update
