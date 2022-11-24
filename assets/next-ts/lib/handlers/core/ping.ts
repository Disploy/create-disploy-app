import type { ButtonHandler } from "disploy";

const Ping: ButtonHandler = {
  customId: "ping-:userId",

  async run(interaction) {
    const originalUser = await interaction.params.getUserParam("userId");
    const clicker = interaction.user;

    return void interaction.reply({
      content: `hello world!!!!!!!! (clicked by ${clicker}) [made by ${originalUser}]`,
    });
  },
};

export default Ping;
