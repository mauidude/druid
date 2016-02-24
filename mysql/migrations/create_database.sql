-- create a druid database, make sure to use utf8 as encoding
CREATE SCHEMA  IF NOT EXISTS events;
ALTER DATABASE events charset=utf8;

-- create a druid user, and grant it all permission on the database we just created
GRANT ALL ON events.* TO 'druid'@'localhost' IDENTIFIED BY 'druid';
