
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

-- INSERT INTO activity(
--     name,
--     calories_per_hour,
--     duration_minutes,
--     total_calories
--   )
-- VALUES ('Skiing, water skiing', 435, 60, 435),
--   (gen_random_uuid(),
--     'Cross country snow skiing, slow',
--     508,
--     60,
--     508
--   ),
--   (gen_random_uuid(),
--     'Cross country skiing, moderate',
--     581,
--     60,
--     581
--   ),
--   (gen_random_uuid(),
--     'Cross country skiing, vigorous',
--     653,
--     60,
--     653
--   ),
--   (gen_random_uuid(),
--     'Cross country skiing, racing',
--     1016,
--     60,
--     1016
--   ),
--   (gen_random_uuid(),
--     'Cross country skiing, uphill',
--     1198,
--     60,
--     1198
--   ),
--   (gen_random_uuid(),
--     'Snow skiing, downhill skiing, light',
--     363,
--     60,
--     363
--   ),
--   (gen_random_uuid(),
--     'Downhill snow skiing, moderate',
--     435,
--     60,
--     435
--   ),
--   (gen_random_uuid(),
--     'Downhill snow skiing, racing',
--     581,
--     60,
--     581
--   ),
--   (gen_random_uuid(),
--     'Coaching: football, basketball, soccer…',
--     290,
--     60,
--     290
--   ),
--   (gen_random_uuid(),'Canoeing, rowing, light', 217, 60, 217),
--   (gen_random_uuid(),
--     'Canoeing, rowing, moderate',
--     508,
--     60,
--     508
--   ),
--   (gen_random_uuid(),
--     'Canoeing, rowing, vigorous',
--     871,
--     60,
--     871
--   ),
--   (gen_random_uuid(),
--     'Crew, sculling, rowing, competition',
--     871,
--     60,
--     871
--   ),
--   (gen_random_uuid(),
--     'Cycling, mountain bike, bmx',
--     617,
--     60,
--     617
--   ),
--   (gen_random_uuid(),'Table tennis, ping pong', 290, 60, 290),
--   (gen_random_uuid(),'Playing tennis', 508, 60, 508),
--   (gen_random_uuid(),
--     'Playing basketball, non game',
--     435,
--     60,
--     435
--   ),
--   (gen_random_uuid(),
--     'Coaching: football, basketball, soccer…',
--     290,
--     60,
--     290
--   ),
--   (gen_random_uuid(),'Playing volleyball', 217, 60, 217),
--   (gen_random_uuid(),'Water volleyball', 217, 60, 217),
--   (gen_random_uuid(),
--     'Coaching: football, basketball, soccer…',
--     290,
--     60,
--     290
--   ),
--   (gen_random_uuid(),'Playing soccer', 508, 60, 508),
--   (gen_random_uuid(),
--     'Football or baseball, playing catch',
--     181,
--     60,
--     181
--   ),
--   (gen_random_uuid(),'Softball or baseball', 363, 60, 363),
--   (gen_random_uuid(),'Ballroom dancing, slow', 217, 60, 217),
--   (gen_random_uuid(),'Ballroom dancing, fast', 399, 60, 399),
--   (gen_random_uuid(),'Stretching, hatha yoga', 290, 60, 290),
--   (gen_random_uuid(),
--     'Martial arts, kick boxing',
--     726,
--     60,
--     726
--   ),
--   (gen_random_uuid(),
--     'Martial arts, judo, karate, jujitsu',
--     726,
--     60,
--     726
--   ),
--   (gen_random_uuid(),
--     'Martial arts, judo, karate, jujitsu',
--     726,
--     60,
--     726
--   ),
--   (gen_random_uuid(),'Hockey, ice hockey', 581, 60, 581),
--   (gen_random_uuid(),'Windsurfing, sailing', 217, 60, 217),
--   (gen_random_uuid(),
--     'Surfing, body surfing or board surfing',
--     217,
--     60,
--     217
--   ),
--   (gen_random_uuid(),
--     'Whitewater rafting, kayaking, canoeing',
--     363,
--     60,
--     363
--   ),
--   (gen_random_uuid(),
--     'Whitewater rafting, kayaking, canoeing',
--     363,
--     60,
--     363
--   ),
--   (gen_random_uuid(),
--     'Whitewater rafting, kayaking, canoeing',
--     363,
--     60,
--     363
--   ),
--   (gen_random_uuid(),'Windsurfing, sailing', 217, 60, 217),
--   (gen_random_uuid(),
--     'Sailing, yachting, ocean sailing',
--     217,
--     60,
--     217
--   ),
--   (gen_random_uuid(),'Skiing, water skiing', 435, 60, 435),
--   (gen_random_uuid(),'Aerobics, step aerobics', 617, 60, 617),
--   (gen_random_uuid(),'Water aerobics', 290, 60, 290),
--   (gen_random_uuid(),
--     'Water aerobics, water calisthenics',
--     290,
--     60,
--     290
--   ),
--   (gen_random_uuid(),
--     'Basketball, shooting baskets',
--     327,
--     60,
--     327
--   ),
--   (gen_random_uuid(),'Water jogging', 581, 60, 581),
--   (gen_random_uuid(),
--     'Golf, walking and carrying clubs',
--     327,
--     60,
--     327
--   ),
--   (gen_random_uuid(),
--     'Golf, walking and pulling clubs',
--     312,
--     60,
--     312
--   ),
--   (gen_random_uuid(),
--     'Horseback riding, walking',
--     181,
--     60,
--     181
--   ),
--   (gen_random_uuid(),
--     'Pushing stroller or walking with children',
--     181,
--     60,
--     181
--   ),
--   (gen_random_uuid(),'Race walking', 472, 60, 472),
--   (gen_random_uuid(),
--     'Cycling, <10 mph, leisure bicycling',
--     290,
--     60,
--     290
--   ),
--   (gen_random_uuid(),'Unicycling', 363, 60, 363),
--   (gen_random_uuid(),
--     'Stationary cycling, very light',
--     217,
--     60,
--     217
--   ),
--   (gen_random_uuid(),
--     'Stationary cycling, light',
--     399,
--     60,
--     399
--   ),
--   (gen_random_uuid(),
--     'Stationary cycling, moderate',
--     508,
--     60,
--     508
--   ),
--   (gen_random_uuid(),
--     'Stationary cycling, vigorous',
--     762,
--     60,
--     762
--   ),
--   (gen_random_uuid(),
--     'Stationary cycling, very vigorous',
--     908,
--     60,
--     908
--   );
