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
**Bunster** is an open source project maintained and driven by the community, But we have a dedicated team
that takes the final decision on contributions and feature requests. And keeps contact with the rest of the community to ensure a healthy and friendly
environment for everyone.

Say hello to our awesome team.

<VPTeamMembers size="small" :members="members" />
