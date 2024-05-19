import { relations, sql } from "drizzle-orm";
import { pgTable, serial, timestamp, uuid, varchar } from "drizzle-orm/pg-core";

import { tasks } from "./tasks.schema";

export const categories = pgTable("categories", {
  id: serial("id").primaryKey(),
  uuid: uuid("uuid")
    .unique()
    .default(sql`gen_random_uuid()`),
  name: varchar("name", { length: 50 }).notNull(),
  deletedAt: timestamp("deleted_at", { withTimezone: true }),
  createdAt: timestamp("created_at", { withTimezone: true }).default(
    sql`now()`
  ),
  updatedAt: timestamp("updated_at", { withTimezone: true })
    .default(sql`now()`)
    .$onUpdate(() => new Date()),
});

export const categoriesRelations = relations(categories, ({ many }) => ({
  tasks: many(tasks),
}));
