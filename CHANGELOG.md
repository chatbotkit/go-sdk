# Changelog

All notable changes to the ChatBotKit Go SDK are documented in this file. The
format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and
this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

For releases prior to `0.2.0`, see the auto-generated notes on the
[GitHub Releases](https://github.com/chatbotkit/go-sdk/releases) page.

## [0.4.0] - 2026-06-27

### Added

- Secret token minting and request proxying. `SecretClient.Mint` /
  `ContactSecretClient.Mint` mint a usable token from a secret (`oauth`/`jwt`
  secrets only; owner-only) and return `{ Token, ExpiresAt }`.
  `SecretClient.Proxy` / `ContactSecretClient.Proxy` proxy a request through a
  secret — the credential is injected server-side (it never leaves the platform)
  and the raw upstream `*http.Response` is returned verbatim. A non-2xx status
  (including `409 authorization_required`) is returned, not an error; the caller
  closes `resp.Body`. Backed by the new internal `httpclient.DoRaw`.

## [0.3.0] - 2026-06-26

### Added

- `State` lifecycle field on the `Skillset` and `Ability` resources, backed by
  the new `ResourceState` enum (`ResourceStateEnabled` / `ResourceStateDisabled`).
  A skillset or ability can now be toggled off without deleting it. Available on
  the create, update, fetch, and list types.

## [0.2.0] - 2026-06-22

### Added

- `SkillServer` integration client (`client.Integration.SkillServer`) with
  `List`, `Fetch`, `Create`, `Update`, and `Delete`. The Skill Server
  integration exposes a skillset's abilities as a text-first HTTP API.
- `Site` client under `Space` (`client.Space.Site`) with `List`, `Fetch`,
  `Create`, `Update`, and `Delete`, keyed by the parent space ID. A space site
  binds a `<label>.chatbotkit.space` subdomain to static content served from a
  space's storage.
- `alias` field on PartnerUser requests for instance identification.
- `BotID` field on the Bulletin struct for bot association.

### Changed

- Re-generated request/response types from the latest API spec, including the
  `alias` field now present across integration create/update requests.

### Fixed

- Re-pointed conversation message role and parameter type aliases after a
  generated type rotation.
