# Changelog

## v1.0.9

- build(deps): bump alpine from 3.23 to 3.24
- build(deps): bump github.com/steadybit/action-kit/go/action_kit_sdk
- build(deps): bump github.com/steadybit/advice-kit/go/advice_kit_sdk
- build(deps): bump github.com/steadybit/discovery-kit/go/discovery_kit_sdk
- build(deps): bump github.com/steadybit/event-kit/go/event_kit_api
- build(deps): bump github.com/steadybit/extension-kit
- build(deps): bump github.com/steadybit/preflight-kit/go/preflight_kit_sdk/v2
- chore(deps): bump go to 1.26.5 (#175)
- chore(deps): bump go-openapi/swag/loading to fix go mod tidy (#177)
- chore(deps): bump golang.org/x/net to v0.55.0 (CVE-2026-39821) (#164)
- chore: add Claude Code workflows (#169)
- chore: silence SonarQube finding on secrets: inherit in Claude workflows
- ci: skip build on .trivyignore.yml-only changes [skip ci]
- feat: drop "Monday" from workflow name and add release_type input
- refactor: register extension index via exthttp.RegisterRevisionedHandler (#176)

## v1.0.8

- Update dependencies

## v1.0.7

- Support discovery group attribute via `STEADYBIT_EXTENSION_DISCOVERY_GROUP` env var (or `discovery.group` Helm value) — when set, the extension adds `steadybit.group=<value>` to every discovered target
- Update dependencies

## v1.0.6

- Bump Go to 1.26.3
- Update dependencies

## v1.0.5

- Bump Go to 1.25.9
- Support if-none-match for the extension list endpoint
- Update dependencies

## v1.0.4

- feat(chart): split image.name into image.registry + image.name
- Support global.priorityClassName
- Update alpine packages in Docker image to address CVEs
- Update dependencies

## v1.0.3

- Update dependencies

## v1.0.2

- Update dependencies

## v1.0.1

- Update dependencies

## v1.0.0

 - Initial release
