#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
wshdir="$workspace/src/github.com/sammy007"
if [ ! -L "$wshdir/open-wiseplat-pool" ]; then
    mkdir -p "$wshdir"
    cd "$wshdir"
    ln -s ../../../../../. open-wiseplat-pool
    cd "$root"
fi

# Set up the environment to use the workspace.
# Also add Godeps workspace so we build using canned dependencies.
GOPATH="$workspace"
GOBIN="$PWD/build/bin"
export GOPATH GOBIN

# Run the command inside the workspace.
cd "$wshdir/open-wiseplat-pool"
PWD="$wshdir/open-wiseplat-pool"

# Launch the arguments with the configured environment.
exec "$@"
