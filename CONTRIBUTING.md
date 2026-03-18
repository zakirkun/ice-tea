# Contributing to Ice Tea

Thank you for your interest in contributing to the Ice Tea Security Scanner! 🍵

This project relies on the community to build comprehensive and cutting-edge security `SKILLs` and improve internal core logic.

## 1. Getting Started

- Ensure you have **Go 1.21+** installed.
- Ensure you have a C compiler configured on your machine (`gcc` or `clang` via MSYS2, Xcode, etc.), which is required by `go-tree-sitter`.
- Create a fork of this repository.

## 2. Issues & Reporting

- **Bug Reports**: Open an issue detailing the error, how to reproduce it, the OS, and the Go version.
- **Feature Requests**: We love new ideas! Make sure to provide a strong use-case argument in your request.
- **False Positives**: If the pattern matching engine identifies safe code as vulnerable, please submit an issue or open a PR fixing the rule! Alternatively, verify if the LLM validation flag (`--enable-llm`) correctly filtered out the FP.

## 3. Contributing Custom SKILLs

Currently, the most important contribution is expanding our `skills/` directory! Ice Tea uses a language-agnostic detection schema.

1. Ensure the vulnerability you wish to add doesn't already exist.
2. Read the [How to Create SKILLs](docs/08-how-to-create-skills.md) guide thoroughly.
3. Generate the `SKILL.md` file (making sure to use proper markdown formatting and frontmatter metadata) and `patterns.yaml` holding your test logic.
4. Add a test case demonstrating the vulnerability to `testdata/vulnerable/<lang>/`.
5. Create a PR. 

## 4. Development Workflow

If you're looking to modify the Go logic:
1. `make build` compiles your changes to `bin/ice-tea`.
2. run `go test ./...` frequently. 
3. Code changes should include corresponding unit tests to maintain coverage.

## 5. Code of Conduct

Please engage respectfully with maintainers and other contributors. We hold ourselves to a standard of professional and patient communication.
