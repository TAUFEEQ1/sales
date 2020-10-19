# Sales App Demo Buffalo Backend!

First ensure that you have buffalo installed. https://gobuffalo.io/en/docs/getting-started/installation

## Database Setup

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... 
that are appropriate for your environment.

For this example the relevant database is mysql with credentials in the database.yml
i.e {
username:laboremus,
password:laboremus,
database:sales}

### Create Your Databases
Use the sales.sql dump file to initialise your database.

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## Dependencies
The command $ go get -u -v -f all will install all listed dependencies

## Compiled Binaries.
The command $ buffalo build will compile the application. 
