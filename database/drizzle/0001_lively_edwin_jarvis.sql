CREATE TABLE IF NOT EXISTS "tags" (
	"id" serial PRIMARY KEY NOT NULL,
	"uuid" uuid DEFAULT gen_random_uuid(),
	"name" varchar(50) NOT NULL,
	"color" varchar(7) NOT NULL,
	"deleted_at" timestamp with time zone,
	"created_at" timestamp with time zone DEFAULT now(),
	"updated_at" timestamp with time zone DEFAULT now(),
	CONSTRAINT "tags_uuid_unique" UNIQUE("uuid")
);
--> statement-breakpoint
CREATE TABLE IF NOT EXISTS "tasks_to_tags" (
	"task_id" serial NOT NULL,
	"tag_id" serial NOT NULL,
	CONSTRAINT "tasks_to_tags_task_id_tag_id_pk" PRIMARY KEY("task_id","tag_id")
);
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "tasks_to_tags" ADD CONSTRAINT "tasks_to_tags_task_id_tasks_id_fk" FOREIGN KEY ("task_id") REFERENCES "public"."tasks"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
--> statement-breakpoint
DO $$ BEGIN
 ALTER TABLE "tasks_to_tags" ADD CONSTRAINT "tasks_to_tags_tag_id_tags_id_fk" FOREIGN KEY ("tag_id") REFERENCES "public"."tags"("id") ON DELETE no action ON UPDATE no action;
EXCEPTION
 WHEN duplicate_object THEN null;
END $$;
