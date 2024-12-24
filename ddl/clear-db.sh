#!/bin/bash
psql -U postgres -d postgres -p 5432 -h localhost -c  "DROP SCHEMA deals CASCADE; DROP SCHEMA auth CASCADE"