import { defineConfig } from "vitepress";

export default defineConfig({
	title: "Bunster",
	description:
		"Compile shell scripts to native self-contained executable programs",
	themeConfig: {
		nav: [
			{ text: "Documentation", link: "/introduction" },
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
					{ text: "Introduction", link: "/introduction" },
					{ text: "Installation", link: "/installation" },
					{ text: "Supported features", link: "/supported-features" },
					{ text: "CLI", link: "/cli" },
					{
						text: "Features",
						collapsed: false,
						items: [
							{ text: "Simple commands", link: "/features/simple-commands" },
							{ text: "Redirections", link: "/features/redirections" },
							{ text: "Pipelines", link: "/features/pipelines" },
							{ text: "Lists", link: "/features/lists" },
							{
								text: "Asynchronous commands",
								link: "/features/async-commands",
							},
							{
								text: "Variables & Environment",
								link: "/features/variables-and-environment",
							},
							{
								text: "Environment files",
								link: "/features/environment-files",
							},
							{
								text: "Groups And Sub-Shells",
								link: "/features/groups-and-subshells",
							},
							{
								text: "Loops",
								link: "/features/loops",
							},
							{
								text: "Deferred Commands",
								link: "/features/deferred-commands",
							},
							{
								text: "Test commands",
								link: "/features/test-commands",
							},
							{
								text: "Arithmetics",
								link: "/features/arithmetics",
							},
							{
								text: "Embedding",
								link: "/features/embedding",
							},
							{ text: "Builtin commands", link: "/features/builtins" },
						],
					},
					{ text: "Contributing", link: "/contributing" },
					{ text: "Developers Guideline", link: "/developers" },
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
				""
			);

		pageData.frontmatter.head ??= [];
		pageData.frontmatter.head.push([
			"link",
			{ rel: "canonical", href: canonicalUrl },
		]);
	},
});
