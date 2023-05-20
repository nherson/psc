Base work
---
* ~~Ent models~~
* ~~Conversion from ufc struct to Ent model~~
* ~~Bulk upsert of historical data (mage target)~~
* ~~API: upcoming~~ REQUIRES TESTING
* ~~Mage: historical import (import:all and import:event XXXX)~~
* ~~Add timestamps to events for when they happened~~
* Clean up multiple connection functions (and rollback function)
* ~~Setup Chakra UI basics~~
* ~~Build UI parity with old site~~
* Refactor fightodds prompt+fuzzy-matcher into its own package, to be reused by tapology

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
* ~~Add "corner" to FightResult for a fighter, for visual consistency (Red left, Blue right)~~
* Include venue information for events
* Use Grid to equally space the fighter result cards