import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Relay",
  description: "docs for relay",
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
    ],

    sidebar: [
      {
        text: 'Quick Start',
        items: [
          { text: 'Installation', link: '/quick-start/installation' },
          { text: 'First Steps', link: '/quick-start/first-steps' },
        ]
      },
      {
        text: 'Make Relay your own',
        items: [
          { text: 'Theme', link: '/customization/theme' },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/yashranjan1/relay' }
    ]
  }
})
