# `cosmstyle` - Sonr Design System

A comprehensive monorepo for Sonr's design system, featuring Golang Templ components and web components for building blockchain applications.

## What's inside?

This monorepo includes several packages organized for building cohesive blockchain interfaces:

### Core Packages

- `ui`: Golang [Templ](https://templ.guide/) components for server-side rendered UIs
  - `cards`: Reusable card components for displaying blockchain data
  - More component categories to come
- `packages/shoelace`: Custom web components based on Shoelace, optimized for blockchain applications
- `packages/cosmes`: A tree-shakeable, framework-agnostic ESM alternative to CosmJS and Cosmos Kit

## Getting Started with Templ Components

Our Golang Templ components provide type-safe, efficient server-side rendering:

```go
import "github.com/sonr-io/design-system/ui/cards"

// In your handler
func renderProfilePage(w http.ResponseWriter, r *http.Request) {
    accountData := cards.AccountCardData{
        Address: "sonr1abcd...",
        Name: "Alice",
        Handle: "alice",
        Block: "42069"
    }
    
    component := cards.AccountCard(accountData)
    component.Render(r.Context(), w)
}
```

## Development

### Prerequisites

- Go 1.21+
- Templ CLI installed
- Node.js 18+ (for web components)

### Building Templ Components

Generate Go code from Templ templates:

```bash
templ generate
```

### Working with Web Components

Install dependencies and start the development server:

```bash
cd packages/shoelace
npm install
npm start
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
