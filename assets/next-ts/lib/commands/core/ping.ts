import type { Command } from "disploy";

const Ping: Command = {
  name: "ping",
  description: "pong!",

  async run(interaction) {
    return void interaction.reply({
      content: "hello world!!!!!!!!",
      components: [
        {
          type: 1,
          components: [
            {
              type: 2,
              label: "Click me!",
              style: 1,
              custom_id: `ping-${interaction.user.id}`,
            },
            {
              type: 2,
              label: "Anonymous mode",
              style: 2,
              emoji: {
                name: "🫢",
              },
              custom_id: `ping`,
            },
          ],
        },
      ],
    });
  },
};

export default Ping;
