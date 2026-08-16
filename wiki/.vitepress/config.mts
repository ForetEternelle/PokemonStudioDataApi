import { defineConfig } from 'vitepress'

export default defineConfig({
  title: "PokemonStudioDataApi",
  description: "The official documentation of the Pokemon Studio data api",
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Installation', link: '/installation' },
      { text: 'Development', link: '/dev/setup' },
      { text: 'API', link: '/api' },
    ],

    sidebar: [
      {
        text: 'Guide',
        items: [
          { text: 'Installation', link: '/installation' },
          { text: 'Configuration', link: '/configuration' },
        ],
      },
      {
        text: 'Development',
        items: [
          { text: 'Setup', link: '/dev/setup' },
          { text: 'Project Structure', link: '/dev/structure' },
          { text: 'Code Documentation', link: '/dev/code' },
        ],
      },
      {
        text: 'API Reference',
        items: [
          { text: 'Swagger UI', link: '/api' },
        ],
      },
    ],

    socialLinks: [
    ]
  }
})
