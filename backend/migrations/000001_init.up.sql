CREATE TABLE recipes (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    cook_duration BIGINT NOT NULL,
    instructions TEXT NOT NULL,
    image_url TEXT NOT NULL,
    meal TEXT NOT NULL CHECK (meal IN ('breakfast', 'lunch', 'dinner', 'snack'))
);
