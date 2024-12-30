---
sidebar: false
---

<script setup>
import { VPTeamMembers } from 'vitepress/theme'

const members = [
  {
    avatar: 'https://www.github.com/yassinebenaid.png',
    name: 'Yassine Benaid',
    title: 'Creator',
    links: [
      { icon: 'github', link: 'https://github.com/yassinebenaid' },
      { icon: 'linkedin', link: 'https://www.linkedin.com/in/yassinebenaid' },
      { icon: 'reddit', link: 'https://www.reddit.com/user/yassinebenaid' }
    ]
  },
]
</script>

# Maintainers
**Bunster** is an open source project driven by the public community, But we have a dedicated team
that oversees the final decisions to ensure the project's direction aligns with its vision. The team also maintains open communication with the community
to sustain a welcoming and collaborative environment for everyone.

Say hello to our awesome team.

<VPTeamMembers size="small" :members="members" />
