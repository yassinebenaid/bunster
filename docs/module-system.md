# Module System

> **Status: 🚧 Planned for a future release — not yet available.**

The Module System is part of Bunster's long-term vision, as outlined in the [README](https://github.com/yassinebenaid/bunster#vision).

## What It Will Be

The Module System is intended to:

- **Allow you to publish and consume shell scripts as reusable libraries**, similar to package ecosystems like npm, PyPI, or crates.io.
- **Include a builtin package manager** integrated into Bunster for fetching, installing, and managing these library dependencies.
- **Make script composition easy**, so you can import shared functions and utilities from community-maintained packages.

## Current Status

This feature has not been implemented yet. The Bunster project is in its early stages (pre-v1.0), and development is currently focused on achieving full bash compatibility and stabilizing the core compiler.

## What You Can Do Today

While there is no module system yet, you can still organize reusable code manually:

- **Use `bunster generate`** to produce Go source code from your scripts. This creates a complete Go module directory that you can inspect or further process.
- **Organize shared functions in separate shell files** and use `source` (if your scripts are run in a real shell) or copy them into your Bunster project.
- **Watch the repository** for updates on the Module System by starring the [GitHub repo](https://github.com/yassinebenaid/bunster) or following the [issues](https://github.com/yassinebenaid/bunster/issues).

## Timeline

There is no confirmed release date for the Module System. Based on the project's roadmap, it is expected after the stable `v1.0.0` release, once full bash compatibility is achieved and the core feature set is complete.
