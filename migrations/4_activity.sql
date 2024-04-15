CREATE TABLE "activity" (
                            "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                            "user_id" integer,
                            "name" varchar(255),
                            "calories_per_hour" float(8),
                            "duration_minutes" float(8),
                            "total_calories" float(8),
                            "created_at" timestamp DEFAULT (now()),
                            "updated_at" timestamp DEFAULT null
);

CREATE TABLE "activity_user" (
                                 "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                 "user_id" integer UNIQUE,
                                 "activity_id" integer UNIQUE,
                                 "created_at" timestamp DEFAULT (now()),
                                 "updated_at" timestamp DEFAULT null
);
