#!/bin/bash

MODE=development
#MODE=production
#MODE=none

npm run build -- --color --mode "${MODE}"
