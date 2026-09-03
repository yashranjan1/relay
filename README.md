# Relay - Ditch the Electron Bloat with Style

[![tests](https://github.com/yashranjan1/relay/actions/workflows/ci-cd.yml/badge.svg?branch=main)](https://github.com/yashranjan1/relay/actions/workflows/ci-cd.yml)

> **Note:** This project is a fork of [overhttps/req](https://github.com/overhttps/req).
> Active development from me is moving here for the foreseeable future.

## About

`relay` is a lightweight, terminal-based API client.
**Current Status**: Early development (alpha). Core HTTP execution features are still in progress.

The goal is to provide a fast, minimal terminal interface for creating, sending,
and inspecting HTTP requests interactively from the command line.

> The app works completely offline with no external dependencies required.

Read the full docs over here -
[Relay Docs](https://yashranjan1.github.io/relay/)

## Installation

Install `relay` using `go install`:

```bash
# Install the latest stable release
go install github.com/yashranjan1/relay@latest

# Or install a specific version (e.g., v0.1.0-alpha.4)
go install github.com/yashranjan1/relay@v0.1.0-alpha.4
```

### Requirements

- Go version 1.24.4

## Usage

After installing `relay`, you can run it using this command.

```
relay
```

## Libraries Used

### Terminal UI (by Charm.sh)

- [bubbletea](https://github.com/charmbracelet/bubbletea) — A powerful, fun TUI
  framework for Go
- [bubbles](https://github.com/charmbracelet/bubbles) — Pre-built components for
  TUI apps
- [lipgloss](https://github.com/charmbracelet/lipgloss) — Terminal style/layout
  DSL
