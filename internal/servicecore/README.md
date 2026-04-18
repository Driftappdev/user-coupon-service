# Service Core Layer

This folder is service-specific application logic.
It should use `pkg/wrapper` for shared libraries and must avoid direct imports of `github.com/driftappdev/libpackage/*` from business handlers.
