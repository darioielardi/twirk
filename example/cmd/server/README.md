# twirk example server #

This binary is an example twirk server. It's meant mostly to be read
and to be used in conjunction with the neighboring client binary.

When a request is made, the server will log the statsd messages it
would have sent, so you'll see stuff like this:

    -> % ./server
    incr twirk.total.requests: 1 @ 1.000000
    incr twirk.MakeHat.requests: 1 @ 1.000000
    incr twirk.total.responses: 1 @ 1.000000
    incr twirk.MakeHat.responses: 1 @ 1.000000
    incr twirk.status_codes.total.200: 1 @ 1.000000
    incr twirk.status_codes.MakeHat.200: 1 @ 1.000000
    time twirk.all_methods.response: 370.695µs @ 1.000000
    time twirk.MakeHat.response: 370.695µs @ 1.000000
    time twirk.status_codes.all_methods.200: 370.695µs @ 1.000000
    time twirk.status_codes.MakeHat.200: 370.695µs @ 1.000000
