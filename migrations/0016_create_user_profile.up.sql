CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY,
    display_name TEXT,
    phone_number TEXT,
    physical_gender TEXT,
    birthday DATE,
    profile_picture_url TEXT,
    profile_picture_background_url TEXT,
    updated_at TIMESTAMP DEFAULT now(),

    CONSTRAINT fk_user_profile
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
