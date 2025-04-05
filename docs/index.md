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
    - theme: alt
      text: Github
      link: https://github.com/yassinebenaid/bunster
features:
  - title: Static binaries
    icon:
      src: ./assets/landing/binary.svg
    details: |
      Scripts compiled with bunster are not just wrappers around your script, nor do they rely on any external shells on your system.
  - title: Bash compatible
    icon:
      src: ./assets/landing/bash.svg
    details: Bunsters aims to be compatible with bash. expecting that exising bash scripts do not have to be edited to work with bunster.
  - title: Module System
    details: You are not limited to write all of your code in a single file. Your code can be distributed across as many files as needed.
  - title: Package Manager
    details: Bunster has a buitlin package manager that makes it easy to publish and consume modules as libraries.
  - title: Environment Files
    icon:
      src: ./assets/landing/dotenv.png
    details: .env files are nativily supported in bunster. Allowing you to load variables from .env files at runtime.
  - title: Static Asset Embedding
    icon:
      src: ./assets/landing/static-assets.png
    details: You can embed files and directories within your compiled program. And use them as if they were normal files in the system at runtime.
---
