<div align="center">
	<br />
	<p>
		<a href="https://disploy.dev"><img src="https://disploy.dev/img/logo.svg" alt="disploy" width="200" /></a>
	</p>
    <p>
		<a href="https://vercel.com/?utm_source=disploy&utm_campaign=oss"><img src="https://www.datocms-assets.com/31049/1618983297-powered-by-vercel.svg" alt="Vercel" /></a>
	</p>
    <h3>
        Next.js (TypeScript)
    </h3>
	<br />
	<p>
		<a href="https://discord.gg/E3z8MDnTWn"><img src="https://img.shields.io/discord/901426442242498650?color=5865F2&logo=discord&logoColor=white" alt="Disploy's Discord server" /></a>
		<a href="https://github.com/disploy/Disploy/actions"><img src="https://github.com/Disploy/disploy/actions/workflows/tests.yml/badge.svg" alt="Tests status" /></a>
	</p>

</div>

## Overview

This template uses Next.js as your server with Disploy, this makes it easy to integrate a website that interfaces with your bot.

## Workflow

- Migrate database

```sh
$ yarn push-db
# or
$ npm run push-db
# or
$ pnpm run push-db
```

- Use Prisma Studio to make yourself an admin

```sh
$ yarn prisma studio
# or
$ npx prisma studio
# or
$ pnpm prisma studio
```

- Deploy commands

Visit http://localhost:3000/admin and click "Deploy Commands"

- Start

```sh
$ yarn dev
# or
$ npm run dev
# or
$ pnpm run dev
```

- Tunnel

```sh
# ngrok is a great way to reverse proxy to the public internet
$ ngrok http 3000
```

Make sure to set "INTERACTIONS ENDPOINT URL" in your application to the ngrok url appended with `/api/interactions`

## Need Help?

https://discord.gg/E3z8MDnTWn - Join our Discord server for support and updates!
