#!/bin/bash
# eval $1 $2 $3 $4 
# Tip: To run this as entrypoint you have to put the shebang in your script and run chmod a+rx in your script too.
echo "$@"
exec "$@"