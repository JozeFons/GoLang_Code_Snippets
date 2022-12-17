# GoLang_Code_Snippets

# MySQL

-- MySQL_Backup.go --

This example uses the mysqldump command to create a backup of the mydatabase database, and writes the output to a file called mydatabase.sql. You can then use this file to restore the database if needed.

-- MySQL_Restore.go --

This example prompts the user for confirmation before restoring the database, and only proceeds with the restore if the user enters "y". Also you can bypass this promt by commeting out.

# PostgreSQL

-- PostgreSQL_Backup.go --

This example uses the pg_dump command to create a binary file format (-Fc) backup of the mydatabase database, and writes the output to a file called mydatabase.dump. You can then use this file to restore the database if needed.



