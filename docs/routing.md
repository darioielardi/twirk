---
id: "routing"
title: "HTTP Routing and Serialization"
sidebar_label: "How twirk routes requests"
---

Routing and Serialization is handled by twirk. All you really need to know is
that "it just works". However you may find this interesting and useful for
debugging.

### HTTP Routes

twirk works over HTTP 1.1; all RPC methods map to routes that follow the format:

```
POST /twirk/<package>.<Service>/<Method>
```

The `<package>` name is whatever value is used for `package` in the `.proto`
file where the service was defined. The `<Service>` and `<Method>` names are
CamelCased just as they would be in Go.

For example, to call the `MakeHat` RPC method on the example
[Haberdasher service](https://github.com/darioielardi/twirk/wiki/Usage-Example:-Haberdasher)
the route would be:

```
POST /twirk/twirk.example.haberdasher.Haberdasher/MakeHat
```

### Content-Type Header (json or protobuf)

The `Content-Type` header is required and must be either `application/json` or
`application/protobuf`. JSON is easier for debugging (particularly when making
requests with cURL), but Protobuf is better in almost every other way. Please
use Protobuf in real code. See
[Protobuf and JSON](https://github.com/darioielardi/twirk/wiki/Protobuf-and-JSON)
for more details.

### JSON serialization

The JSON format should match the
[official spec](https://developers.google.com/protocol-buffers/docs/proto3#json)'s
rules for JSON serialization. In a nutshell: names are `camelCased`, _all_
fields must be set, _no_ extra fields may be set, and `null` means "I want to
leave this field blank".

### Error responses

Errors returned by twirk servers use non-200 HTTP status codes and always have
JSON-encoded bodies (even if the request was Protobuf-encoded). The body JSON
has three fields `{type, msg, meta}`. For example:

```
POST /twirk/twirk.example.haberdasher.Haberdasher/INVALIDROUTE

404 Not Found
{
    "type": "bad_route",
    "msg": "no handler for path /twirk/twirk.example.haberdasher.Haberdasher/INVALIDROUTE",
    "meta": {"twirk_invalid_route": "POST /twirk/twirk.example.haberdasher.Haberdasher/INVALIDROUTE"}
}
```

## Making requests on the command line with cURL

It's easy to hand-write a twirk request on the command line.

For example, a cURL request to the Haberdasher example `MakeHat` RPC would look
like this:

```sh
curl --request "POST" \
     --location "http://localhost:8080/twirk/twirk.example.haberdasher.Haberdasher/MakeHat" \
     --header "Content-Type:application/json" \
     --data '{"inches": 10}' \
     --verbose
```

We need to signal twirk that we're sending JSON data (instead of protobuf), so
it can use the right deserializer. If we were using protobuf, the `--header`
would be `Content-Type:application/protobuf` (and `--data` a protobuf-encoded
message).

The `Size` request in JSON is `{"inches": 10}`, matching the Protobuf message
type:

```protobuf
message Size {
   int32 inches = 1;
}
```

The JSON response from twirk would look something like this (`--verbose` stuff
omitted):

```json
{"inches":1, "color":"black", "name":"bowler"}
```

Matching the Protobuf message type:

```protobuf
message Hat {
  int32 inches = 1;
  string color = 2;
  string name = 3;
}
```
