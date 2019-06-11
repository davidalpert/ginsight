![Ginsight](/docs/img/ginsight-logo.png)

Ginsight is a go API client with a CLI wrapper to talk to the Jira Insight API

# Building

```bash
# Install dependencies
$ go get -d ./...

# Build
$ ./build.sh
```

# CLI Usage

## Configuration

Create the configuration file at `~/.gapi.yaml` with your Jira details. See `.gapi.example.yaml`.

```yaml
jira:
  base_url: "https://jira.example.com"
  username: "username"
  password: "passwordGoesHere"
tls:
  insecure: false
debug: false
```

Test the configuration with `ginsight config --validate`

```
$ ginsight config --validate
Looking for config file: 
Using config file: /Users/jneufeld/.gapi.yaml
Validating your Insight client configuration...

Successfully authenticated to https://jira.example.com as 'jneufeld' (Jordan Neufeld)
```