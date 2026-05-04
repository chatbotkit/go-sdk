# Portable Embedded Agent

Demonstrates shipping a ChatBotKit agent as a single portable executable.

This example bundles:

- Embedded skills loaded from `embed.FS`
- An obfuscated bot ID
- An obfuscated scoped token that can only mint sessions for one bot
- A stateful agent that runs with the short-lived session token returned by ChatBotKit

## What it demonstrates

- Embedding a `skills/` directory at compile time with `//go:embed`
- Obfuscating the bot ID and scoped session-minting token with `go generate`
- Creating a temporary bot session with `client.Bot.Session.Create`
- Reusing the returned `conversationId` and `token` for stateful execution
- Producing a single binary with no `.env` dependency at runtime

## Important security note

This is a portability-first pattern, not strong secret protection.

The embedded bot ID and scoped token are hidden from casual inspection and from tools like `strings`, but anyone who controls the binary can eventually recover what the binary can use. The safety comes from using a narrowly scoped token that can only mint temporary sessions for one bot.

## Usage

1. Create a scoped token in ChatBotKit that only allows this route:

```text
bot/YOUR_BOT_ID/session/create
```

2. Regenerate the embedded secret file with your real values:

```bash
export CHATBOTKIT_BOT_ID="YOUR_BOT_ID"
export CHATBOTKIT_SESSION_MINT_TOKEN="YOUR_SCOPED_SESSION_MINT_TOKEN"

go generate .
```

3. Build the binary:

```bash
go build -trimpath -ldflags="-s -w" -o portable-agent .
```

4. Run it:

```bash
./portable-agent "Review this Go function: func add(a, b int) int { return a - b }"
```

If you skip step 2, the example still compiles, but it exits at runtime because the checked-in placeholder credentials are not usable.

## Check that the values are not present as plain strings

```bash
strings ./portable-agent | grep "YOUR_BOT_ID"
strings ./portable-agent | grep "YOUR_SCOPED_SESSION_MINT_TOKEN"
```

Those placeholder strings should not appear after you regenerate `secrets_gen.go` with real values and build again.
