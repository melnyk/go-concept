# Copilot Agent Onboarding Guide

This repository defines reusable, generic "concept" interfaces for Go (cache, key-value store) plus minimal "empty" no-op implementations used for prototyping, testing, or as documentation scaffolding. Your role: rapidly add new concept interfaces, keep code idiomatic, tested, and formatted. This guide orients you so you spend time coding—not rediscovering basics. Concrete implementations will be added in another modules or repositories.

## 0. Copilot Directives
- Scope: keep changes minimal and related to the task; do not refactor unrelated code.
- Go version: target Go 1.22; avoid new dependencies unless explicitly justified.
- Layout: new interface in `root/xyz.go`; empty impl in `xyz/empty.go`; optional concrete impl in `xyz/<impl>/`.
- Semantics:
  - Cache: `Get` returns `(zero value, false)` when missing/expired; zero TTL means no expiration; `Delete` accepts multiple keys.
  - KVStore: missing key returns `(zero value, ErrKeyNotFound)`.
- Quality: add/update table‑driven tests; assert interface satisfaction with `var _ concept.X[...] = (*T[...])(nil)`; run `go test ./...` locally.
- Interfaces: prefer small, behavior‑focused surfaces; do not add methods unless broadly useful.

## 1. Project Summary
- Purpose: Provide small, generic Go interfaces ("concepts") such as `Cache` and `KVStore` under module path `go.melnyk.org/concept` to standardize abstractions across projects.
- Current concepts: `Cache` (TTL-aware API surface) and `KVStore` (basic CRUD + enumeration). Errors: `ErrNotImplemented`, `ErrKeyNotFound`.
- Implementations: `cache/empty.Empty` and `kvstore/Empty` are deliberately inert (always miss / return sentinel errors) to allow dependency injection, testing, and gradual rollout.

## 2. Tech Stack
- Language: Go 1.25 (see `go.mod`). Keep backward compatibility with 1.20 unless explicitly bumped.
- Dependencies: None external currently. Adding deps raises maintenance cost—justify in PR description.
- Build/Test: Standard `go build ./...` and `go test ./...`. No custom scripts, no code generation.

## 3. Repository Structure
```
/ (module root)
  cache.go        -> Public `Cache` interface (generic)
  kvstore.go      -> Public `KVStore` interface (generic)
  errors.go       -> Shared sentinel errors
  examples_test.go -> (build tag "ignore") illustrative examples; not run by default
  cache/empty.go  -> `cache.Empty` no-op cache + test
  kvstore/kvstore.go -> `kvstore.Empty` no-op KV store + test
  README.md, CONTRIBUTING.md, CODE_OF_CONDUCT.md, LICENSE
  .github/ (this file)
```
No hidden build scripts, Makefiles, or CI config present (as of authoring). Assume plain Go toolchain.

## 4. Coding Guidelines
- Style: Follow standard Go conventions; run `gofmt` (or `go fmt ./...`). Maintain goimports ordering.
- Generics: Keep type parameters minimal (`[K comparable, V any]`). Avoid over-constraining unless needed.
- Interfaces: Small, behavior-focused. Do not add methods without clear cross-implementation value.
- Errors: Reuse existing sentinel errors when semantically correct. Introduce new sentinels in `errors.go` with a short comment. Prefer wrapping (`fmt.Errorf("...: %w", err)`) for contextual propagation.
- Documentation: Exported identifiers must have leading comments. Keep them succinct and action-oriented.
- Tests: Table-driven where multiple cases exist. For new implementations, cover: happy path, not-found, overwrite, delete semantics, TTL expiry (for caches), enumeration stop conditions.
- Panics: Avoid; prefer returning errors. Only panic for programmer misconfiguration (rare here).
- Dependencies: Favor standard library. If adding a third-party module, pin versions indirectly via Go's module proxy (normal) and explain rationale in PR.
- Performance: Only optimize after a demonstrable need; keep APIs simple.
- Concurrency: Implementations intended for production use should be safe for concurrent use by multiple goroutines. Call out if an implementation is not goroutine‑safe.

## 5. Common Tasks (Recipes)
Add new concept interface:
1. Create `xyz.go` at root with `package concept`.
2. Define minimal interface (aim < ~6 methods) with generics.
3. Add sentinel errors if necessary.
4. Update `README.md` (Current concepts list).
5. Provide at least one empty or reference implementation under a subpackage (`xyz/empty.go`).
6. Add tests for the implementation (`xyz/empty_test.go`).
7. Add interface satisfaction assertions where applicable.

Expose example usage:
- Use `_test.go` files with Example* functions (without the build tag) to integrate into `go test` docs.
 - Note: `examples_test.go` under the `ignore` build tag demonstrates shapes; its type parameters may not compile as-is (e.g., `any` does not satisfy `comparable` for `K`). Do not copy blindly.

## 6. Testing Expectations
- Command: `go test ./...` must pass cleanly (0 failures). Current suite tests only "empty" implementations.
- Avoid flaky time-based tests: abstract time via an interface when adding TTL logic.
- Keep example or integration tests fast (<1s total ideally).

## 7. Error & Edge Cases to Consider
- Cache: Missing key (should return zero value + false). Expired entries (if implemented) behave like missing.
- KVStore: Non-existent key returns `ErrKeyNotFound`. Save collisions (define overwrite semantics explicitly—default: overwrite silently) document in code comment.
- Enumeration: Must stop early if callback returns false; test that no further callbacks occur.

## 8. Documentation & Comments
- Keep README concise. For large additions, prefer package doc comment (create `doc.go` file) instead of bloating README.
- Update concept lists whenever adding/deprecating an interface.

## 9. Tooling & Commands (macOS/Linux)
- Build library: `go build ./...`
- Run tests: `go test ./...`
- Format: `go fmt ./...` (CI expectation per CONTRIBUTING)
- Lint (optional local): If you install `golangci-lint`, run `golangci-lint run` (not required currently). Do not assume its presence.

## 10. Contribution Workflow (Simplified)
1. Branch from `master`.
2. Implement change + tests.
3. Ensure: build passes, tests green, code formatted.
4. Commit message: short imperative summary first line, blank line, optional details.
5. Open PR referencing issue (if any). Indicate any new external dependencies explicitly.

## 11. Design Principles
- Minimal surface: Add only broadly useful abstractions.
- Composability: Interfaces first, empty implementations.
- Clarity over cleverness: Prefer explicit code to reflective or unsafe patterns.
- Test-first for observable behaviors.
 - Concurrency by default for shared components; document exceptions.

## 12. When Unsure
- Inspect existing `Empty` implementations for pattern consistency.
- Default to returning sentinel errors and zero values rather than panicking.
- Keep generics parameter order `[K comparable, V any]` for consistency.


## 14. Quick Checklist Before PR
- [ ] Interfaces documented
- [ ] Implementation asserts interface satisfaction
- [ ] Tests added & passing
- [ ] README updated (concept list) if new interface
- [ ] No stray build tags excluding essential examples
- [ ] No unused imports / unreachable code

---
This guide should remain < ~2 pages. Update it whenever adding concepts, implementations, build steps, or tooling.
