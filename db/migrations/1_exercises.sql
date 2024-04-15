CREATE TABLE "total_exercise_stats" (
                                        "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                        "user_id" integer UNIQUE,
                                        "activity_id" integer,
                                        "session_name" varchar(255),
                                        "number_of_times" integer,
                                        "total_duration_hours" integer,
                                        "total_duration_minutes" integer,
                                        "total_duration_seconds" integer,
                                        "total_calories_burned" integer,
                                        "created_at" timestamp DEFAULT (now()),
                                        "updated_at" timestamp
);

CREATE TABLE "total_exercise_session" (
                                          "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                          "user_id" integer UNIQUE,
                                          "activity_id" integer,
                                          "total_duration_hours" integer,
                                          "total_duration_minutes" integer,
                                          "total_duration_seconds" integer,
                                          "total_calories_burned" integer,
                                          "created_at" timestamp DEFAULT (now()),
                                          "updated_at" timestamp
);

CREATE TABLE "exercise_session" (
                                    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                    "user_id" integer,
                                    "activity_id" integer,
                                    "session_name" varchar(255),
                                    "start_time" timestamp,
                                    "end_time" timestamp,
                                    "duration" float(8),
                                    calories_burned float(8),
                                    "created_at" timestamp DEFAULT (now())
);
