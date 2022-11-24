import { App } from "disploy";
import commands from "./commands/commands";
import handlers from "./handlers/handlers";

const clientId = process.env.DISCORD_CLIENT_ID;
const token = process.env.DISCORD_TOKEN;
const publicKey = process.env.DISCORD_PUBLIC_KEY;

if (!clientId || !token || !publicKey) {
  throw new Error("Missing environment variables");
}

export const ExampleApp = new App({
  logger: {
    debug: true,
  },
});

ExampleApp.start({
  clientId,
  token,
  publicKey,
  commands,
  handlers,
});
