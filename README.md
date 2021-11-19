# linenoise-server

linenoise-server is a very simple web server that provides
[linenoise](https://git.sr.ht/~mcornick/linenoise) as a service.

## Usage

There is only one method, `GET /noise` called thus:

`curl http://HOSTNAME:8080/v1/noise?length=x&upper=x&lower=x&digit=x`

If any parameter is not specified, a default of 16 is used for `length`,
and `true` for `upper`, `lower` and/or `digit`.

## Contributing

Bug reports are welcome on Sourcehut at
https://todo.sr.ht/~mcornick/linenoise. Patches can be sent to
<~mcornick/linenoise@lists.sr.ht>.

This project is intended to be a safe, welcoming space for
collaboration, and contributors are expected to adhere to the
[Contributor Covenant](https://www.contributor-covenant.org/) code of
conduct.

## License

linenoise-server is available as open source under the terms of the MIT
License.
