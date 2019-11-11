# yaml2json

A small tool for converting YAML files into JSON

## Installation

```
curl -sLo yaml2json https://github.com/ssuareza/yaml2json/releases/download/v0.0.1/yaml2json-v0.0.1-darwin-amd64
chmod +x yaml2json
sudo mv yaml2json /usr/local/bin/
```

## Usage

### stdin pipe:

`cat file.yml | yaml2json`

### or specify a file:

`yaml2json path/file.yml`
