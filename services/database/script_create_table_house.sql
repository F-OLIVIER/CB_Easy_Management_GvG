
-- Table des maisons
-- 
-- CREATE TABLE IF NOT EXISTS GroupGvG (
--     ID INTEGER PRIMARY KEY,
--     User_ID INTEGER NOT NULL,
--     GroupNumber INTEGER NOT NULL,
--     Unit1 VARCHAR(50) DEFAULT "",
--     Unit2 VARCHAR(50) DEFAULT "",
--     Unit3 VARCHAR(50) DEFAULT "",
--     Unit4 VARCHAR(50) DEFAULT "",
--     FOREIGN KEY (User_ID) REFERENCES Users (ID)
-- );

-- CREATE TABLE IF NOT EXISTS GroupTypeAtt (
--     ID INTEGER PRIMARY KEY,
--     User_ID INTEGER NOT NULL,
--     GroupNumber INTEGER NOT NULL,
--     Unit1 VARCHAR(50) DEFAULT "",
--     Unit2 VARCHAR(50) DEFAULT "",
--     Unit3 VARCHAR(50) DEFAULT "",
--     Unit4 VARCHAR(50) DEFAULT "",
--     FOREIGN KEY (User_ID) REFERENCES Users (ID)
-- );

-- CREATE TABLE IF NOT EXISTS GroupTypeDef (
--     ID INTEGER PRIMARY KEY,
--     User_ID INTEGER NOT NULL,
--     GroupNumber INTEGER NOT NULL,
--     Unit1 VARCHAR(50) DEFAULT "",
--     Unit2 VARCHAR(50) DEFAULT "",
--     Unit3 VARCHAR(50) DEFAULT "",
--     Unit4 VARCHAR(50) DEFAULT "",
--     FOREIGN KEY (User_ID) REFERENCES Users (ID)
-- );

-- CREATE TABLE IF NOT EXISTS NameGroupGvG (
--     ID INTEGER PRIMARY KEY,
--     GroupNumber INTEGER NOT NULL,
--     NameGroup VARCHAR(150) NOT NULL
-- );




-- CREATE TABLE IF NOT EXISTS Caserne (
--     ID INTEGER PRIMARY KEY,
--     User_ID INTEGER NOT NULL,
--     Unit1 INTEGER DEFAULT 0,
--     Unit2 INTEGER DEFAULT 0,
--     Unit3 INTEGER DEFAULT 0,
--     Unit4 INTEGER DEFAULT 0,
--     Unit5 INTEGER DEFAULT 0,
--     ...
--     FOREIGN KEY (User_ID) REFERENCES Users (ID)
-- );

-- CREATE TABLE IF NOT EXISTS CaserneMaitrise (
--     ID INTEGER PRIMARY KEY,
--     User_ID INTEGER NOT NULL,
--     Unit1 INTEGER DEFAULT 0,
--     Unit2 INTEGER DEFAULT 0,
--     Unit3 INTEGER DEFAULT 0,
--     Unit4 INTEGER DEFAULT 0,
--     Unit5 INTEGER DEFAULT 0,
--     ...
--     FOREIGN KEY (User_ID) REFERENCES Users (ID)
-- );
