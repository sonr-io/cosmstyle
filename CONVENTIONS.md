# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands
- `task gen:templ`: Generate templ files
- `task build`: Build Nebula (runs build:cosmes and build:shoelace)
- `pnpm build` (in packages/cosmes): Build package to dist folder
- `pnpm dev` (in packages/cosmes): Watch and rebuild on changes

## Test Commands
- `pnpm test` (in packages/cosmes): Run unit tests
- `pnpm test:suite` (in packages/cosmes): Run full test suite (lint, typecheck, tests)

## Code Style
- **Go**: Follow Go 1.23.4 standards; group imports (stdlib first); use meaningful camelCase names
- **Templ**: Files follow `name.templ` pattern; don't edit generated `*_templ.go` files
- **UI**: Use TailwindCSS, Alpine.js, and Shoelace components (`sl-` prefix)
- **Error Handling**: Return errors with proper context; use consistent patterns
- **Git**: Use conventional commit format (feat:, fix:, docs:, etc.)

## Project Structure
- `ui/`: Frontend components organized by type (cards, inputs, layouts, etc.)
- `packages/`: Core SDK packages including cosmes and shoelace
- `tests/`: Test fixtures and integration tests