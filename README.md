# Filternet Checker
A simple CLI application for checking whether provided domains are censored or not.

## Usage
You just need to provide a configuration file that contains DNS resolvers of different ISPs and a list of domains to check.
1. First of all, you have to build the app. \
`go mod download && go build . -o filternet`
2. Make your own configuration file or use the default configuration
3. Provide some domains and check them.\
`./filternet -domain youtube.com,whatsapp.com -config default.json`

```
$ ./filternet -help
Usage of ./filternet:
  -config string
    	Config file path which contains resolvers (default "default.json")
  -domains string
    	Domains to check. Separated with, (default "twitter.com,instagram.com,facebook.com")
  -no-color
    	Disable color output
```

## Configuration file
Currently, DNS-over-UDP and DNS-over-HTTPS resolvers are supported.\
Each resolver object must have `name` and `servers` keys. 
- `name`: Name of the provider.
- `servers`: list of resolvers. Note that UDP addresses must contain a port and HTTPS addresses must start with `https://`

***Examples are available in the `default.json` file.***
