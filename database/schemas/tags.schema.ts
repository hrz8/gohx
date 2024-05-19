import { relations, sql } from "drizzle-orm";
import { pgTable, serial, timestamp, uuid, varchar } from "drizzle-orm/pg-core";

import { tasksToTags } from "./tasks_to_tasks.schema";

export const tags = pgTable("tags", {
  id: serial("id").primaryKey(),
  uuid: uuid("uuid")
    .unique()
    .default(sql`gen_random_uuid()`),
  name: varchar("name", { length: 50 }).notNull(),
  color: varchar("color", { length: 7 }).notNull(),
  deletedAt: timestamp("deleted_at", { withTimezone: true }),
  createdAt: timestamp("created_at", { withTimezone: true }).default(
    sql`now()`
  ),
  updatedAt: timestamp("updated_at", { withTimezone: true })
    .default(sql`now()`)
    .$onUpdate(() => new Date()),
});

export const tagsRelations = relations(tags, ({ many }) => ({
  tasksToTags: many(tasksToTags),
}));
