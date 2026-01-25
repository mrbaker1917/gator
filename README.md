# gator

Welcome to gator! 
A fun way to get, store, and browse RSS feeds.

Before you can run this on your local machine, you need to make sure you have Postgres and Go installed.

To install gator, type into terminal: `go install gator`

You will also need to set up a config file, called ".gatorconfig.json" in your home directory.

Once those set-up steps are done, you can run gator in any terminal:

1. Navigate to the root folder.
2. Enter: `go run .` followed by any of the following commands:
# register <your name>
# login <your name> (use this, once you have registered, if there are multiple users registered to use the app on your computer)
# addfeed <name of RSS feed> <url of feed>
# feeds (shows feeds already added)
# follow <url of feed in feeds> (to follow an already added feed)
# following (shows the feeds you are following)
# agg <time between RSS feed requests> (this will update your feeds; to stop it getting them, press `control + C`)
# browse <number of posts> (to browse posts from a feed)

Enjoy!
