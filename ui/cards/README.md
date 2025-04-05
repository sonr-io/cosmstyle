# `AccountCard`

A card that displays a user's account information.

## Properties


| Name | Type | Description |
| ---- | ---- | ----------- |
| `Address` | `string` | The address of the user's account. |
| `Name` | `string` | The name of the user. |
| `Handle` | `string` | The handle of the user. |
| `Block` | `string` | The block number of the user's account. |


## Usage

```templ
import "github.com/sonr-io/nebula/ui/cards"

templ View() {
  @cards.AccountCard(cards.AccountCardData{
    Address: "0x1234567890123456789012345678901234567890",
    Name: "Sonr",
    Handle: "sonr-io",
    Block: "1234567890",
  })
}
```

---

# `PermissionsCard`

A card that displays a user's permissions requested by a Service.

## Properties


