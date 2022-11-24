import type { NextApiRequest, NextApiResponse } from "next";
import { unstable_getServerSession } from "next-auth";
import { ExampleApp } from "../../lib/main";
import { prisma } from "../../lib/prisma";
import { authOptions } from "./auth/[...nextauth]";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  const session = await unstable_getServerSession(req, res, authOptions);

  if (!session || !session?.user) {
    res.status(401).json({ error: "No session" });
    return;
  }

  const user = await prisma.user.findFirst({
    where: { id: session.user.id },
  });

  if (!user || !user.admin) {
    res.status(401).json({ error: "No permission" });
    return;
  }

  await ExampleApp.commands.syncCommands(false);

  res.status(200).json({ name: "Deployed Commands" });
}
