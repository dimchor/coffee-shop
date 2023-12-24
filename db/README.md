# MySQL

This container runs MySQL and uses a `.env` file to store the password. Create
the file in the parent directory and assign the password to the 
`DB_ROOT_PASSWORD` environmental variable.

This folder contains database dumps. A folder `mysql` is generated where all of 
the database's files are stored, so that it doesn't have go through the 
initialisation each time the container is started.
