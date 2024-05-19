CREATE TYPE "public"."task_status" AS ENUM('open', 'in_progress', 'completed', 'wont_do');
ALTER TABLE "tasks" ADD CONSTRAINT "tasks_category_id_categories_id_fk" FOREIGN KEY ("category_id") REFERENCES "public"."categories"("id") ON DELETE cascade ON UPDATE no action;
ALTER TABLE "tasks_to_tags" ADD CONSTRAINT "tasks_to_tags_task_id_tasks_id_fk" FOREIGN KEY ("task_id") REFERENCES "public"."tasks"("id") ON DELETE no action ON UPDATE no action;
ALTER TABLE "tasks_to_tags" ADD CONSTRAINT "tasks_to_tags_tag_id_tags_id_fk" FOREIGN KEY ("tag_id") REFERENCES "public"."tags"("id") ON DELETE no action ON UPDATE no action;
