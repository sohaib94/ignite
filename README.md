# Ignite <!-- omit in toc -->

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Get Started](#get-started)
  - [Install](#install)
  - [Run](#run)
- [Other TODO](#other-todo)
- [Contributing](#contributing)

## Introduction

Go cli tool to kickstart new project based on an input yaml file.

This includes the ability to:
- Locally
  - [x] Create directory for repo
  - [x] Create any subdirectories and empty files in those directories
  - [ ] Clone entire repos into a directory
  - [ ] Clone specific directories / files from repo
- Remote repo
  - [ ] Initialise personal remote repo
  - [ ] Initialise organisational remote repo
  - [ ] Init commit with created local repo
  - [ ] Setup webhooks for tools with repo
    - [ ] Slack
    - [ ] Monday board

## Prerequisites

- Golang 1.22

## Get Started

### Install

`go install github.com/sohaib94/ignite`

### Run

1. Create an ignite with a structure following [this example](./testdata/ignite_example.yml)
2. Run `ignite generate -f <path_to_ignite_file> -o <path_to_output_dir>`

## Other TODO

- [ ] Setup deploying and allowing install via `brew`
- [ ] Setup Github Issues to track work

## Contributing

Feel free to clone down and raise a PR for any changes
