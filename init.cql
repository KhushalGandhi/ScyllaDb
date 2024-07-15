CREATE KEYSPACE IF NOT EXISTS my_keyspace WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE my_keyspace;

CREATE TABLE IF NOT EXISTS todos (
                                     id UUID PRIMARY KEY,
                                     user_id UUID,
                                     title TEXT,
                                     description TEXT,
                                     status TEXT,
                                     created TIMESTAMP,
                                     updated TIMESTAMP
);