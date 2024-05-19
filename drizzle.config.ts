import dotenv from "dotenv";
dotenv.config();

import type { Config } from "drizzle-kit";

const DATABASE_URL = process.env.DATABASE_URL;

if (!DATABASE_URL) {
  throw new Error("cannot resolve database connection!");
}

export default {
  migrations: {
    table: "migrations",
    schema: "public",
  },
  schema: "./database/drizzle/*.schema.ts",
  out: "./database/schema",
  dialect: "postgresql",
  dbCredentials: {
    url: DATABASE_URL,
  },
} satisfies Config;
