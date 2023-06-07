/*
 Jack, in December 2022:

 I found a problem with Gorm, or at least my usage of it several months back :/

 I'm not confident that this is actually a *bug* in Gorm, but ideally Gorm would've either thrown an error or handled
 it better.

 What follows is a very long comment with all the details so the context of "why on earth are we renaming foreign key
 constraints" doesn't die with me.

 Q: What's the problem?
 A: I think that when Gorm automigrates a struct that has multiple one-to-one relationships to another struct it
    only builds a foreign key constraint for the last relating field.
    In other words, if you set up multiple https://gorm.io/docs/has_one.html relations between two structs, they
    don't all get automigrated and set up properly--only the last occurring relation in the source struct gets set
    up with a foreign key.

 Q: Why "I think"?
 A: Because I had even more of an edge case--the two fields that had a one-to-one relation were actually named the
    same. I had embedding the same struct twice with just a different embeddedPrefix
    (https://gorm.io/docs/models.html#Embedded-Struct) and each copy wanted its own one-to-one relation to a row in
    another table.

 Q: Why on earth were you embedding the same struct twice; why did you come across this issue?
 A: Every Chart Release stores some version information. Every Changeset *stores that info twice*--once as "from" and
    once as "to", describing both sides of the diff. Inside Sherlock's model, I represented those fields as a common
    struct (ChartReleaseVersion) and just embedded it with Gorm into the structs for ChartRelease and Changeset.

 Q: Okay... but we don't use Gorm's automigrator anyway because it's a bit of an opaque system. How did the error
    occur in *Sherlock*?
 A: See you're forgetting that I had never once used SQL before writing Sherlock v2. So my workflow for anything
    significant was "write it in Go, let Gorm automigrate it against local Postgres, use JetBrains DataGrip to generate
    SQL for 'resolve structural diff between this local database and Sherlock's remote database,' and then paste it
    into one of these migration files." That worked great except that an issue with Gorm's automigrator made its way
    through into Sherlock's actual remote database.
    000015_add_changeset.up.sql contains the issue. See that the v2_changesets table has to_app_version_id and
    to_chart_version_id set up with foreign key constraints but the "from" fields earlier in the table didn't get the
    same treatment. Note that the constraint names omit the "to" embeddedPrefix entirely--rather than Gorm writing
    invalid SQL and errorring during automigration, it seems to handle the conflict by having the last one win.

 Q: Wow, alright. Why did you just notice now?
 A: Well, I actually noticed that something was up a while back. Because there wasn't a foreign key constraint on the
    "from" fields, Gorm couldn't load data from the association properly. I didn't know the reason why at the time, but
    the empty association nested structs I did notice. The associations got used in the controller to assemble the API
    output, and that was behaving weird, so I rewrote that little bit in
    https://github.com/broadinstitute/sherlock/pull/71. It was only when adding more one-to-one relationships for
    version-follow behavior that I went back and figured out the SQL from back then so I could add my new fields.

 Q: What's the fix?
 A: Here I'm renaming the original foreign key constraints to include the "to" and I'm adding the missing ones. When
    Gorm automigrates it won't know about that, but we don't use automigration outside of development and presumably
    Gorm works fine with the association being in its (current) broken state. I'm going to keep around some of the
    defensive-design I added in that earlier PR so that the controller won't literally omit data if this issue
    resurfaces due to automigration use in development.
 */

alter table v2_changesets
    rename constraint fk_v2_changesets_app_version to fk_v2_changesets_to_app_version;

alter table v2_changesets
    rename constraint fk_v2_changesets_chart_version to fk_v2_changesets_to_chart_version;

alter table v2_changesets
    add constraint fk_v2_changesets_from_app_version foreign key (from_app_version_id) references v2_app_versions;

alter table v2_changesets
    add constraint fk_v2_changesets_from_chart_version foreign key (from_chart_version_id) references v2_chart_versions;
