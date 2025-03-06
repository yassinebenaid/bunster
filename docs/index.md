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
  - title: Module System
    details: Bunsters aims to provide a module system makes it easy to share and consume scripts as libraries.
  - title: Bash compatible
    details: Bunsters aims to be compatible with bash. expecting that exising bash scripts do not have to be edited to work with bunster.
  - title: Environment Files
    details: We have a first class support for environment files. allowing you to load variables from .env files.
---
