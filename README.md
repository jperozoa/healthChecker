# healthChecker

Do http request to an URL and print the status header of the response

## Use

``` shell
healthChecker https://localhost/health
```

## Options

You can use the insecure parameter to avoid checking the TLS certificate

``` shell
healthChecker -insecure https://localhost/health
```
