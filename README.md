# GoLang_Code_Snippets

# MySQL

-- MySQL_Backup.go --

This example uses the "mysqldump" command to create a backup of the mydatabase database, and writes the output to a file called "mydatabase.sql". You can then use this file to restore the database if needed.

-- MySQL_Restore.go --

This example uses the "mysql" command to restore the mydatabase database from the "mydatabase.sql" file.

Additionally this example prompts the user for confirmation before restoring the database, and only proceeds with the restore if the user enters "y". Optionally you can bypass this promt by commeting out the code called "Prompt the user for confirmation".

# PostgreSQL

-- PostgreSQL_Backup.go --

This example uses the "pg_dump" command to create a binary file format (-Fc) backup of the "mydatabase" database, and writes the output to a file called "mydatabase.dump". You can then use this file to restore the database if needed.

-- PostgreSQL_Restore.go --

This example uses the "pg_restore" command to restore the "mydatabase" database from the "mydatabase.dump" file.

Additionally this example prompts the user for confirmation before restoring the database, and only proceeds with the restore if the user enters "y". Optionally you can bypass this promt by commeting out the code called "Prompt the user for confirmation".

# Slices_Manipulation

-- Slice_Manipulation.go --

Some practical examples “How To” remove a specific element or an element on specific index from slices[] or convert slices into different data type with GO

This is an example and you may need to add extra code to prevent some errors!

Visit and test in Playground: https://go.dev/play/p/4Hyilpr5POB
