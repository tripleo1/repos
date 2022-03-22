# repos
Download repos to a directory straight from your browser

## Usage 

Create a bookmarklet in your browser toolbar (set to visible if not enabled) with the following configuration:

* Name: Clone to repos
* URL: javascript:location.href = "http://127.0.0.1:3000/clone?q=" + location.href;

## Installation

I forgot how to compile it.  It's probably `git clone` and `go build` and `cp -a repos ~repos/bin/`

```sh
adduser repos
mkdir ~repos/Repos
chown repos.repos ~repos/Repos
systemctl enable repos.service
```
