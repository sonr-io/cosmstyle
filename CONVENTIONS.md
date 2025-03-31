# Nebula Development Guidelines

## Build Commands
- Generate templ files: `task templ-gen` or `make generate`
- Build project: `make build` or `task nebula:build`
- Publish: `make release` or `task nebula:publish`

## Code Style Guidelines

### Go
- Use Go 1.23.4 standards
- Group imports: standard library first, then third-party packages
- Use meaningful variable/function names in camelCase (e.g., `shortenAddress`)
- Document public functions with comments
- Use error handling with proper context

### Templ
- Component files follow `name.templ` pattern
- Maintain consistent indentation in HTML templates
- Group related components in package directories
- Use clear component naming that describes functionality
- Generated Go files (name_templ.go) should not be edited manually

### HTML/CSS
- Use TailwindCSS classes for styling
- Follow BEM-like naming for custom CSS classes
- Maintain responsive design principles
- Custom components use the `sl-` prefix

### Git Workflow
- Commit messages follow conventional format (feat:, fix:, refactor:, etc.)
- Maintain clean git history (no merge conflicts in PRs)