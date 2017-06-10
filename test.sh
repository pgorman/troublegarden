#!/bin/sh
rm /tmp/troublegarden.db
sqlite3 /tmp/troublegarden.db < create_test_database.sql
$HOME/repo/go/bin/troublegarden -html ./html -sqlite /tmp/troublegarden.db -tlscert /tmp/test.crt -tlskey /tmp/test.key
