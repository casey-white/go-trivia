CREATE TABLE quizzes (
    id SERIAL PRIMARY KEY,
    category VARCHAR(255),
    type VARCHAR(255),
    difficulty VARCHAR(255),
    question TEXT,
    correct_answer TEXT
);