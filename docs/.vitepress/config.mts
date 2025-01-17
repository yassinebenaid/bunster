import { defineConfig } from "vitepress";

export default defineConfig({
  title: "Bunster",
  description:
    "Compile shell scripts to native self-contained executable programs",
  themeConfig: {
    nav: [
      { text: "Documentation", link: "/quick-start" },
      { text: "Installation", link: "/installation" },
      { text: "Maintainers", link: "/maintainers" },
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
          { text: "Usage", link: "/usage" },
          { text: "Contributing", link: "/contributing" },
        ],
      },
    ],

    socialLinks: [
      { icon: "github", link: "https://github.com/yassinebenaid/bunster" },
    ],

    footer: {
      message: "Released under the 3-Clause BSD License.",
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
  lastUpdated: true,
  sitemap: {
    hostname: "https://bunster.netlify.app",
  },
  cleanUrls: true,

  // we set the canonical url
  transformPageData(pageData) {
    const canonicalUrl =
      `https://bunster.netlify.app/${pageData.relativePath}`.replace(
        /\.md$/,
        "",
      );

    pageData.frontmatter.head ??= [];
    pageData.frontmatter.head.push([
      "link",
      { rel: "canonical", href: canonicalUrl },
    ]);
  },
});
