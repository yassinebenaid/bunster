---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "Bunster"
  text: "Compile shell scripts to native executable programs"
  tagline: A shell compiler that turns your scripts into self-contained executable programs
  image:
    src: /logo.png
    alt: Bunster
  actions:
    - theme: brand
      text: Installation
      link: /installation
    - theme: alt
      text: Documentation
      link: /introduction
features:
  - title: Static binaries
    details: |
      Scripts compiled with bunster are not just wrappers around your script, nor do they rely on any external shells on your system.
  - title: Bash compatible
    details: Bunsters aims to be compatible with bash. expecting that exising bash scripts do not have to be edited to work with bunster.
  - title: Module System
    details: You no longer have to write your scripts in a single file. Bunster allows you to distribute code across multiple files and directories thought of as single unit called module.
  - title: Environment Files
    details: .env files are nativily supported in bunster. Allowing you to load variables from .env files at runtime.
  - title: Static Asset Embedding
    details: You can embed files and directories within your compiled program. And use them as if they were normal files in the system at runtime.
  - title: Deferred commands
    details: Bunster allows you to defer the execution of a command to the end of the program or function. Useful for commands that perform clean up.
---
