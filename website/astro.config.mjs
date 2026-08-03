import { defineConfig } from 'astro/config';
import sitemap from '@astrojs/sitemap';
import starlight from '@astrojs/starlight';

export default defineConfig({
  site: 'https://mrqwenty.github.io',
  base: '/fast-context-protocol',
  integrations: [
    starlight({
      title: 'CGP',
      description: 'Context Governance Protocol — governed context infrastructure for AI systems.',
      favicon: '/fast-context-protocol/favicon.svg',
      logo: {
        src: './src/assets/cgp-mark.svg',
        alt: 'CGP',
        replacesTitle: false,
      },
      social: {
        github: 'https://github.com/MrQwenty/fast-context-protocol',
      },
      editLink: {
        baseUrl: 'https://github.com/MrQwenty/fast-context-protocol/edit/master/website/src/content/docs/',
      },
      customCss: ['./src/styles/custom.css'],
      sidebar: [
        {
          label: 'Start here',
          items: [
            { label: 'Introduction', link: '/' },
            { label: 'Quick start', slug: 'getting-started/quick-start' },
            { label: 'Project status', slug: 'project/status' },
          ],
        },
        {
          label: 'Core concepts',
          items: [
            { label: 'Context Contract', slug: 'concepts/context-contract' },
            { label: 'Context Graph', slug: 'concepts/context-graph' },
            { label: 'Receipts', slug: 'reference/receipts' },
          ],
        },
        {
          label: 'Architecture',
          items: [
            { label: 'System overview', slug: 'architecture/overview' },
            { label: 'Request lifecycle', slug: 'architecture/request-lifecycle' },
          ],
        },
        {
          label: 'Privacy & security',
          items: [
            { label: 'Local Privacy Gateway', slug: 'privacy/local-gateway' },
            { label: 'Anonymization model', slug: 'privacy/anonymization' },
            { label: 'Context Trust Firewall', slug: 'security/trust-firewall' },
          ],
        },
        {
          label: 'Governance',
          items: [
            { label: 'EU governance direction', slug: 'governance/eu-compliance' },
            { label: 'Open-source strategy', slug: 'project/open-source' },
          ],
        },
        {
          label: 'Integrations',
          items: [
            { label: 'CGP and MCP', slug: 'integrations/mcp' },
          ],
        },
        {
          label: 'Reference',
          items: [
            { label: 'HTTP API', slug: 'reference/http-api' },
            { label: 'Repository layout', slug: 'reference/repository-layout' },
          ],
        },
        {
          label: 'Project',
          items: [
            { label: 'Roadmap', slug: 'project/roadmap' },
            { label: 'Naming and compatibility', slug: 'project/naming' },
          ],
        },
      ],
    }),
    sitemap(),
  ],
});
