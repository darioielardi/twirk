---
id: "hooks"
title: "Server Hooks"
sidebar_label: "Server Hooks"
---

The service constructor can use `hooks *twirk.ServerHooks` to plug in additional
functionality:

```go
func NewHaberdasherServer(svc Haberdasher, hooks *twirk.ServerHooks) http.Handler
```

These _hooks_ provide a framework for side-effects at important points while a
request gets handled. You can do things like log requests, record response times
in statsd, authenticate requests, and so on.

To enable hooks for your service, you pass a `*twirk.ServerHooks` in to the
server constructor. The `ServerHooks` struct holds a pile of callbacks, each of
which receives the current context.Context, and can provide a new
context.Context (possibly including a new value through the
[`context.WithValue`](https://godoc.org/golang.org/x/net/context#WithValue)
function).

Check out
[the godoc for `ServerHooks`](http://godoc.org/github.com/darioielardi/twirk#ServerHooks)
for information on the specific callbacks. For an example hooks implementation,
[`github.com/darioielardi/twirk/hooks/statsd`](https://github.com/darioielardi/twirk/blob/master/hooks/statsd/)
is a good tutorial.
