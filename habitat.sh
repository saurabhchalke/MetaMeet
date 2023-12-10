#!/bin/sh

# Loads environment variables from a dot env (.env) file before calling the specified command.

# Parse -f and assign it to $OPTARG. Also assigns 'f' to $opt but that is not needed.
getopts f: opt
# If a path has been provided for the dot env file use that otherwise look in the working directory.
DOT_ENV_FILE_PATH=${OPTARG:-./.env}
# Shift option parameters away so that $1 becomes the command to run rather than '-f'.
shift $((OPTIND-1))
# Set the shell to export all variables set from this point onwards.
set -a
# If running Jest it will lose its colored output unless we set this.
FORCE_COLOR=true
# Load the dot env file.
. $DOT_ENV_FILE_PATH
# Execute the command.
exec $@
