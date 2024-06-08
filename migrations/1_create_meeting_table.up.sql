CREATE TABLE IF NOT EXISTS meetings (
                                        id SERIAL PRIMARY KEY,
                                        title TEXT NOT NULL,
                                        date TIMESTAMP NOT NULL,
                                        participants TEXT NOT NULL,
                                        summary TEXT
);