import {
  Block,
  Button,
  Card,
  ColGrid,
  Tab,
  TabList,
  Text,
  Title,
} from "@tremor/react";
import axios from "axios";
import { GetServerSidePropsContext } from "next";
import { unstable_getServerSession } from "next-auth";
import { useState } from "react";
import { prisma } from "../lib/prisma";
import { authOptions } from "./api/auth/[...nextauth]";

function CommandDeployer() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  async function handleDeploy() {
    setLoading(true);
    setError(null);
    axios
      .get("/api/deploy")
      .then((res) => setLoading(false))
      .catch(() => setError("Something went wrong"));
  }

  return (
    <Button
      text={loading ? "loading" : "Deploy Globally"}
      disabled={loading}
      handleClick={handleDeploy}
    />
  );
}

export async function getServerSideProps(context: GetServerSidePropsContext) {
  const session = await unstable_getServerSession(
    context.req,
    context.res,
    authOptions
  );

  if (!session) {
    return {
      redirect: {
        destination: "/",
        permanent: false,
      },
    };
  }

  const user = await prisma.user.findFirst({
    where: {
      email: session.user?.email,
    },
  });

  if (!user?.admin) {
    return {
      redirect: {
        destination: "/",
        permanent: false,
      },
    };
  }

  return {
    props: {
      session,
    },
  };
}

export default function Admin() {
  const [selectedView, setSelectedView] = useState(1);
  return (
    <>
      <main>
        <Title>Admin Dashboard</Title>
        <TabList
          defaultValue={1}
          handleSelect={(value) => setSelectedView(value)}
          marginTop="mt-6"
        >
          <Tab value={1} text="Stats" />
          <Tab value={2} text="Tools" />
        </TabList>

        {selectedView === 1 ? (
          <>
            <ColGrid
              numColsMd={2}
              numColsLg={3}
              gapX="gap-x-6"
              gapY="gap-y-6"
              marginTop="mt-6"
            >
              <Card>
                {/* Placeholder to set height */}
                <div className="h-28" />
              </Card>
            </ColGrid>
          </>
        ) : (
          <Block marginTop="mt-6">
            <Card>
              <Text>Deploy Application Commands</Text>
              <CommandDeployer />
            </Card>
          </Block>
        )}
      </main>
    </>
  );
}
