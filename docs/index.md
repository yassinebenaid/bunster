---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "Bunster"
  text: "Compile shell scripts to native executable programs"
  tagline: A shell compiler that turns your scripts into a self-contained executable programs
  image:
    src: /logo.png
    alt: Bunster
  actions:
    - theme: brand
      text: Installation
      link: /installation
    - theme: alt
      text: Documentation
      link: /quick-start
features:
  - title: Native Program Generation
    details: |
      Scripts compiled with Bunster are not just wrappers around your script, nor do they rely on any external shell on your system.
  - title: Wite Once, Run Everywhere
    details: Bunster offers static linking. Your scripts are compiled to a statically linked binary that runs on every machine.
  - title: Module System
    details: Bunsters's module system makes it easy to shares your scripts as a versioned modules to be used by others.
---
