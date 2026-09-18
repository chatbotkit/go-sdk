# Changelog

All notable changes to the ChatBotKit Go SDK are documented in this file. The
format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and
this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

For releases prior to `0.2.0`, see the auto-generated notes on the
[GitHub Releases](https://github.com/chatbotkit/go-sdk/releases) page.

## [Unreleased]

## [0.5.0] - 2026-09-18

### Added

- `Options.Token` is the new name for the API credential. `Options.Secret` is
  deprecated and still works; `Token` wins when both are set.
- `client.Decision.Create` asks a decision model typed questions (boolean,
  choice, score) about a state and returns an answer with probabilities for
  each. The request is the hand-written `sdk.DecisionCreateRequest`, whose
  state, instructions and criteria take any JSON value, because the generated
  request type cannot hold the named options of a choice question.

### Fixed

- A path prefix on `Options.BaseURL` was dropped, so a platform served under a
  sub-path (for example `http://localhost:3000/cbk`) was unreachable. The prefix
  is now preserved for regular, raw and streaming requests. Plain `http` base
  URLs work for local use.
- The SDK did not build: `types/aliases.go` still pointed at generated names
  that the last types regeneration had renamed. The aliases are updated, and
  the names consumers use are unchanged.

### Changed

- **BREAKING:** skillset ability link fields renamed. On create/update/fetch/
  list/export, `secretId` / `fileId` / `botId` / `spaceId` are now
  `linkedSecretId` / `linkedFileId` / `linkedBotId` / `linkedSpaceId`. Inline
  conversation `extensions.skillsets[].abilities[]` entries use
  `linkedSecretId` (and the new `linkedSpaceId`). GraphQL `Ability` relations
  `secret` / `file` / `bot` / `space` are now `linkedSecret` / `linkedFile` /
  `linkedBot` / `linkedSpace`. There are no compatibility aliases; upgrade
  together with the platform deploy.
- Regenerated types also pick up unrelated API changes since the previous
  regeneration (2026-08-19), grouped by resource:
  - **BREAKING:** Dataset: `store` removed from `DatasetCreateRequest`,
    `DatasetFetchResponse`, `DatasetListResponseItem` and
    `DatasetListStreamItemData` (the platform now has a single vector store;
    the REST API accepts and ignores `store`).
  - Conversation: `expiresAt` (epoch ms, auto-delete) added to
    `ConversationUpdateRequest`, `ConversationFetchResponse`,
    `ConversationListResponseItem` and `ConversationListStreamItemData`.
  - Memory: `expiresAt` added to `MemoryCreateRequest`, `MemoryUpdateRequest`,
    `MemoryFetchResponse`, `MemoryListResponseItem` and
    `MemoryListStreamItemData`.
  - Task: `expiresAt` added to `TaskCreateRequest`, `TaskUpdateRequest`,
    `TaskFetchResponse`, `TaskListResponseItem` and `TaskListStreamItemData`;
    `resumeAt` (when a paused run resumes, null while running) added to
    `TaskExecutionListResponseItem` and `TaskExecutionListStreamItemData`.
  - Policy: `state` (`enabled` | `disabled`) added to `PolicyCreateRequest`,
    `PolicyUpdateRequest`, `PolicyFetchResponse`, `PolicyListResponseItem` and
    `PolicyListStreamItemData`.
  - WhatsApp integration: `appSecret` (Meta app secret for webhook signature
    validation, masked as `********` on read) added to
    `IntegrationWhatsAppCreateRequest`, `IntegrationWhatsAppUpdateRequest`,
    `IntegrationWhatsAppFetchResponse`, `IntegrationWhatsAppListResponseItem`
    and `IntegrationWhatsAppListStreamItemData`; `idempotencyKey` added to
    `WhatsappInitiateRequest`.
  - GitHub integration: `allowFrom` (allowed senders) added to
    `GithubIntegrationCreateRequest`.
- `agent.convertMessageExtensions` now forwards `LinkedSpaceID` alongside
  `LinkedSecretID` for inline abilities.

## [0.4.0] - 2026-06-27

### Added

- Secret token minting and request proxying. `SecretClient.Mint` /
  `ContactSecretClient.Mint` mint a usable token from a secret (`oauth`/`jwt`
  secrets only; owner-only) and return `{ Token, ExpiresAt }`.
  `SecretClient.Proxy` / `ContactSecretClient.Proxy` proxy a request through a
  secret — the credential is injected server-side (it never leaves the platform)
  and the upstream `*http.Response` is returned as-is, success or error (the
  caller closes `resp.Body`).
- `AuthorizationRequiredError` (aliased into the `sdk` package; detect with
  `errors.As`) carrying the `URL` the user must visit to authorize. It is
  returned when a secret or connection has not been authenticated yet
  (`409 authorization_required`) — by `Mint`, by any normal route, and by
  `Proxy` (which otherwise passes every genuine upstream response through
  untouched). Generic API `Error`s now also carry the HTTP `Status`.

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
