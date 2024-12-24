#!/bin/bash
psql -U postgres -d postgres -p 5432 -h localhost < "master.sql"