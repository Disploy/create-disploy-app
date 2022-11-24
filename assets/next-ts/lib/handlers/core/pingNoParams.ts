import type { ButtonHandler } from "disploy";

const PingNoParams: ButtonHandler = {
  customId: "ping",

  async run(interaction) {
    return void interaction.reply({
      content: `hello world!!!!!!!! (clicked by ${interaction.user})`,
    });
  },
};

export default PingNoParams;
