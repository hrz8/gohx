import { relations, sql } from "drizzle-orm";
import {
  integer,
  pgEnum,
  pgTable,
  serial,
  timestamp,
  uuid,
  varchar,
} from "drizzle-orm/pg-core";

import { categories } from "./categories.schema";
import { tasksToTags } from "./tasks_to_tasks.schema";

export const taskStatus = pgEnum("task_status", [
  "open",
  "in_progress",
  "completed",
  "wont_do",
]);

export const tasks = pgTable("tasks", {
  id: serial("id").primaryKey(),
  uuid: uuid("uuid")
    .unique()
    .default(sql`gen_random_uuid()`),
  categoryId: serial("category_id")
    .references(() => categories.id, { onDelete: "cascade" })
    .notNull(),
  name: varchar("name", { length: 50 }).notNull(),
  description: varchar("description", { length: 200 }),
  order: integer("order").notNull(),
  status: taskStatus("status").notNull(),
  dueDate: timestamp("due_date", { withTimezone: true }),
  deletedAt: timestamp("deleted_at", { withTimezone: true }),
  createdAt: timestamp("created_at", { withTimezone: true }).default(
    sql`now()`
  ),
  updatedAt: timestamp("updated_at", { withTimezone: true }).default(
    sql`now()`
  ),
});

export const tasksRelations = relations(tasks, ({ one, many }) => ({
  category: one(categories, {
    fields: [tasks.categoryId],
    references: [categories.id],
  }),
  tasksToTags: many(tasksToTags),
}));
