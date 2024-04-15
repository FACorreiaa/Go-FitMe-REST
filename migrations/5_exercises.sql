CREATE TABLE IF NOT EXISTS user_exercises (
                                              id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                              user_id integer NOT NULL,
                                              exercise_id UUID NOT NULL,
                                              created_at timestamp DEFAULT (now()),
                                              FOREIGN KEY (user_id) REFERENCES users(id),
                                              FOREIGN KEY (exercise_id) REFERENCES exercise_list(id)
);

CREATE TABLE IF NOT EXISTS "exercise_list" (
                                               "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                               "name" varchar(255),
                                               "type" varchar(255),
                                               "muscle" varchar(255),
                                               "equipment" varchar(255),
                                               "difficulty" varchar(255),
                                               "instructions" text,
                                               "video" varchar(255),
                                               "custom_created" boolean DEFAULT true,
                                               "created_at" timestamp DEFAULT (now()),
                                               "updated_at" timestamp DEFAULT null
);
