# Golang Legacy Rescue: From Over-Engineered to Maintainable

This repository presents a simplified but realistic case of a Golang microservice that became unmaintainable due to excessive abstraction and misuse of design patterns.

It includes:
- `v1_java_style/`: an overengineered version using Java-like SOLID patterns.
- `v2_refactored/`: a cleaned-up, idiomatic, and testable version.
- `rescue-notes.md`: commentary on decisions and trade-offs.

## 🔧 Refactoring Progress

### ✅ Step 1 – Simplified UserService

- Replaced overengineered service layer with a clean, testable processor.
- Removed unnecessary factory and interface layers.
- Created a production-grade test file covering success, failure, and context scenarios.
- Folder: `v2_refactored/api/user_service.go`
- Tests: `v2_refactored/tests/user_service_test.go`

### ✅ Step 2 – Clean UserHandler with Observability

- Replaced tightly coupled handler with a modular HTTP handler using dependency injection.
- Introduced structured logging using slog for visibility and traceability.
- Added input validation and standardized JSON responses.
- Folder: `v2_refactored/api/user_handler.go`