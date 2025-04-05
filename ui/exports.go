// Package ui provides a comprehensive library of Golang Templ components
// for building modern, type-safe web applications with a focus on blockchain interfaces.
package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/sonr-io/nebula/ui/cards"
	"github.com/sonr-io/nebula/ui/dropdowns"
	"github.com/sonr-io/nebula/ui/inputs"
	"github.com/sonr-io/nebula/ui/layouts"
)

// ╭──────────────────────────────────────────────────────────╮
// │                Card Component Exports                    │
// ╰──────────────────────────────────────────────────────────╯

// AccountCardData represents the data needed for an AccountCard component
type AccountCardData = cards.AccountCardData

// AccountCard renders a user's account information in a card format
var AccountCard = cards.AccountCard

// ╭──────────────────────────────────────────────────────────╮
// │               Dropdown Component Exports                 │
// ╰──────────────────────────────────────────────────────────╯

// Coin represents a cryptocurrency with its ticker, name, and default status
type Coin = dropdowns.Coin

// AssetsDropdown renders a dropdown selector for blockchain assets
var AssetsDropdown = dropdowns.AssetsDropdown

// CoinsDropdown renders a dropdown selector for cryptocurrencies
var CoinsDropdown = dropdowns.CoinsDropdown

// CoinOption renders an individual option in the coins dropdown
var CoinOption = dropdowns.CoinOption

// ╭──────────────────────────────────────────────────────────╮
// │                Input Component Exports                   │
// ╰──────────────────────────────────────────────────────────╯

// HandleState represents the validation state of a handle input
type HandleState = inputs.HandleState

// HandleState constants
const (
	HandleStateInitial = inputs.HandleStateInitial
	HandleStateValid   = inputs.HandleStateValid
	HandleStateInvalid = inputs.HandleStateInvalid
)

// Handle input components
var (
	HandleInitial = inputs.HandleInitial
	HandleError   = inputs.HandleError
	HandleSuccess = inputs.HandleSuccess
)

// Human verification slider components
var (
	HumanInitial = inputs.HumanInitial
	HumanError   = inputs.HumanError
	HumanSuccess = inputs.HumanSuccess
)

// Passkey input components
var (
	PasskeyInitial = inputs.PasskeyInitial
	PasskeyError   = inputs.PasskeyError
	PasskeySuccess = inputs.PasskeySuccess
)

// ╭──────────────────────────────────────────────────────────╮
// │                Layout Component Exports                  │
// ╰──────────────────────────────────────────────────────────╯

// Layout components
var (
	// Base layout components
	Head      = layouts.Head
	Root      = layouts.Root
	Container = layouts.Container
	Columns   = layouts.Columns
	Rows      = layouts.Rows
	Separator = layouts.Separator
	Clsx      = layouts.Clsx

	// Form layout components
	Form   = formWithChildren
	Header = layouts.Header
	Body   = layouts.Body
	Footer = layouts.Footer
	Cancel = layouts.Cancel
	Submit = layouts.Submit

	// Third-party library components
	Tailwind  = layouts.Tailwind
	Alpine    = layouts.Alpine
	Dexie     = layouts.Dexie
	Htmx      = layouts.Htmx
	Turnstile = layouts.Turnstile
	Nebula    = layouts.Nebula
)

// Helper function to make Form accept children
func formWithChildren(action, id string, children ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		// Render form opening tag
		err := layouts.Form(action, id).Render(ctx, w)
		if err != nil {
			return err
		}

		// Render all children
		for _, child := range children {
			err := child.Render(ctx, w)
			if err != nil {
				return err
			}
		}

		// Close the form tag
		_, err = io.WriteString(w, "</form>")
		return err
	})
}
