CREATE TYPE "user_role" AS ENUM ('user', '_admin');
CREATE TYPE "user_gender" AS ENUM ('male', 'female');

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

-- start_time, end_time, duration, calories_burned, created_at
CREATE TABLE "user_personal_data" (
  "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "user_id" integer,
  "firstname" varchar(255),
  "lastname" varchar(255),
  "gender" user_gender DEFAULT 'male',
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);
CREATE TABLE "user_bio_data" (
  "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "user_id" integer,
  "weight" float(8),
  "height" float(8),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);
CREATE TABLE "account" (
  "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "user_id" integer,
  "type" varchar(255),
  "provider" varchar(255),
  "providerAccountId" varchar(255),
  "refresh_token" varchar(255),
  "access_token" varchar(255),
  "expires_at" integer,
  "token_type" varchar(255),
  "scope" varchar(255),
  "id_token" varchar(255),
  "session_state" varchar(255),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);
CREATE TABLE "users" (
                         "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) UNIQUE NOT NULL,
  "role" user_role NOT NULL DEFAULT 'user',
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);


CREATE TABLE IF NOT EXISTS "weight_measure" (
                                  "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                  "user_id" integer,
                                  "weight_value" decimal(10,2),
                                  "created_at" timestamp DEFAULT (now()),
                                  "updated_at" timestamp DEFAULT null
);

CREATE TABLE IF NOT EXISTS "water_intake" (
  "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "user_id" integer,
  "quantity" decimal(10,2),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);

CREATE TABLE IF NOT EXISTS "waist_line" (
                                "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                "user_id" integer,
                                "quantity" decimal(10,2),
                                "created_at" timestamp DEFAULT (now()),
                                "updated_at" timestamp DEFAULT null
);

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

CREATE TABLE IF NOT EXISTS user_exercises (
                                              id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                              user_id integer NOT NULL,
                                              exercise_id UUID NOT NULL,
                                              created_at timestamp DEFAULT (now()),
                                              FOREIGN KEY (user_id) REFERENCES users(id),
                                              FOREIGN KEY (exercise_id) REFERENCES exercise_list(id)
);

select * from exercise_list where id = '0268b875-9602-4a06-8738-7b38006882e8';
select * from exercise_list;
select * from workout_plan_detail;

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

select * from workout_plan;
select * from workout_day;
select * from workout_plan_detail;
select * from exercise_list;

SELECT
    wd.day AS workout_day,
    wp.description AS workout_description,
    el.name AS exercise_name,
    el.type AS exercise_type,
    el.muscle AS exercise_muscle,
    el.equipment AS exercise_equipment,
    el.difficulty AS exercise_difficulty,
    el.instructions AS exercise_instructions,
    el.video AS exercise_video
FROM workout_plan_detail wpd
         JOIN workout_plan wp ON wpd.workout_plan_id = wp.id
         JOIN exercise_list el ON el.id = ANY(wpd.exercises)
         JOIN workout_day wd ON wd.workout_plan_id = '07ef66c2-2d92-4e33-9c83-4c3984c0dc15'
        AND wp.user_id = 40;

