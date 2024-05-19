import { pgTable, primaryKey, serial } from "drizzle-orm/pg-core";

import { tasks } from "./tasks.schema";
import { tags } from "./tags.schema";

export const tasksToTags = pgTable(
  "tasks_to_tags",
  {
    taskId: serial("task_id")
      .notNull()
      .references(() => tasks.id),
    tagId: serial("tag_id")
      .notNull()
      .references(() => tags.id),
  },
  (t) => ({
    pk: primaryKey({ columns: [t.taskId, t.tagId] }),
  })
);
