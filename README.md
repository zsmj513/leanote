# leanote

[![Join the chat at https://gitter.im/humboldtux/leanote](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/humboldtux/leanote?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

[Leanote](http://leanote.org/) Go command line app

# How to use it

This is beta software for now, i will update these sections later.
For now, consider looking at the `Development` section.

## Installation

TODO

## Usage

```
LEANOTE_EMAIL="your@email.com" LEANOTE_PWD="secret" leanote u
```

By default, the cli use `https://leanote.com/api` as host API. If you want to use your own Leanote API:

```
LEANOTE_APIURL="https://yourleanote.com/api" LEANOTE_EMAIL="your@email.com" LEANOTE_PWD="secret" leanote u
```

# Development

Leanote sync mechanism refer to [evernote](http://dev.evernote.com/media/pdf/edam-sync.pdf).

```
git clone https://github.com/humboldtux/leanote.git
cd leanote
```

Prepare pre-requisites:

```
make prepare
```

Installation:

```
make
```

# TODO

- [ ] Better error handling
- [ ] Better api response handling
- [ ] Tests
- [ ] Hash passwords
- [ ] Configuration from file
- [ ] By default, don't list deleted items
- [ ] Usn and Seq are not used properly

Auth API:
- [ ] Logout at the end of cli execution?

User API:
- [ ] UpdateLogo not working

Notebooks API:
- [ ] update not working

Notes API:
- [ ] Add note not implemented
- [ ] Update note not implemented

Files API:
- [ ] GetImage not implemented
- [ ] GetAttach not implemented
- [ ] GetAllAttachs not implemented
