import { defineConfig } from "vitepress";

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Bunster",
  description:
    "Compile shell scripts to native self-contained executable programs",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: "Documentation", link: "/quick-start" },
      { text: "Installation", link: "/installation" },
    ],
    logo: "/logo.png",
    search: {
      provider: "local",
    },
    sidebar: [
      {
        items: [
          { text: "Quick Start", link: "/quick-start" },
          { text: "Installation", link: "/installation" },
          { text: "Supported Features", link: "/supported-features" },
        ],
      },
    ],

    socialLinks: [
      { icon: "github", link: "https://github.com/yassinebenaid/bunster" },
    ],

    footer: {
      message: "Released under the GPLv3 License.",
      copyright: "Copyright © 2024-present Yassine Benaid",
    },
    editLink: {
      pattern:
        "https://github.com/yassinebenaid/bunster/edit/master/docs/:path",
    },
  },
  head: [
    ["link", { rel: "manifest", href: "/site.webmanifest" }],
    ["link", { rel: "icon", type: "image/x-icon", href: "/favicon.ico" }],
  ],
});
