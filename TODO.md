Base work
---
* Ent models
* Conversion from ufc struct to Ent model
* Bulk upsert of historical data (mage target)

Monthly process
---
* Load in data from previous month (TODO: mage target)
* Delete UpcomingFights and UpcomingEvents (TODO: mage target)
* Handle upcoming events and fights
  <!-- * Create new fighter nodes (TODO: make mage target for this) -->
  * Register required fighter aliases (TODO: mage target)
  * Register upcoming fights and events, make new fighters as needed, wire in odds from odds site
    * Would be cool if this had some sort of "preview" output before committing to the DB

Nice to haves
---
* Mage target to _update_ odds
  * Make it an API endpoint and hit it on a cron?
* 