CREATE TYPE "user_role" AS ENUM ('user', '_admin');
CREATE TYPE "user_gender" AS ENUM ('male', 'female');



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
