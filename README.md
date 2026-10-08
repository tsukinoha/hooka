# Hooka

Hooka [hu:kɑ] is a Go library for sending messages to Microsoft Teams through
[Workflows](https://support.microsoft.com/office/creating-a-workflow-from-a-channel-in-teams-242eb8f2-f328-45be-b81f-9817b51a5f0e)
(Power Automate) webhooks.

It builds [Adaptive Cards](https://adaptivecards.io/) with a small typed API and posts them as Teams messages.

## Features

- Build Adaptive Cards with `TextBlock`, `Image`, `ImageSet`, `FactSet`, `Container`, `ColumnSet` and `Column`
- Setters accept values case-insensitively and drop invalid ones, so the card JSON stays valid
- The card's `version` is raised automatically to the highest version of the elements it contains
- Webhook URLs are checked: only `https` URLs on `logic.azure.com` or `powerplatform.com` are accepted
- `context.Context` support for cancellation and timeouts
- No dependencies outside the Go standard library

## Requirements

- Go 1.26 or later
- A Teams Workflows webhook URL. In Teams, open the channel's **Workflows** menu and create a flow from the
  "Send webhook alerts to a channel" template (or a similar one) to get the URL.

## Installation

```sh
go get github.com/tsukinoha/hooka
```

## Usage

```go
package main

import (
	"encoding/json"
	"log"

	"github.com/tsukinoha/hooka"
	"github.com/tsukinoha/hooka/adaptive_card"
)

func main() {
	teams, err := hooka.NewTeams("https://xxxx.logic.azure.com/workflows/xxxx")
	if err != nil {
		log.Fatal(err)
	}

	// Build a card.
	card := adaptive_card.New()
	card.SetLang("en")

	title := adaptive_card.NewTextBlock("Deployment finished")
	title.SetSize("large")
	title.SetWeight("bolder")
	card.Append(title)

	facts := adaptive_card.NewFactSet()
	facts.Append(adaptive_card.NewFact("Service", "api-server"))
	facts.Append(adaptive_card.NewFact("Status", "Success"))
	card.Append(facts)

	// Attach the card and send it.
	teams.Attach(card)
	data, err := json.Marshal(teams)
	if err != nil {
		log.Fatal(err)
	}
	if err := teams.Send(data); err != nil {
		log.Fatal(err)
	}
}
```

### Cancellation and timeouts

`SendContext` cancels the request when the context is done. `Send` uses a 30-second timeout.

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
if err := teams.SendContext(ctx, data); err != nil {
	log.Fatal(err)
}
```

### Sending your own JSON

`Send` posts any JSON payload, so you can also send a message you built yourself:

```go
err := teams.Send([]byte(`{"type":"message","attachments":[ ... ]}`))
```

### Layouts

`Container` and `Column` hold other elements, and `ColumnSet` places `Column`s side by side:

```go
left := adaptive_card.NewColumn()
left.SetWidth("auto")
left.Append(adaptive_card.NewImage("https://example.com/icon.png"))

right := adaptive_card.NewColumn()
right.SetWidth("stretch")
right.Append(adaptive_card.NewTextBlock("Hello, Teams!"))

columns := adaptive_card.NewColumnSet()
columns.Append(left)
columns.Append(right)
card.Append(columns)
```

`Column.SetWidth` accepts `"auto"`, `"stretch"` or a number as a string (for example `"2"`), which sets a relative width.

### Errors

`Send` and `SendContext` return an error when the request fails or Teams responds with a non-2xx status.
The error message includes the status and the beginning of the response body.

## Elements

| Element      | Constructor            | Setters |
|--------------|------------------------|---------|
| `TextBlock`  | `NewTextBlock(text)`   | `SetColor`, `SetSize`, `SetWeight`, `SetHorizontalAlignment`, `SetWrap`, `SetMaxLines`, `SetSubtle`, `SetSpacing`, `SetSeparator`, `SetId` |
| `Image`      | `NewImage(url)`        | `SetAltText`, `SetSize`, `SetStyle`, `SetHorizontalAlignment`, `SetSpacing`, `SetSeparator`, `SetId` |
| `ImageSet`   | `NewImageSet()`        | `Append`, `SetImageSize`, `SetSpacing`, `SetSeparator`, `SetId` |
| `FactSet`    | `NewFactSet()`         | `Append(NewFact(title, value))`, `SetSpacing`, `SetSeparator`, `SetId` |
| `Container`  | `NewContainer()`       | `Append`, `SetStyle`, `SetSpacing`, `SetSeparator`, `SetId` |
| `ColumnSet`  | `NewColumnSet()`       | `Append`, `SetHorizontalAlignment`, `SetSpacing`, `SetSeparator`, `SetId` |
| `Column`     | `NewColumn()`          | `Append`, `SetWidth`, `SetStyle`, `SetSpacing`, `SetSeparator`, `SetId` |

See the [Adaptive Cards schema explorer](https://adaptivecards.io/explorer/) for what each property means.
An invalid value passed to a setter clears the property, so Teams uses its default.

## Development

```sh
go vet ./...
go test ./...
```

## License

Hooka is distributed under the MIT License. See [LICENSE](LICENSE) or https://opensource.org/licenses/mit-license.php.

