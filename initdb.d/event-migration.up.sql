CREATE TABLE event (
id serial PRIMARY KEY,
user_id serial NOT NULL,
title VARCHAR ( 50 ),
description VARCHAR ( 50 ),
);