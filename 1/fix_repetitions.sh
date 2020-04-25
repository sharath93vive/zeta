#!/bin/bash -x

awk '{
    printf "%s", $1
    for (position=2; position<=NF; position++) {
        if ($position != $(position-1)) {
            printf "%s%s", OFS, $position
        }
    }
    print ""
}' $1

