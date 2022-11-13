# go-healthcheck-action
> A example with simple GitHub action that contains a minimal Go application to test URL response code.
Uses [Container Action](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-a-docker-container-action). 

[![Build](https://github.com/atrakic/go-healthcheck-action/actions/workflows/go.yml/badge.svg)](https://github.com/atrakic/go-healthcheck-action/actions/workflows/go.yml)
[![Integration Test](https://github.com/atrakic/go-healthcheck-action/actions/workflows/integration.yml/badge.svg)](https://github.com/atrakic/go-healthcheck-action/actions/workflows/integration.yml)

## Usage

### Example workflow

```yaml
name: My Workflow
on: [push, pull_request]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: Run action
      uses: atrakic/go-healthcheck-action@main
      with:
        url: http://github.com
```

### Inputs

| Input                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `url`  | An mandatory input    | Url to test

### Outputs

| Output                                             | Description                                        |
|------------------------------------------------------|-----------------------------------------------|
| `statusCode`  | An example output (returns '200')    | https://httpwg.org/specs/rfc9110.html#overview.of.status.codes

### Using outputs

How to use your outputs in another action:

```yaml
steps:
- uses: actions/checkout@master
- name: Run action
  id: myaction
  uses: atrakic/go-healthcheck-action@main
  with:
    url: http://github.com

- name: Check outputs
    run: |
    echo "Outputs - ${{ steps.myaction.outputs.statusCode }}"
```
