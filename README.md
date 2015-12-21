Cyclic
-------------

![cyclic-image](https://cdn2-resources.ableton.com/resources/filer_thumbnails/public/2012/12/11/cyclic-waves.png__600x600_q85_crop_subsampling-2_upscale.jpg)


A library for running sql migrations in unit tests that need to access a database.


## Rationale

Often web applications use sql databases to store and present data.

In a pure unit test world, this dependency would be mocked and the application code would be unit tested itself.

In the harsh cold reality we live in we'd rather have our sql that is being executed validated and validate our sql is returning or modifying data correctly.

## What it does

Cyclic uses the golang [migrate](https://github.com/mattes/migrate) library to run schema migrations on a psuedo randomly named database and provides a connection to that database to use for tests.

## Requirements

- Supports mysql only for now.
- Your migrations should be in `./sql`
- Need database named `cyclic` listening on 127.0.0.1:3306
- Super user named `cyclic`
  - password `cyclic`

For each run of `WithDatabase` will create a db, run migrations and return a connection to the database created for use in your unit tests.  See mysql_test.go for an example.
