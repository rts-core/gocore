<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/rts-core/gocore">
    <img src="images/logo.png" alt="Logo" height="120">
  </a>

  <p align="center">
    Reusable helper functionality for RTS golang applications
    <br />
    <a href="https://github.com/rts-core/gocore/issues">Report Bug</a> |
    <a href="https://github.com/rts-core/gocore/issues">Request Feature</a>
  </p>
</p>

### Built With

* [Go 1.18](https://go.dev//)

<!-- GETTING STARTED -->
## Getting Started

To reference this library:

```shell
go get github.com/rts-core/gocore
```

NOTE: this is a private repository, this may require you to do private repo workarounds. Notes provided below.

## Documentation

Documentation is hosted as standard [Go Documentation](). 

## Working with Private Repositories in Go

Private repos are easiest to work with through either 2-factor authentication or SSH. Setting up either on there system
is left as an exercise for the reader (and google) as this will change from OS to OS and version to version.

However, getting Git to play nicely with this setup is often much harder to manage, so these steps are documented here.

### HTTPS issues

To get github to stop attempting all calls as https calls use the following command:
```shell
git config --global --add url."git@github.com:".insteadOf "https://github.com/"
```
this will do the correct replacement. This adds the following to your ~/.gitconfig and can be removed from there if you
want to revert the change.
```text
[url "git@github.com:"]
	insteadOf = https://github.com/
```

### Version does not exist issue

The following command will tell go to treat the repos with the following names as private repos and not
to check the checksum on them.

```shell
go env -w GOPRIVATE=github.com/rts-core/*
```

Note: this will replace your currently existing GOPRIVATE variable. If for example you are in need of multiple entries
GOPRIVATE is comma delimited and you can extend your current GOPRIVATE with something like:

```shell
go env -w GOPRIVATE=$(go env GOPRIVATE),github.com/rts-core/*

```

<!-- CONTACT -->
## Contact

Project Link: [https://github.com/rts-core/gocore](https://github.com/rts-core/gocore)

Copyright [2022] [Real Technology Solutions LLC]