CREATE TABLE IF NOT EXISTS "workout_plan" (
                                              "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                              "user_id" integer,
                                              "description" varchar(255),
                                              "notes" varchar(255),
                                              "rating" integer DEFAULT 10,
                                              "created_at" timestamp DEFAULT (now()),
                                              "updated_at" timestamp DEFAULT null,
                                              FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS "workout_plan_detail" (
                                                     "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                                     "workout_plan_id" UUID,
                                                     "day" varchar(100),
                                                     "exercises" uuid[],
                                                     "created_at" timestamp DEFAULT (now()),
                                                     FOREIGN KEY (workout_plan_id) REFERENCES workout_plan(id)

);




CREATE TABLE IF NOT EXISTS "workout_day" (
                                             "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                             "workout_plan_id" UUID,
                                             "day" varchar(255),
                                             "created_at" timestamp DEFAULT (now()),
                                             "updated_at" timestamp DEFAULT null,
                                             FOREIGN KEY (workout_plan_id) REFERENCES workout_plan(id)
);



CREATE TABLE "food" (
                        "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "name" varchar(255),
  "calories" float(8),
  "serving_size" float(8),
  "protein" float(8),
  "fat_total" float(8),
  "fat_saturated" float(8),
  "carbohydrates_total" float(8),
  "fiber" float(8),
  "sugar" float(8),
  "sodium" float(8),
  "potassium" float(8),
  "cholesterol" float(8),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);
CREATE TABLE "meal_type" (
                             "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "user_id" integer UNIQUE,
  "food_id" integer UNIQUE,
  "meal_number" integer,
  "meal_description" varchar(255),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);
CREATE TABLE "meal_plan" (
                             "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "user_id" integer UNIQUE,
  "meal_type_id" integer UNIQUE,
  "description" varchar(255),
  "notes" varchar(255),
  "total_calories" float(8),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null,
  "rating" integer DEFAULT 10
);
CREATE TABLE "favourite" (
                             "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "user_id" integer UNIQUE,
  "exercise_id" integer UNIQUE,
  "activity_id" integer UNIQUE,
  "food_id" integer UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);
CREATE TABLE "recipe" (
                          "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "user_id" integer UNIQUE,
  "food_id" integer UNIQUE,
  "description" varchar(255),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);
CREATE TABLE "recipe_user" (
                               "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "recipe_id" integer UNIQUE,
  "user_id" integer UNIQUE,
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
CREATE TABLE "meal_plan_meal_type" (
                                       "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "meal_plan_id" integer UNIQUE,
  "meal_type_id" integer UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);
CREATE TABLE "meal_plan_user" (
                                  "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "meal_plan_id" integer UNIQUE,
  "user_id" integer UNIQUE,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT null
);

CREATE TABLE "water_intake_user" (
                                     "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  "water_intake_id" integer UNIQUE,
  "user_id" integer UNIQUE
);

CREATE TABLE IF NOT EXISTS user_macro_distribution (
                                                       id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                                       user_id INTEGER NOT NULL,
                                                       age BIGINT NOT NULL,
                                                       height DOUBLE PRECISION NOT NULL,
                                                       weight DOUBLE PRECISION NOT NULL,
                                                       gender VARCHAR NOT NULL,
                                                       system VARCHAR NOT NULL,
                                                       activity VARCHAR NOT NULL,
                                                       activity_description VARCHAR NOT NULL,
                                                       objective VARCHAR NOT NULL,
                                                       objective_description VARCHAR NOT NULL,
                                                       calories_distribution VARCHAR NOT NULL,
                                                       calories_distribution_description VARCHAR NOT NULL,
                                                       protein INTEGER NOT NULL,
                                                       fats INTEGER NOT NULL,
                                                       carbs INTEGER NOT NULL,
                                                       bmr INTEGER NOT NULL,
                                                       tdee INTEGER NOT NULL,
                                                       goal INTEGER NOT NULL,
                                                       created_at TIMESTAMP NOT NULL
);


COMMENT ON COLUMN "food"."serving_size" IS 'grams';
COMMENT ON COLUMN "food"."protein" IS 'grams';
COMMENT ON COLUMN "food"."fat_total" IS 'grams';
COMMENT ON COLUMN "food"."fat_saturated" IS 'grams';
COMMENT ON COLUMN "food"."carbohydrates_total" IS 'grams';
COMMENT ON COLUMN "food"."fiber" IS 'grams';
COMMENT ON COLUMN "food"."sugar" IS 'grams';
COMMENT ON COLUMN "food"."sodium" IS 'miligrams';
COMMENT ON COLUMN "food"."potassium" IS 'miligrams';
COMMENT ON COLUMN "food"."cholesterol" IS 'miligrams';
-- ALTER TABLE "user_personal_data"
-- ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
-- ALTER TABLE "user_bio_data"
-- ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
-- ALTER TABLE "account"
-- ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
-- ALTER TABLE "user"
-- ADD FOREIGN KEY ("id") REFERENCES "water_intake_user" ("user_id");
-- ALTER TABLE "water_intake"
-- ADD FOREIGN KEY ("id") REFERENCES "water_intake_user" ("water_intake_id");
-- ALTER TABLE "user"
-- ADD FOREIGN KEY ("id") REFERENCES "workout_plan_user" ("user_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "workout_plan"
-- ADD FOREIGN KEY ("id") REFERENCES "workout_plan_user" ("workout_plan_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "workout_plan"
-- ADD FOREIGN KEY ("id") REFERENCES "workout_day_plan" ("workout_plan_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "workout_day"
-- ADD FOREIGN KEY ("id") REFERENCES "workout_day_plan" ("workout_day_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "exercise"
-- ADD FOREIGN KEY ("id") REFERENCES "workout_day_exercise" ("exercise_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "workout_day"
-- ADD FOREIGN KEY ("id") REFERENCES "workout_day_exercise" ("workout_day_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "meal_plan"
-- ADD FOREIGN KEY ("id") REFERENCES "meal_plan_user" ("meal_plan_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "user"
-- ADD FOREIGN KEY ("id") REFERENCES "meal_plan_user" ("user_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "meal_type"
-- ADD FOREIGN KEY ("id") REFERENCES "meal_plan_meal_type" ("meal_type_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "meal_plan"
-- ADD FOREIGN KEY ("id") REFERENCES "meal_plan_meal_type" ("meal_plan_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "meal_type"
-- ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "food"
-- ADD FOREIGN KEY ("id") REFERENCES "meal_type" ("food_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "activity"
-- ADD FOREIGN KEY ("user_id") REFERENCES "activity_user" ("activity_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "user"
-- ADD FOREIGN KEY ("id") REFERENCES "activity_user" ("user_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "favourite"
-- ADD FOREIGN KEY ("food_id") REFERENCES "food" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "favourite"
-- ADD FOREIGN KEY ("activity_id") REFERENCES "activity" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "favourite"
-- ADD FOREIGN KEY ("exercise_id") REFERENCES "exercise" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "favourite"
-- ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "user"
-- ADD FOREIGN KEY ("id") REFERENCES "recipe" ("food_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "recipe"
-- ADD FOREIGN KEY ("user_id") REFERENCES "recipe_user" ("user_id") ON DELETE CASCADE ON UPDATE NO ACTION;
-- ALTER TABLE "user"
-- ADD FOREIGN KEY ("id") REFERENCES "recipe_user" ("user_id") ON DELETE CASCADE ON UPDATE NO ACTION;
--activity values
INSERT INTO activity(
    name,
    calories_per_hour,
    duration_minutes,
    total_calories
  )
VALUES ('Skiing, water skiing', 435, 60, 435),
  (gen_random_uuid(),
    'Cross country snow skiing, slow',
    508,
    60,
    508
  ),
  (gen_random_uuid(),
    'Cross country skiing, moderate',
    581,
    60,
    581
  ),
  (gen_random_uuid(),
    'Cross country skiing, vigorous',
    653,
    60,
    653
  ),
  (gen_random_uuid(),
    'Cross country skiing, racing',
    1016,
    60,
    1016
  ),
  (gen_random_uuid(),
    'Cross country skiing, uphill',
    1198,
    60,
    1198
  ),
  (gen_random_uuid(),
    'Snow skiing, downhill skiing, light',
    363,
    60,
    363
  ),
  (gen_random_uuid(),
    'Downhill snow skiing, moderate',
    435,
    60,
    435
  ),
  (gen_random_uuid(),
    'Downhill snow skiing, racing',
    581,
    60,
    581
  ),
  (gen_random_uuid(),
    'Coaching: football, basketball, soccer…',
    290,
    60,
    290
  ),
  (gen_random_uuid(),'Canoeing, rowing, light', 217, 60, 217),
  (gen_random_uuid(),
    'Canoeing, rowing, moderate',
    508,
    60,
    508
  ),
  (gen_random_uuid(),
    'Canoeing, rowing, vigorous',
    871,
    60,
    871
  ),
  (gen_random_uuid(),
    'Crew, sculling, rowing, competition',
    871,
    60,
    871
  ),
  (gen_random_uuid(),
    'Cycling, mountain bike, bmx',
    617,
    60,
    617
  ),
  (gen_random_uuid(),'Table tennis, ping pong', 290, 60, 290),
  (gen_random_uuid(),'Playing tennis', 508, 60, 508),
  (gen_random_uuid(),
    'Playing basketball, non game',
    435,
    60,
    435
  ),
  (gen_random_uuid(),
    'Coaching: football, basketball, soccer…',
    290,
    60,
    290
  ),
  (gen_random_uuid(),'Playing volleyball', 217, 60, 217),
  (gen_random_uuid(),'Water volleyball', 217, 60, 217),
  (gen_random_uuid(),
    'Coaching: football, basketball, soccer…',
    290,
    60,
    290
  ),
  (gen_random_uuid(),'Playing soccer', 508, 60, 508),
  (gen_random_uuid(),
    'Football or baseball, playing catch',
    181,
    60,
    181
  ),
  (gen_random_uuid(),'Softball or baseball', 363, 60, 363),
  (gen_random_uuid(),'Ballroom dancing, slow', 217, 60, 217),
  (gen_random_uuid(),'Ballroom dancing, fast', 399, 60, 399),
  (gen_random_uuid(),'Stretching, hatha yoga', 290, 60, 290),
  (gen_random_uuid(),
    'Martial arts, kick boxing',
    726,
    60,
    726
  ),
  (gen_random_uuid(),
    'Martial arts, judo, karate, jujitsu',
    726,
    60,
    726
  ),
  (gen_random_uuid(),
    'Martial arts, judo, karate, jujitsu',
    726,
    60,
    726
  ),
  (gen_random_uuid(),'Hockey, ice hockey', 581, 60, 581),
  (gen_random_uuid(),'Windsurfing, sailing', 217, 60, 217),
  (gen_random_uuid(),
    'Surfing, body surfing or board surfing',
    217,
    60,
    217
  ),
  (gen_random_uuid(),
    'Whitewater rafting, kayaking, canoeing',
    363,
    60,
    363
  ),
  (gen_random_uuid(),
    'Whitewater rafting, kayaking, canoeing',
    363,
    60,
    363
  ),
  (gen_random_uuid(),
    'Whitewater rafting, kayaking, canoeing',
    363,
    60,
    363
  ),
  (gen_random_uuid(),'Windsurfing, sailing', 217, 60, 217),
  (gen_random_uuid(),
    'Sailing, yachting, ocean sailing',
    217,
    60,
    217
  ),
  (gen_random_uuid(),'Skiing, water skiing', 435, 60, 435),
  (gen_random_uuid(),'Aerobics, step aerobics', 617, 60, 617),
  (gen_random_uuid(),'Water aerobics', 290, 60, 290),
  (gen_random_uuid(),
    'Water aerobics, water calisthenics',
    290,
    60,
    290
  ),
  (gen_random_uuid(),
    'Basketball, shooting baskets',
    327,
    60,
    327
  ),
  (gen_random_uuid(),'Water jogging', 581, 60, 581),
  (gen_random_uuid(),
    'Golf, walking and carrying clubs',
    327,
    60,
    327
  ),
  (gen_random_uuid(),
    'Golf, walking and pulling clubs',
    312,
    60,
    312
  ),
  (gen_random_uuid(),
    'Horseback riding, walking',
    181,
    60,
    181
  ),
  (gen_random_uuid(),
    'Pushing stroller or walking with children',
    181,
    60,
    181
  ),
  (gen_random_uuid(),'Race walking', 472, 60, 472),
  (gen_random_uuid(),
    'Cycling, <10 mph, leisure bicycling',
    290,
    60,
    290
  ),
  (gen_random_uuid(),'Unicycling', 363, 60, 363),
  (gen_random_uuid(),
    'Stationary cycling, very light',
    217,
    60,
    217
  ),
  (gen_random_uuid(),
    'Stationary cycling, light',
    399,
    60,
    399
  ),
  (gen_random_uuid(),
    'Stationary cycling, moderate',
    508,
    60,
    508
  ),
  (gen_random_uuid(),
    'Stationary cycling, vigorous',
    762,
    60,
    762
  ),
  (gen_random_uuid(),
    'Stationary cycling, very vigorous',
    908,
    60,
    908
  );
