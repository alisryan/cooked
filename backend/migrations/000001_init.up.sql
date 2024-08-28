CREATE TABLE recipes (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    cook_duration BIGINT NOT NULL,
    instructions TEXT NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    meal VARCHAR(10) NOT NULL CHECK (meal IN ('breakfast', 'lunch', 'dinner', 'snack'))
);
