Base work
---
* ~~Ent models~~
* ~~Conversion from ufc struct to Ent model~~
* ~~Bulk upsert of historical data (mage target)~~
* ~~API: upcoming~~ REQUIRES TESTING
* ~~Mage: historical import (import:all and import:event XXXX)~~
* ~~Add timestamps to events for when they happened~~
* Clean up multiple connection functions (and rollback function)
* Setup Chakra UI basics
* Build UI parity with old site

Monthly process
---
* ~~Load in data from previous month (TODO: mage target)~~
* Delete UpcomingFights and UpcomingEvents (TODO: mage target)
* Handle upcoming events and fights
  <!-- * Create new fighter nodes (TODO: make mage target for this) -->
  * Register required fighter aliases (TODO: mage target)
  * Register upcoming fights and events, make new fighters as needed, wire in odds from odds site
    * Would be cool if this had some sort of "preview" output before committing to the DB

Nice to haves
---
* API to update upcoming events, fighters, odds, etc
  * Make it an API endpoint and hit it on a cron?



Pre-upsert counts

events: 315
fighters: 1369
fights: 3622
fighter_results: 7244