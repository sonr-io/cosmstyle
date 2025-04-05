package ui_test

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/sonr-io/nebula/ui"
)

// Helper function to render components to string
func renderComponent(t *testing.T, component templ.Component) string {
	t.Helper()
	var sb strings.Builder
	err := component.Render(context.Background(), &sb)
	if err != nil {
		t.Fatalf("Failed to render component: %v", err)
	}
	return sb.String()
}

// Test the AccountCard component
func TestAccountCard(t *testing.T) {
	data := ui.AccountCardData{
		Address: "sonr1abc123def456ghi789jkl",
		Name:    "Test User",
		Handle:  "testuser",
		Block:   "12345",
	}

	html := renderComponent(t, ui.AccountCard(data))

	// Simple assertions to ensure the component renders correctly
	if !strings.Contains(html, data.Handle) {
		t.Errorf("AccountCard should contain handle: %s", data.Handle)
	}

	if !strings.Contains(html, data.Block) {
		t.Errorf("AccountCard should contain block number: %s", data.Block)
	}

	if !strings.Contains(html, data.Name) {
		t.Errorf("AccountCard should contain name: %s", data.Name)
	}
}

// Test the AssetsDropdown component
func TestAssetsDropdown(t *testing.T) {
	html := renderComponent(t, ui.AssetsDropdown())

	// Check for some common assets
	assets := []string{"Sonr", "Bitcoin", "Ethereum"}
	for _, asset := range assets {
		if !strings.Contains(html, asset) {
			t.Errorf("AssetsDropdown should contain asset: %s", asset)
		}
	}
}

// Test the CoinsDropdown component
func TestCoinsDropdown(t *testing.T) {
	html := renderComponent(t, ui.CoinsDropdown())

	// Simple assertions to ensure the component renders correctly
	if !strings.Contains(html, "custom-tag") {
		t.Errorf("CoinsDropdown should have the custom-tag class")
	}

	// Check for some common coins
	coins := []string{"SNR", "BTC", "ETH"}
	for _, coin := range coins {
		if !strings.Contains(html, coin) {
			t.Errorf("CoinsDropdown should contain coin: %s", coin)
		}
	}
}

// Test the CoinOption component
func TestCoinOption(t *testing.T) {
	// Test default coin
	defaultCoin := ui.Coin{
		Ticker:    "SNR",
		Name:      "Sonr",
		IsDefault: true,
	}

	html := renderComponent(t, ui.CoinOption(defaultCoin))

	if !strings.Contains(html, "selected") {
		t.Errorf("Default coin option should be selected")
	}

	if !strings.Contains(html, defaultCoin.Name) {
		t.Errorf("CoinOption should contain coin name: %s", defaultCoin.Name)
	}

	// Test non-default coin
	nonDefaultCoin := ui.Coin{
		Ticker:    "ATOM",
		Name:      "Cosmos",
		IsDefault: false,
	}

	html = renderComponent(t, ui.CoinOption(nonDefaultCoin))

	if strings.Contains(html, "selected") {
		t.Errorf("Non-default coin option should not be selected")
	}

	if !strings.Contains(html, nonDefaultCoin.Name) {
		t.Errorf("CoinOption should contain coin name: %s", nonDefaultCoin.Name)
	}
}

// Test the HandleInput components
func TestHandleInput(t *testing.T) {
	// Test the initial state
	html := renderComponent(t, ui.HandleInitial())

	if !strings.Contains(html, "handle") || !strings.Contains(html, "at-sign") {
		t.Errorf("HandleInitial should contain handle input with at-sign icon")
	}

	// Test the error state
	html = renderComponent(t, ui.HandleError("testhandle", "Handle already taken"))

	if !strings.Contains(html, "testhandle") {
		t.Errorf("HandleError should contain the provided value")
	}

	if !strings.Contains(html, "Handle already taken") {
		t.Errorf("HandleError should contain the provided error message")
	}

	// Test the success state
	html = renderComponent(t, ui.HandleSuccess("testhandle"))

	if !strings.Contains(html, "testhandle") || !strings.Contains(html, "disabled") {
		t.Errorf("HandleSuccess should contain the provided value and be disabled")
	}
}

// Test the HumanSlider components
func TestHumanSlider(t *testing.T) {
	// Test the initial state
	html := renderComponent(t, ui.HumanInitial(3, 4))

	if !strings.Contains(html, "What is 3 + 4?") {
		t.Errorf("HumanInitial should contain the correct math question")
	}

	// Test the error state
	html = renderComponent(t, ui.HumanError(2, 5))

	if !strings.Contains(html, "What is 2 + 5?") {
		t.Errorf("HumanError should contain the correct math question")
	}

	if !strings.Contains(html, "Invalid Human Sum") {
		t.Errorf("HumanError should contain the error message")
	}

	// Test the success state
	html = renderComponent(t, ui.HumanSuccess())

	if !strings.Contains(html, "Success! Welcome Human") {
		t.Errorf("HumanSuccess should contain success message")
	}

	if !strings.Contains(html, "disabled") {
		t.Errorf("HumanSuccess should be disabled")
	}
}

// Test the PasskeyInput components
func TestPasskeyInput(t *testing.T) {
	addr := "sonr1abc123"
	handle := "testuser"
	challenge := "randomchallenge"

	// Test the initial state
	html := renderComponent(t, ui.PasskeyInitial(addr, handle, challenge))

	if !strings.Contains(html, "Register Passkey") {
		t.Errorf("PasskeyInitial should contain the Register Passkey text")
	}

	// Test the error state
	html = renderComponent(t, ui.PasskeyError(addr, handle, challenge))

	if !strings.Contains(html, "text-red-500") {
		t.Errorf("PasskeyError should contain a red icon")
	}

	// Test the success state
	html = renderComponent(t, ui.PasskeySuccess(addr, handle, challenge))

	if !strings.Contains(html, "text-green-500") {
		t.Errorf("PasskeySuccess should contain a green icon")
	}
}

// Test the Form layout components
func TestFormLayouts(t *testing.T) {
	// Test Form
	html := renderComponent(t, ui.Form("/submit", "test-form", templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := w.Write([]byte("Form Content"))
		return err
	})))

	if !strings.Contains(html, "Form Content") || !strings.Contains(html, "test-form") {
		t.Errorf("Form should render its children and have the correct ID")
	}

	// Test Submit
	html = renderComponent(t, ui.Submit("Submit Form"))

	if !strings.Contains(html, "Submit Form") {
		t.Errorf("Submit should contain the provided text")
	}

	if !strings.Contains(html, "arrow-right") {
		t.Errorf("Submit should contain the arrow icon")
	}
}

// Test the third-party library components
func TestThirdPartyLibraries(t *testing.T) {
	// Test the Head component which includes the third-party libraries
	html := renderComponent(t, ui.Head("Test Page", "0.0.11"))

	// Check that all the required scripts are included
	libraries := []string{
		"tailwindcss.com",
		"alpinejs",
		"dexie",
		"htmx.org",
		"nebula",
	}

	for _, lib := range libraries {
		if !strings.Contains(html, lib) {
			t.Errorf("Head should include %s script", lib)
		}
	}
}
