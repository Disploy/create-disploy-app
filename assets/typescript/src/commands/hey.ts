import type { Command } from "disploy";

const HeyCommand: Command = {
  name: "hey",
  description: "heyy!",

  async run(interaction) {
    interaction.deferReply();

    await new Promise((resolve) => setTimeout(resolve, 5000));

    return void interaction.editReply({
      content: `Hello! (with artificial delay)`,
    });
  },
};

export default HeyCommand;
